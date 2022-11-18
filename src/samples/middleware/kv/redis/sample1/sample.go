package main

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"sync"
	"time"
)

var rbd redis.UniversalClient

func initClient() (err error) {
	rbd = redis.NewClient(&redis.Options{
		Addr:     "192.168.136.24:6379",
		Password: "password",
		DB:       0,
	})

	_, err = rbd.Ping().Result()
	return err
}

func initSentinelClient() (err error) {
	rbd = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"192.168.136.24:26379"},
		Password:      "",
	})
	_, err = rbd.Ping().Result()
	return err
}

func initClusterClient() (err error) {
	rbd = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{":7000", ":7001"},
		Password: "",
	})
	_, err = rbd.Ping().Result()
	return err
}

func setExample() {
	err := rbd.Set("score", 100, 0).Err()
	if err != nil {
		log.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := rbd.Get("score").Result()
	if err != nil {
		log.Printf("get score failed, err:%v\n", err)
		return
	}
	log.Println("score", val)

	val2, err := rbd.Get("name").Result()
	if err == redis.Nil {
		log.Println("name does not exist")
	} else if err != nil {
		log.Printf("get name failed, err:%v\n", err)
		return
	} else {
		log.Println("name", val2)
	}
}

func zsetExample() {
	zsetKey := "language_rank"
	languages := []redis.Z{
		redis.Z{Score: 90.0, Member: "Golang"},
		redis.Z{Score: 98.0, Member: "Java"},
		redis.Z{Score: 95.0, Member: "Python"},
		redis.Z{Score: 97.0, Member: "JavaScript"},
		redis.Z{Score: 99.0, Member: "C/C++"},
	}
	// zadd
	num, err := rbd.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		log.Printf("zadd failed, err:%v\n", err)
		return
	}
	log.Printf("zadd %d succ.\n", num)

	// 把Golang的分数加10
	newScore, err := rbd.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		log.Printf("zincrby failed, err:%v\n", err)
		return
	}
	log.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret, err := rbd.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		log.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		log.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rbd.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		log.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

func pipelineExample() {
	pipe := rbd.Pipeline()
	incr := pipe.Incr("pipeline_counter")
	pipe.Expire("pipeline_counter", time.Hour)

	_, err := pipe.Exec()
	log.Println(incr.Val(), err)
}

func pipelineExample1() {
	var incr *redis.IntCmd
	_, err := rbd.Pipelined(func(pipe redis.Pipeliner) error {
		incr = pipe.Incr("pipelined_counter")
		pipe.Expire("pipelined_counter", time.Hour)
		return nil
	})
	log.Println(incr.Val(), err)
}

// MULTI/EXEC命令
func txpipelineExample() {
	pipe := rbd.TxPipeline()
	incr := pipe.Incr("tx_pipeline_counter")
	pipe.Expire("tx_pipeline_counter", time.Hour)

	_, err := pipe.Exec()
	log.Println(incr.Val(), err)
}

// 在watch期间，如果key值变化，pipeline中的命令不会执行
func watchExample() {
	key := "watch_count"
	err := rbd.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			pipe.Set(key, n+1, 0)
			return nil
		})
		return err
	}, key)
	if err != nil {
		log.Printf("watch failed, err:%v\n", err)
	}
}

const routineCount = 100

func txSet() {
	increment := func(key string) error {
		txf := func(tx *redis.Tx) error {
			// 获得当前值或零值
			n, err := tx.Get(key).Int()
			if err != nil && err != redis.Nil {
				return err
			}

			// 实际操作（乐观锁定中的本地操作）
			n++
			// 仅在监视的Key保持不变的情况下运行
			_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
				// pipe 处理错误情况
				pipe.Set(key, n, 0)
				return nil
			})
			return err
		}

		for retries := routineCount; retries > 0; retries-- {
			err := rbd.Watch(txf, key)
			if err != redis.TxFailedErr {
				return err
			}
			// 乐观锁丢失
		}
		return errors.New("increment reached maximum number of retries")
	}

	var wg sync.WaitGroup
	wg.Add(routineCount)
	for i := 0; i < routineCount; i++ {
		go func() {
			defer wg.Done()

			if err := increment("counter3"); err != nil {
				fmt.Println("increment error:", err)
			}
		}()
	}
	wg.Wait()

	n, err := rbd.Get("counter3").Int()
	fmt.Println("ended with", n, err)
}

// TODO: pub/sub
func pubsubExample() {
	channelName := "mychannel1"
	pubsub := rbd.Subscribe(channelName)

	// Wait for confirmation that subscription is created before publishing anything.
	_, err := pubsub.Receive()
	if err != nil {
		panic(err)
	}

	// Go channel which receives messages.
	ch := pubsub.Channel()

	// Publish a message.
	err = rbd.Publish(channelName, "hello").Err()
	if err != nil {
		panic(err)
	}

	time.AfterFunc(time.Second, func() {
		// When pubsub is closed channel is closed too.
		_ = pubsub.Close()
	})

	// Consume message.
	for msg := range ch {
		log.Println(msg.Channel, msg.Payload)
	}
}

func main() {
	err := initClient()
	if err != nil {
		log.Fatal("init error")
	}

	pubsubExample()
}
