package database

import (
	"context"
	"errors"
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

type MDatabase struct {
	Client  *mongo.Client
	DB      *mongo.Database
	Context context.Context
}

func NewMDBDefault() (*MDatabase, error) {
	return NewMDB("127.0.0.1", 27017, "manager")
}

func NewMDB(serverIP string, port int, dbName string) (*MDatabase, error) {
	connString := fmt.Sprintf("mongodb://%s:%d", serverIP, port)
	_, err := url.Parse(connString)
	if err != nil {
		println(err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := options.Client().ApplyURI(connString)
	opts.SetAuth(options.Credential{AuthMechanism: "SCRAM-SHA-1", AuthSource: "manager", Username: "admin", Password: "000000"})
	opts.SetMaxPoolSize(64)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	ctxPing, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctxPing, readpref.Primary())
	if err != nil {
		println(err)
		return nil, err
	}
	db := client.Database(dbName)
	return &MDatabase{DB: db, Client: client, Context: ctx}, nil
}

func (d *MDatabase) Close() {
	_ = d.Client.Disconnect(d.Context)
}

func BsonToOdj(val interface{}, obj interface{}) (err error) {
	data, err := bson.Marshal(val)
	if err != nil {
		return err
	}
	_ = bson.Unmarshal(data, obj)
	return nil
}

func ObjToBson(obj interface{}) (b bson.M, err error) {
	data, err := bson.Marshal(&obj)
	if err != nil {
		return nil, err
	}
	b = bson.M{}
	err = bson.Unmarshal(data, b)
	if err != nil {
		return nil, err
	}
	return b, err
}

func (o *MDatabase) CountDoc(col string) (size int64, err error) {
	if o.DB == nil || o.Client == nil {
		return 0, errors.New("not init connect and database")
	}
	table := o.DB.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if size, err = table.CountDocuments(ctx, bson.D{}); err != nil {
		return 0, err
	}
	return size, err
}

func (o *MDatabase) FindOne(col string, filter bson.M) (bson.M, error) {
	if o.DB == nil || o.Client == nil {
		return nil, errors.New("not init connect and database")
	}
	table := o.DB.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var result bson.M
	err := table.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (o *MDatabase) FindOneDelete(col string, filter bson.M) (bson.M, error) {
	if o.DB == nil || o.Client == nil {
		return nil, errors.New("not init connect and database")
	}
	table := o.DB.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var result bson.M
	err := table.FindOneAndDelete(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//查询单条数据后修改该数据
func (o *MDatabase) FindOneUpdate(col string, filter bson.M, update bson.M) (bson.M, error) {
	if o.DB == nil || o.Client == nil {
		return nil, errors.New("not init connect and database")
	}
	table := o.DB.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var result bson.M
	err := table.FindOneAndUpdate(ctx, filter, update).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//查询单条数据后替换该数据(以前的数据全部清空)
func (o *MDatabase) FindOneReplace(col string, filter bson.M, replace bson.M) (bson.M, error) {
	if o.DB == nil || o.Client == nil {
		return nil, errors.New("not init connect and database")
	}
	table := o.DB.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var result bson.M
	err := table.FindOneAndUpdate(ctx, filter, replace).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (o *MDatabase) FindMore(col string, filter bson.M, opts ...*options.FindOptions) ([]bson.M, error) {
	if o.DB == nil || o.Client == nil {
		return nil, errors.New("not init connect and database")
	}
	table := o.DB.Collection(col)
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

func (o *MDatabase) InsertOne(col string, elem interface{}) (err error) {
	table := o.DB.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if _, err := table.InsertOne(ctx, elem); err != nil {
		return err
	}
	return err
}

func (o *MDatabase) InsertMany(col string, elemArray []interface{}) (err error) {
	table := o.DB.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if _, err := table.InsertMany(ctx, elemArray); err != nil {
		return err
	}
	return err
}

func (o *MDatabase) UpdateOne(col string, filter bson.M, update bson.M) (err error) {
	if o.DB == nil || o.Client == nil {
		return errors.New("not init connect and database")
	}
	table := o.DB.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = table.UpdateOne(ctx, filter, update)
	return err
}

func (o *MDatabase) UpdateMany(col string, filter bson.M, update bson.M) (err error) {
	if o.DB == nil || o.Client == nil {
		return errors.New("not init connect and database")
	}
	table := o.DB.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = table.UpdateMany(ctx, filter, update)
	return err
}

func (o *MDatabase) DeleteOne(col string, filter bson.M) (err error) {
	if o.DB == nil || o.Client == nil {
		return errors.New("not init connect and database")
	}
	table := o.DB.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = table.DeleteOne(ctx, filter)
	return err
}

func (o *MDatabase) DeleteMany(col string, filter bson.M) (err error) {
	if o.DB == nil || o.Client == nil {
		return errors.New("not init connect and database")
	}
	table := o.DB.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = table.DeleteMany(ctx, filter)
	return err
}
