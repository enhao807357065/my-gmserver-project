package internal

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/gate"
	"reflect"
	"server/msg"
	"fmt"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.Hello{}, handleHello)

	//handler(&msg.Register{}, handlerRegiest)

}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.Hello)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("hello %v", m.Name)

	// 给发送者回应一个 Hello 消息
	a.WriteMsg(&msg.Hello{
		Name: "client",
	})
}

func handlerRegiest(args []interface{}) {
	m := args[0].(*msg.Register)

	a := args[1].(gate.Agent)

	fmt.Println("m: ", m)
	//err := c.Orm().C("user").Insert(&m)
	//fmt.Println("err: ", err)

	a.WriteMsg(&msg.Back{Message: "success"})
}
