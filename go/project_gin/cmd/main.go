// hello world example
package main

import (
	"context"
	"flag"
	"log"
	"os"

	"lovelake.cn/texiusi/pkg/config"
	"lovelake.cn/texiusi/pkg/web"
)

var (
	configFlag = flag.String("config", "./etc/config.yaml", "配置文件，default: ./etc/config.yaml")
)

func main() {

	// 设置Context
	_, cancel := context.WithCancel(context.Background())
	flag.Parse()

	// 读取配置文件
	_, err := config.New(configFlag)
	if err != nil {
		log.Fatal(err.Error())
	}

	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	// 配置日志

	// 配置路由

	// 配置Cache

	// 配置Database

	// 启动HTTP
	if err := web.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	cancel()
}
