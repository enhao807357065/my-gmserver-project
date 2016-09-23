package internal

import (
	"github.com/name5566/leaf/gate"
	"server/conf"
	"server/game"
	"server/msg"
	"time"
	"github.com/name5566/leaf/timer"
	"fmt"
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Server.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		TCPAddr:         conf.Server.TCPAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       msg.Processor,
		AgentChanRPC:    game.ChanRPC,
	}

	//game.ChanRPC.Go("NewAgent", "Test", "test")
	//broadcast()

}

// 测试代码：每隔10秒群发广播
func broadcast() {

	time.AfterFunc(time.Second * 10, func() {
		d := timer.NewDispatcher(10)	// chan长度10

		// cron expr
		cronExpr, err := timer.NewCronExpr("* * * * * *")
		if err != nil {
			return
		}

		// cron
		//var c *timer.Cron
		d.CronFunc(cronExpr, func() {

			fmt.Println("len(game.OnlineUsers): ", len(msg.OnlineUsers))
			for a := range msg.OnlineUsers {
				fmt.Println("a: ", a.UserData())
				//user := a.UserData().(*msg.Register)
				//if user.Account == "test" {
				//	//a.Close()	// 关闭连接?
				//	//a.Destroy()	// 关闭连接?
				//	//delete(msg.OnlineUsers, a)
				//}
				//a.WriteMsg(&msg.Back{Message: "success"})
			}
		})

		// dispatch
		//(<-d.ChanTimer).Cb()

		for {
			(<-d.ChanTimer).Cb()	// 不会自动执行重复执行逻辑，需多次调用

			timer := time.NewTimer(time.Second * 10)

			<-timer.C

		}

	})

}
