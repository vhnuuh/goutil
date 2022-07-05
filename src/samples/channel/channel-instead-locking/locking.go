package main

import (
	"fmt"
	"sync"
	"time"
)

type Pool struct {
	Mu    sync.Mutex
	Tasks []*Task
}

func lockProducer(pool *Pool) {
	for i := 1; ; i++ {
		pool.Mu.Lock()
		fmt.Println("producer acquire lock")
		task := &Task{
			Score: i,
		}
		pool.Tasks = append(pool.Tasks, task)
		pool.Mu.Unlock()
		fmt.Println("producer release lock")
		time.Sleep(1e9)
	}
}

func lockWorker(pool *Pool) {
	for {
		pool.Mu.Lock()
		fmt.Println("worker acquire lock")
		if len(pool.Tasks) != 0 {
			task := pool.Tasks[0]
			pool.Tasks = pool.Tasks[1:]
			fmt.Println(task.Score)
		}
		pool.Mu.Unlock()
		fmt.Println("worker release lock")
		time.Sleep(1e9)
	}
}
