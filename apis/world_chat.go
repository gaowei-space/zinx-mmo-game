package apis

import (
	"fmt"
	"zinx-mmo-game/core"
	"zinx-mmo-game/pb"
	"zinx/ziface"
	"zinx/znet"

	"google.golang.org/protobuf/proto"
)

// 世界聊天的路由业务

type WorldChatApi struct {
	znet.BaseRouter
}

func (wc *WorldChatApi) Handle (request ziface.IRequest) {
	// 1. 解析客户端传递进来的proto协议
	proto_msg := &pb.Talk{}
	err := proto.Unmarshal(request.GetData(), proto_msg)
	if err != nil {
		fmt.Println("Talk unmarshal err", err)
		return
	}

	// 2. 获取玩家ID
	pid, err := request.GetConnection().GetProperty("pid")

	// 3. 根据pid得到当前玩家的player对象
	player := core.WorldManagerObj.GetPlayerByPid(pid.(int32))


	// 4. 将消息广播给其他全部在线的玩家
	player.Talk(proto_msg.Content)
}