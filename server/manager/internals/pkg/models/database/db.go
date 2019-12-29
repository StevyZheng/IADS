package database

import (
	"github.com/globalsign/mgo"
	"log"
	"time"
)

var (
	Session *mgo.Session
)

const (
	host   = "127.0.0.1"
	source = "manager"
	user   = "admin"
	pass   = "000000"
)

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{host},
		Direct:    false,
		Timeout:   time.Second * 1,
		Source:    source,
		Username:  user,
		Password:  pass,
		PoolLimit: 1024,
	}
	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalf("Can't connect to mongo, go error %s\n", err)
	}
	Session = s
}

func connect(db, collation string) (*mgo.Session, *mgo.Collection) {
	s := Session.Copy()
	c := s.DB(db).C(collation)
	s.SetMode(mgo.Monotonic, true)
	return s, c
}

func Insert(db, collection string, docs ...interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Insert(docs)
}

func IsExist(db, collection string, query interface{}) bool {
	ms, c := connect(db, collection)
	defer ms.Close()
	count, _ := c.Find(query).Count()
	return count > 0
}

func IsEmpty(db, collection string) bool {
	ms, c := connect(db, collection)
	defer ms.Close()
	count, err := c.Count()
	if err != nil {
		log.Fatal(err)
	}
	return count == 0
}

func Count(db, collection string, query interface{}) (int, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Count()
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

func FindPage(db, collection string, page, limit int, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).Skip(page * limit).Limit(limit).All(result)
}

func Update(db, collection string, query, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Update(query, update)
}

func Remove(db, collection string, query interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Remove(query)
}

func RemoveAll(db, collection string, query interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	_, err := c.RemoveAll(query)
	return err
}
