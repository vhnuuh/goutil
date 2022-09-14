package main

import (
	"fmt"
	"time"
)

func chanProducer(in chan *Task) {
	for i := 1; ; i++ {
		task := &Task{
			Score: i,
		}
		in <- task
		time.Sleep(1e8)
	}
}

func chanWorker(index int, in chan *Task) {
	for {
		t := <-in
		fmt.Printf("worker %d process %d\n", index, t.Score)
		time.Sleep(1e9)
	}
}
