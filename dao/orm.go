package dao

import (
	"gopkg.in/mgo.v2"
	"server/conf"
	"gopkg.in/mgo.v2/bson"
)

var (
	mongo_uri string
	mongo_db string
	MAX_RETRY = 3
	gsession  *mgo.Session
	pool      chan *mgo.Session
)

func init() {

	mongo_uri = conf.Server.MongoDbUrl
	if mongo_uri == "" {
		panic("MongoDbUrl cannot empty!")
	}
	mongo_db = conf.Server.MongoDb
	if mongo_db == "" {
		panic("MongoDb cannot empty!")
	}
	mgo.SetDebug(true)
	mgo.SetStats(true)

	gsession = NewSession()

}

func NewSession() *mgo.Session {
	session, err := mgo.Dial(mongo_uri)
	if err != nil {
		panic(err)
	}
	return session
}

func MC(collection string) *MCollection {
	//	s := NewSession()
	s := gsession.Copy()
	//s := <-pool
	c := s.DB(mongo_db).C(collection)
	return &MCollection{Collection: c, ms: s}
}

type MCollection struct {
	*mgo.Collection
	ms *mgo.Session
}

func (m *MCollection) SwitchToCollection(collection string) {
	m.Collection = m.ms.DB(mongo_db).C(collection)
}

func (m *MCollection) Close() {
	m.ms.Close()
	//pool <- m.ms
}

func selector(fields ...string) bson.M {
	f := bson.M{}
	for _, k := range fields {
		f[k] = 1
	}
	return f
}

func setUpdate(up bson.M) bson.M {
	return bson.M{"$set": up}
}
