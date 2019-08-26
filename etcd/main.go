package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/pkg/transport"
	"time"
)

func main() {
	tlsInfo := transport.TLSInfo{
		CertFile:      `client.pem`,
		KeyFile:       `client-key.pem`,
		TrustedCAFile: `ca.pem`,
	}
	config, err := tlsInfo.ClientConfig()

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"https://172.20.10.12:2379"},
		DialTimeout: 3 * time.Second,
		TLS:         config,
	})
	if err != nil {
		panic(fmt.Sprintf("connect etcd err: %v", err))
	}
	defer cli.Close()

	kv := clientv3.NewKV(cli)
	if _, err = kv.Put(context.TODO(), "/pods/podname", "122333"); err != nil {
		panic(fmt.Sprintf("put ip to etcd err: %v", err))
	}

	getResp, _ := kv.Get(context.TODO(), "/pods/podname")
	if err != nil {
		panic(fmt.Sprintf("get ip from etcd err: %v", err))
	}
	fmt.Printf("get ip from etcd, key: %s, value: %s\n", getResp.Kvs[0].Key, getResp.Kvs[0].Value)
}
