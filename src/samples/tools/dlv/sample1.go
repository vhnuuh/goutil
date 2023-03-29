package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func count(i, j int) int {
	yz := 5
	result := (i + j) * yz
	return result
}

func randHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var (
		i = rand.Intn(20)
		j = rand.Intn(20)
	)

	result := count(i, j)
	_, _ = w.Write([]byte(fmt.Sprintf("%d", result)))
}

func main() {
	http.HandleFunc("/rand", randHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("start server fail: %v", err)
	}
}
