package dao

import (
	"gopkg.in/mgo.v2"
	"server/conf"
	"gopkg.in/mgo.v2/bson"
	"github.com/name5566/leaf/log"
)

var c *mgo.Database

type User struct {
	BId		bson.ObjectId		`bson:"_id"`
	Id		int64			`bson:"id"`
	Name		string			`bson:"name"`
	Age		int			`bson:"age"`
}

func init() {
	session, err := mgo.Dial(conf.Server.MongoDbUrl)
	if err != nil {
		log.Fatal("db %v", err)
	}
	c = session.DB("test")
	//user := User{}
	//err = c.C("user").Find(bson.M{"id": 2}).One(&user)
	//fmt.Println("err: ", err)
	//fmt.Println("user: ", user)

	//err = c.C("user").Insert(User{BId: bson.NewObjectId(), Id: 2, Name: "test2", Age: 22})
	//fmt.Println("err; ", err)

	//info, err := c.C("user").RemoveAll(bson.M{"id": 2})
	//fmt.Println("info: ", info) // info返回受影响的行数等信息
	//fmt.Println("err: ", err)

	//err = c.C("user").Update(bson.M{"id": 2}, bson.M{"$set": bson.M{"name": "刘备"}})
	//fmt.Println("err: ", err)
}

func Orm() *mgo.Database {
	return c
}