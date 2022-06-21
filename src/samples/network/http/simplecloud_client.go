package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const contentTypeJson = "application/json"

const urlPath = "http://127.0.0.1:8554/mcm/paas-cloud/v1/gotty-logs"

type GottyLogRequest struct {
	ClusterId   string `json:"clusterId,omitempty"`
	Namespace   string `json:"namespace,omitempty"`
	Pod         string `json:"pod,omitempty"`
	Container   string `json:"container,omitempty"`
	Command     string `json:"command,omitempty"`
	CommandLine string `json:"commandLine,omitempty"`
	UserId      string `json:"userId,omitempty"`
	ExecuteTime string `json:"executeTime,omitempty"`
}

func FlushLogs(flushLogs []*GottyLogRequest) {
	body, err := json.Marshal(flushLogs)
	fmt.Println(string(body))
	if err == nil {
		resp, err := http.Post(urlPath, contentTypeJson, bytes.NewBuffer(body))
		if err != nil {
			log.Printf("post logs err: %s", err.Error())
		} else {
			defer resp.Body.Close()
		}
	} else {
		log.Printf("marshal logs err: %s", err.Error())
	}
}

func buildLog(title, log *string) *GottyLogRequest {
	parts := strings.Split(*title, "|")
	commands := strings.Split(*log, " ")
	return &GottyLogRequest{
		ClusterId:   parts[0],
		Namespace:   parts[1],
		Pod:         parts[2],
		Container:   parts[3],
		Command:     commands[0],
		CommandLine: *log,
		UserId:      parts[4],
		ExecuteTime: time.Now().Format("2006-01-02 15:04:05"),
	}
}

type LogHttpHandler struct {
	cacheLogReqs []*GottyLogRequest
	flushing     int32
	logMu        sync.Mutex
}

var (
	httpHandler = &LogHttpHandler{
		cacheLogReqs: make([]*GottyLogRequest, 0, 10),
		flushing:     0,
	}
	title   = "1|2|3|4|5"
	content = "kill"
)

func (handler *LogHttpHandler) swapFlush() {
	log.Println("Start swapFlush")
	atomic.AddInt32(&handler.flushing, 1)
	flushLogs := handler.cacheLogReqs
	handler.cacheLogReqs = make([]*GottyLogRequest, 0, 10)
	defer FlushLogs(flushLogs)
	atomic.AddInt32(&handler.flushing, -1)
}

func (handler *LogHttpHandler) Log(title, cmd *string) {
	log.Println("Start swapFlush")
	handler.logMu.Lock()
	logReq := buildLog(title, cmd)
	handler.cacheLogReqs = append(handler.cacheLogReqs, logReq)
	if len(handler.cacheLogReqs) >= 1 {
		handler.swapFlush()
	}
	handler.logMu.Unlock()
}

func main() {

	httpHandler.Log(&title, &content)
	time.Sleep(5 * time.Second)
	fmt.Println("success")
}
