package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/url"
	"runtime/debug"
	"time"
)

type MongoUtils struct {
	Con      *mongo.Client
	Db       *mongo.Database
	ServerIp string
	Port     int
}

func (o *MongoUtils) OpenConn() (con *mongo.Client) {
	connString := fmt.Sprintf("mongodb://%s:%d", o.ServerIp, o.Port)
	_, err := url.Parse(connString)
	if err != nil {
		println(err)
		return
	}
	opts := options.Client().ApplyURI(connString)
	opts.SetAuth(options.Credential{AuthMechanism: "SCRAM-SHA-1", AuthSource: "AuthDb", Username: "admin", Password: "000000"})
	opts.SetMaxPoolSize(64)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	con, err = mongo.Connect(ctx, opts)
	if err != nil {
		println(err)
		return nil
	}
	err = con.Ping(ctx, readpref.Primary())
	if err != nil {
		println(err)
		return nil
	}
	o.Con = con
	return con
}

func (o *MongoUtils) SetDb(db string) {
	if o.Con == nil {
		panic("Connect  is nil...")
	}
	o.Db = o.Con.Database(db)
}

func (o *MongoUtils) FindOne(col string, filter bson.M) (bson.M, error) {
	if o.Db == nil || o.Con == nil {
		return nil, fmt.Errorf("Not init connect and database!")
	}
	table := o.Db.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var result bson.M
	err := table.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (o *MongoUtils) FindMore(col string, filter bson.M) ([]bson.M, error) {
	if o.Db == nil || o.Con == nil {
		return nil, fmt.Errorf("Not init connect and database!")
	}
	table := o.Db.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	cur, err2 := table.Find(ctx, filter)
	if err2 != nil {
		fmt.Print(err2)
		return nil, err2
	}
	defer cur.Close(ctx)
	var resultArr []bson.M
	for cur.Next(ctx) {
		var result bson.M
		err3 := cur.Decode(&result)
		if err3 != nil {
			return nil, err3
		}
		resultArr = append(resultArr, result)
	}
	return resultArr, nil
}

func Bson2Odj(val interface{}, obj interface{}) (err error) {
	data, err := bson.Marshal(val)
	if err != nil {
		return err
	}
	_ = bson.Unmarshal(data, obj)
	return nil
}

func (o *MongoUtils) InsertOne(col string, elem interface{}) (err error) {
	cols := o.Db.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if _, err := cols.InsertOne(ctx, elem); err != nil && cid != nil {
		return err
	}
	return err
}

func (o *MongoUtils) InsertMany(col string, elemArray []interface{}) (err error) {
	cols := o.Db.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if _, err := cols.InsertMany(ctx, elemArray); err != nil {
		return err
	}
	return err
}
