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

	kv.Put(context.TODO(), "/demo/A/B1", "BBB", clientv3.WithPrevKV())
	kv.Put(context.TODO(), "/demo/A/B2", "CCC", clientv3.WithPrevKV())
	kv.Put(context.TODO(), "/demo/A/B3", "DDD", clientv3.WithPrevKV())
	/*
		clientv3.WithFromKey() 表示针对的key操作是大于等于当前给定的key，比较危险
		clientv3.WithPrevKV() 表示返回的 response 中含有之前删除的值，否则
		下面的 delResp.PrevKvs 为空
	*/
	delResp, err := kv.Delete(context.TODO(), "/demo/A/B",
		clientv3.WithFromKey(), clientv3.WithPrevKV())
	if err != nil {
		fmt.Println(err)
	}
	// 查看被删除的 key 和 value 是什么
	if delResp.PrevKvs != nil {
		// if len(delResp.PrevKvs) != 0 {
		for _, kvpair := range delResp.PrevKvs {
			fmt.Println("已删除:", string(kvpair.Key), string(kvpair.Value))
		}
	}
}
