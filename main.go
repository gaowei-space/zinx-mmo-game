package main

import "zinx/znet"

func main() {
	// 创建服务句柄
	s := znet.NewServer("MMO GAME SERVER")

	// 启动
	s.Serve()
}
