package internal

import (
	"github.com/name5566/leaf/gate"
)

var agents = make(map[gate.Agent]struct{})

// 创建和删除连接会调用
func init() {
	// NewAgent 和 CloseAgent 会被 LeafServer 的 gate 模块在连接建立和连接中断时调用
	// RegisterChanRPC 的第一个参数是 ChanRPC 的名字，第二个参数是 ChanRPC 的实现。这里的 NewAgent 和 CloseAgent 会被 LeafServer 的 gate 模块在连接建立和连接中断时调用。
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	//a := args[0].(gate.Agent)
	//_ = a

	a := args[0].(gate.Agent)
	agents[a] = struct{}{}
}

func rpcCloseAgent(args []interface{}) {
	//a := args[0].(gate.Agent)
	//_ = a

	a := args[0].(gate.Agent)
	delete(agents, a)
}
