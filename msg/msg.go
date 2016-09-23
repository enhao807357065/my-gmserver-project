package msg

import (
	"github.com/name5566/leaf/network/json"
	"github.com/name5566/leaf/gate"
)

var Processor = json.NewProcessor()

// 在线的用户集合
var OnlineUsers = make(map[gate.Agent]struct{})

type BackType int
const (
	BackTypeUndefined 	BackType 	= iota	// 无
	BackTypeRegist 				= 1 	// 注册返回类型
	BackTypeLogin 				= 2 	// 登录返回类型
)

func init() {
	Processor.Register(&Hello{})
	Processor.Register(&Register{})
	Processor.Register(&Back{})
	Processor.Register(&Login{})
}

type Hello struct {
	Name string	`json:"name"`
}

type Register struct {
	Id              int64 		`json:"id,omitempty" bson:"_id,omitempty" description:"标识id"`
	Account		string		`json:"account,omitempty" bson:"account,omitempty" description:"账号"`
	Password	string		`json:"password,omitempty" bson:"password,omitempty" description:"密码"`
}

type Login struct {
	Account		string		`json:"account,omitempty" bson:"account,omitempty" description:"账号"`
	Password	string		`json:"password,omitempty" bson:"password,omitempty" description:"密码"`
}

type Back struct {
	Type		BackType		`json:"type"`
	Result		int			`json:"rt"`
	Message		interface{}		`json:"msg,omitempty"`
}