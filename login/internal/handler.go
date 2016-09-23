package internal

import (
	"reflect"
	"server/msg"
	"github.com/name5566/leaf/gate"
	"server/dao"
	"strings"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {

	handleMsg(&msg.Register{}, handlerRegiest)
	handleMsg(&msg.Login{}, handlerLogin)

}

func handlerRegiest(args []interface{}) {

	m := args[0].(*msg.Register)

	a := args[1].(gate.Agent)

	err := dao.InsertUser(m)

	if err != nil {
		a.WriteMsg(&msg.Back{Type: msg.BackTypeRegist, Result: 1, Message: "插入数据库失败!"})
		return
	}

	a.WriteMsg(&msg.Back{Type: msg.BackTypeRegist, Result: 0})
}

func handlerLogin(args []interface{}) {

	m := args[0].(*msg.Login)

	a := args[1].(gate.Agent)

	if len(strings.TrimSpace(m.Account)) <= 0 || len(strings.TrimSpace(m.Password)) <= 0 {
		a.WriteMsg(&msg.Back{Type: msg.BackTypeLogin, Result: 1, Message: "账户或密码格式错误!"})
		return
	}

	dbUser, err := dao.FindUserByAccount(m.Account)
	if err != nil {
		a.WriteMsg(&msg.Back{Type: msg.BackTypeLogin, Result: 2, Message: "查询数据库错误!"})
		return
	}

	if dbUser == nil {
		a.WriteMsg(&msg.Back{Type: msg.BackTypeLogin, Result: 3, Message: "没有此用户!"})
		return
	}

	if dbUser.Password != m.Password {
		a.WriteMsg(&msg.Back{Type: msg.BackTypeLogin, Result: 4, Message: "密码不正确!"})
		return
	}

	a.SetUserData(dbUser)	// 向agent中存入用户基本信息
	msg.OnlineUsers[a] = struct{}{}	// 向在线用户中存入相应信息
	a.WriteMsg(&msg.Back{Type: msg.BackTypeLogin, Result: 0, Message: "登录成功!"})
}