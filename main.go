package main

import (
	"fmt"
	"hello/boot"
	"hello/pkg/server"
	"log"
)

func main() {
	fmt.Println("服务初始化中...")
	if err := boot.Bootstrap("config/config.toml"); nil != err {
		log.Fatal(err)
		return
	}

	svrman.RegisterServer(boot.NewServer())
	if err := svrman.Start(); nil != err {
		log.Fatal(err)
	}

	return
}
