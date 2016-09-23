package dao

import (
	"gopkg.in/mgo.v2"
	"time"
	"server/msg"
	"server/util"
	"gopkg.in/mgo.v2/bson"
)

var (
	db_user = "user"
)

func InsertUser(user *msg.Register) (err error) {

	c := MC(db_user)
	defer c.Close()

	user.Id = util.NewInsertId()
	TRY : {
		err := c.Insert(user)
		if mgo.IsDup(err) {
			time.Sleep(time.Millisecond)
			goto TRY
		}
	}

	return err
}

func FindUserByAccount(account string) (user *msg.Register, err error) {

	c := MC(db_user)
	defer c.Close()

	err = c.Find(bson.M{"account": account}).One(&user)
	return user, err
}