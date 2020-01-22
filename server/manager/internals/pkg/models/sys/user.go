package sys

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"iads/server/manager/internals/pkg/models/database"
	"time"
)

type User struct {
	BaseModel
	Name     string `json:"name" bson:"name,omitempty"`
	Password string `json:"password" bson:"password,omitempty"`
	Email    string `json:"email" bson:"email,omitempty"`
	RoleName string `json:"roleName" bson:"roleName,omitempty"`
}

func (u User) List() (users []User, err error) {
	mdb, err := database.NewMDBDefault()
	if err != nil {
		return nil, err
	}
	table := mdb.DB.Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if cur, err := table.Find(ctx, bson.M{}); err != nil {
		return nil, err
	} else {
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var user User
			err3 := cur.Decode(&user)
			if err3 != nil {
				return nil, err3
			}
			users = append(users, user)
		}
	}
	return users, err
}

func (u User) FindOne() (user User, err error) {
	if mdb, err := database.NewMDBDefault(); err != nil {
		return user, err
	} else {
		table := mdb.DB.Collection("user")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		err = table.FindOne(ctx, bson.M{"name": u.Name}).Decode(&user)
		return user, err
	}
}

func (u *User) AddOne() (err error) {
	if mdb, err := database.NewMDBDefault(); err != nil {
		return err
	} else {
		table := mdb.DB.Collection("user")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		var user User
		if err = table.FindOne(ctx, bson.M{"name": u.Name}).Decode(&user); err != nil {
			if err == mongo.ErrNoDocuments {
				if _, err = table.InsertOne(ctx, u); err != nil {
					return err
				}
				return nil
			} else {
				return errors.New("role is exist")
			}
		}
		return nil
	}
}

func (u *User) UpdateOneFromName(update User) (err error) {
	if mdb, err := database.NewMDBDefault(); err != nil {
		return err
	} else {
		table := mdb.DB.Collection("user")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		if doc, err := database.ObjToBson(update); err != nil {
			return err
		} else {
			err = table.FindOneAndUpdate(ctx, bson.M{"name": u.Name}, bson.M{"$set": doc}).Decode(&u)
			return err
		}
	}
}

func (u *User) DeleteFromName() (err error) {
	if mdb, err := database.NewMDBDefault(); err != nil {
		return err
	} else {
		table := mdb.DB.Collection("user")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		_, err = table.DeleteMany(ctx, bson.M{"name": u.Name})
		return err
	}
}
