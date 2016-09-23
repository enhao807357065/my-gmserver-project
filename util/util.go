package util

import (
	"time"
	"server/msg"
	"github.com/name5566/leaf/gate"
)

// 生成新id(临时解决方案)
func NewInsertId() int64 {
	return time.Now().UnixNano() / 1000000 - 1450427758000 //毫秒
}

// 强制下线某用户(断开连接)
func Offline(a gate.Agent) {
	//a.Close()	// 关闭连接?
	a.Destroy()	// 关闭连接?
	delete(msg.OnlineUsers, a)
}