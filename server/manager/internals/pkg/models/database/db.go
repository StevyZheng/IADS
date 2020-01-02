package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/url"
	_ "runtime/debug"
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

func Bson2Odj(val interface{}, obj interface{}) (err error) {
	data, err := bson.Marshal(val)
	if err != nil {
		return err
	}
	_ = bson.Unmarshal(data, obj)
	return nil
}

func (o *MongoUtils) CountDoc(col string) (size int64, err error) {
	if o.Db == nil || o.Con == nil {
		return 0, fmt.Errorf("Not init connect and database!")
	}
	table := o.Db.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if size, err = table.CountDocuments(ctx, bson.D{}); err != nil {
		return 0, err
	}
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

func (o *MongoUtils) FindOneDelete(col string, filter bson.M) (bson.M, error) {
	if o.Db == nil || o.Con == nil {
		return nil, fmt.Errorf("Not init connect and database!")
	}
	table := o.Db.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var result bson.M
	err := table.FindOneAndDelete(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//查询单条数据后修改该数据
func (o *MongoUtils) FindOneUpdate(col string, filter bson.M, update bson.M) (bson.M, error) {
	if o.Db == nil || o.Con == nil {
		return nil, fmt.Errorf("Not init connect and database!")
	}
	table := o.Db.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var result bson.M
	err := table.FindOneAndUpdate(ctx, filter, update).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//查询单条数据后替换该数据(以前的数据全部清空)
func (o *MongoUtils) FindOneReplace(col string, filter bson.M, replace bson.M) (bson.M, error) {
	if o.Db == nil || o.Con == nil {
		return nil, fmt.Errorf("Not init connect and database!")
	}
	table := o.Db.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var result bson.M
	err := table.FindOneAndUpdate(ctx, filter, replace).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (o *MongoUtils) FindMore(col string, filter bson.M, opts ...*options.FindOptions) ([]bson.M, error) {
	if o.Db == nil || o.Con == nil {
		return nil, fmt.Errorf("Not init connect and database!")
	}
	table := o.Db.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	cur, err2 := table.Find(ctx, filter, opts...)
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

func (o *MongoUtils) InsertOne(col string, elem interface{}) (err error) {
	table := o.Db.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if _, err := table.InsertOne(ctx, elem); err != nil {
		return err
	}
	return err
}

func (o *MongoUtils) InsertMany(col string, elemArray []interface{}) (err error) {
	table := o.Db.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if _, err := table.InsertMany(ctx, elemArray); err != nil {
		return err
	}
	return err
}

func (o *MongoUtils) UpdateOne(col string, filter bson.M, update bson.M) (err error) {
	if o.Db == nil || o.Con == nil {
		return fmt.Errorf("Not init connect and database!")
	}
	table := o.Db.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = table.UpdateOne(ctx, filter, update)
	return err
}

func (o *MongoUtils) UpdateMany(col string, filter bson.M, update bson.M) (err error) {
	if o.Db == nil || o.Con == nil {
		return fmt.Errorf("Not init connect and database!")
	}
	table := o.Db.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = table.UpdateMany(ctx, filter, update)
	return err
}

func (o *MongoUtils) DeleteOne(col string, filter bson.M) (err error) {
	if o.Db == nil || o.Con == nil {
		return fmt.Errorf("Not init connect and database!")
	}
	table := o.Db.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = table.DeleteOne(ctx, filter)
	return err
}

func (o *MongoUtils) DeleteMany(col string, filter bson.M) (err error) {
	if o.Db == nil || o.Con == nil {
		return fmt.Errorf("Not init connect and database!")
	}
	table := o.Db.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = table.DeleteMany(ctx, filter)
	return err
}
