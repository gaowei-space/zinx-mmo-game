package core

import (
	"fmt"
	"math/rand"
	"sync"
	"zinx-mmo-game/pb"
	"zinx/ziface"

	"github.com/golang/protobuf/proto"
)

type Player struct {
	Pid int32 // 用户ID
	Conn ziface.IConnection // 当前玩家的连接
	X float32 // 平面X坐标
	Y float32 // 高度
	Z float32 // 平面y坐标
	V float32 // 旋转0-360角度
}

var PidGen int32 = 1 // 用户ID生成器
var IdLock sync.Mutex // 保护PidGen的互斥锁

func NewPlayer(conn ziface.IConnection) *Player {
	// 生成一个PID
	IdLock.Lock()
	id :=PidGen
	PidGen ++
	IdLock.Unlock()

	p := &Player{
		Pid: id,
		Conn: conn,
		X: float32(160+rand.Intn(10)), // 随机在160坐标点 基于X轴偏移若干坐标
		Y: 0,
		Z: float32(140+rand.Intn(20)), // 随机在140坐标点 基于Y轴偏移若干坐标
		V: 0,
	}

	return p
}

func (p *Player) SendMsg(msgId int32, data proto.Message) {
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("marshal data err", err)
		return
	}

	if p.Conn == nil {
		fmt.Println("conn is closed")
		return
	}

	if err := p.Conn.SendMsg(uint32(msgId), msg); err != nil {
		fmt.Println("sendMsg err", err)
		return
	}

	return
}

func (p *Player) SyncPid() {
	data := &pb.SyncPid{
		Pid: p.Pid,
	}

	p.SendMsg(1, data)
}

func (p *Player) BroadCastStartPosition() {
	data := &pb.BroadCast{
		Pid: p.Pid,
		Tp: 2,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	p.SendMsg(200, data)
}