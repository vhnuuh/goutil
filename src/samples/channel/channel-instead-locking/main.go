package main

import "time"

func run_locking() {
	pool := &Pool{
		Tasks: make([]*Task, 0, 20),
	}
	go lockProducer(pool)
	go lockWorker(pool)
	time.Sleep(100 * 1e9)
}

func run_channel() {
	tasks := make(chan *Task)
	go chanProducer(tasks)
	for i := 0; i < 5; i++ {
		go chanWorker(i, tasks)
	}
	time.Sleep(100 * 1e9)
}

func main() {
	run_channel()
}
