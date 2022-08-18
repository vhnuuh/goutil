package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

func incr() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.136.24:6379",
		Password: "password",
		DB:       0,
	})

	lockKey := "counter_lock"
	counterKey := "counter"

	// lock
	resp := client.SetNX(lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()

	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result: ", lockSuccess)
		return
	}

	// counter++
	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil || err == redis.Nil {
		cntValue++
		resp := client.Set(counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			// log err
			fmt.Println("set value error!")
		}
	}
	fmt.Println("current counter is ", cntValue)

	delResp := client.Del(lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		fmt.Println("unlock success!")
	} else {
		fmt.Println("unlock failed", err)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
}
