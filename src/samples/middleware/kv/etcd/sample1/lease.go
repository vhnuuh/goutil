package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func respLease(keepRespChan <-chan *clientv3.LeaseKeepAliveResponse) {
	for {
		select {
		case keepResp := <-keepRespChan:
			if keepResp == nil {
				log.Println("租约已经失效了")
				return
			} else { // 每秒会续租一次, 所以就会受到一次应答
				log.Println("收到自动续租应答:", keepResp.ID)
			}
		}
	}
}

func main() {
	config := clientv3.Config{
		Endpoints:   []string{"192.168.136.24:2379"}, // 集群列表
		DialTimeout: 5 * time.Second,
	}

	// 建立一个客户端
	client, err := clientv3.New(config)
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个lease（租约）对象
	lease := clientv3.NewLease(client)
	// 申请一个10秒的租约
	leaseGrantResp, err := lease.Grant(context.TODO(), 10)
	if err != nil {
		log.Fatal(err)
	}
	// 拿到租约的ID
	leaseId := leaseGrantResp.ID

	// 主动续约
	//resp, err := lease.KeepAliveOnce(context.TODO(), leaseId)
	// 自动永久续租
	keepRespChan, err := lease.KeepAlive(context.TODO(), leaseId)
	if err != nil {
		log.Fatal(err)
	}

	// 处理续约应答的协程
	go respLease(keepRespChan)

	// 获得kv API子集
	kv := clientv3.NewKV(client)

	// Put一个KV, 让它与租约关联起来, 从而实现10秒后自动过期
	putResp, err := kv.Put(context.TODO(), "/demo/A/B1", "hello", clientv3.WithLease(leaseId))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("写入成功:", putResp.Header.Revision)

	// 定时的看一下key过期了没有
	for {
		getResp, err := kv.Get(context.TODO(), "/demo/A/B1")
		if err != nil {
			log.Fatal(err)
		}
		if getResp.Count == 0 {
			fmt.Println("kv过期了")
			break
		}
		fmt.Println("还没过期:", getResp.Kvs)
		time.Sleep(2 * time.Second)
	}
}
