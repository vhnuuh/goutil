package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	config := clientv3.Config{
		Endpoints:   []string{"192.168.136.24:2379"}, // 集群列表
		DialTimeout: 5 * time.Second,
	}

	// 建立一个客户端
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 用于读写etcd的键值对
	kv := clientv3.NewKV(client)

	_, _ = kv.Put(context.TODO(), "/demo/A/B", "BBB", clientv3.WithPrevKV())
	_, _ = kv.Put(context.TODO(), "/demo/A/C", "CCC", clientv3.WithPrevKV())
	// 	读取/demo/A/为前缀的所有key
	// clientv3.WithPrefix() , clientv3.WithCountOnly() 可以有多个并以 逗号分隔即可
	getResp, err := kv.Get(context.TODO(), "/demo/A/", clientv3.WithPrefix() /*,clientv3.WithCountOnly()*/)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(getResp.Kvs, getResp.Count)
	for _, resp := range getResp.Kvs {
		fmt.Printf("key: %s, value:%s\n", string(resp.Key), string(resp.Value))
	}
}
