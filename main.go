package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {

	config := clientv3.Config{
		Endpoints:   []string{"localhost:2479"},
		DialTimeout: 10 * time.Second,
		Username: "root",
		Password: "12345",
	}
	client, err := clientv3.New(config)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	kv := clientv3.NewKV(client)

	// 写入
	kv.Put(context.TODO(), "name1", "lesroad")
	kv.Put(context.TODO(), "name2", "haha")

	if getResp1, err1 := kv.Get(context.TODO(), "name2"); err != nil {
		fmt.Println(err1)
		return
	} else {
		// 获取成功
		fmt.Println(getResp1.Kvs)
	}

	// 读取name为前缀的所有key
	if getResp, err := kv.Get(context.TODO(), "name", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
		return
	} else {
		// 获取成功
		fmt.Println(getResp.Kvs)
	}

	// 删除name为前缀的所有key
	// if _, err = kv.Delete(context.TODO(), "name", clientv3.WithPrevKV()); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

}
