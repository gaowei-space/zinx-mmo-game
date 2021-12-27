package main

import (
	"fmt"
	"zinx-mmo-game/apis"
	"zinx-mmo-game/core"
	"zinx/ziface"
	"zinx/znet"
)

// 当前客户端建立连接之后的hook函数
func OnConnecionAdd(conn ziface.IConnection) {
	player := core.NewPlayer(conn)

	// 给客户端发送msgID:1的消息，目的是同步客户端当前玩家ID
	player.SyncPid()

	// 给客户端发送msgID:200的消息，目的是广播当前用户端至其他用户端
	player.BroadCastStartPosition()

	// 将当前上线的玩家添加至WorldManager中
	core.WorldManagerObj.AddPlayer(player)

	// 将该连接绑定一个pid属性
	conn.SetProperty("pid", player.Pid)

	// 同步周边玩家上线信息，与显示周边玩家信息
	player.SyncSurrounding()

	fmt.Println("=====> Player pidId = ", player.Pid, " arrived <=====")
}

func main() {
	// 创建服务句柄
	s := znet.NewServer("MMO GAME SERVER")

	// 注册客户端连接建立函数
	s.SetOnConnStart(OnConnecionAdd)

	s.AddRouter(2, &apis.WorldChatApi{})

	// 启动
	s.Serve()
}
