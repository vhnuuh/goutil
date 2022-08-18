package main

import (
	"context"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/grpclog"
	"log"
	"os"
	"time"
)

type ConfigStruct struct {
	Addr           string `json:"addr"`
	AesKey         string `json:"aes_key"`
	HTTPS          bool   `json:"https"`
	Secret         string `json:"secret"`
	PrivateKeyPath string `json:"private_key_path"`
	CertFilePath   string `json:"cert_file_path"`
}

var (
	configPath  = `/configs/remote_config.json`
	appConfig   ConfigStruct
	endpoints   = []string{"http://192.168.136.24:2379"}
	dialTimeout = 5 * time.Second
	cli         *clientv3.Client
	kv          clientv3.KV
	conf        = ConfigStruct{
		Addr:           "127.0.0.1:1080",
		AesKey:         "01B345B7A9ABC00F0123456789ABCDAF",
		HTTPS:          false,
		Secret:         "",
		PrivateKeyPath: "",
		CertFilePath:   "",
	}
)

func init() {
	clientv3.SetLogger(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))
	cfg := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	}

	cli, err := clientv3.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	initConfig()
}

func watchAndUpdate() {
	w := cli.Watcher(configPath, nil)
	go func() {
		// watch 该节点下的每次变化
		for {
			resp, err := w.Next(context.Background())
			if err != nil {
				log.Fatal(err)
			}
			log.Println("new values is ", resp.Node.Value)

			err = json.Unmarshal([]byte(resp.Node.Value), &appConfig)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()
}

func initConfig() {
	val, _ := json.Marshal(conf)
	pr, _ := kv.Put(context.TODO(), configPath, string(val))

	resp, err := kv.Get(context.TODO(), configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(resp.Kvs[0].Value, &appConfig)
	if err != nil {
		log.Fatal(err)
	}
}

func getConfig() ConfigStruct {
	return appConfig
}

func main() {
	// init your app
	conf := getConfig()
	watchAndUpdate()
	fmt.Println(conf)
}
