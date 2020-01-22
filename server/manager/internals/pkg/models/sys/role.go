package sys

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"iads/server/manager/internals/pkg/models/database"
	"time"
)

type Role struct {
	BaseModel
	Name    string `json:"name" bson:"name,omitempty"`
	Explain string `json:"explain" bson:"explain,omitempty"`
}

func (r Role) List() (roles []Role, err error) {
	mdb, err := database.NewMDBDefault()
	if err != nil {
		return nil, err
	}
	table := mdb.DB.Collection("role")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if cur, err := table.Find(ctx, bson.M{}); err != nil {
		return nil, err
	} else {
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var role Role
			err3 := cur.Decode(&role)
			if err3 != nil {
				return nil, err3
			}
			roles = append(roles, role)
		}
	}
	return roles, err
}

func (r Role) FindOne() (role Role, err error) {
	if mdb, err := database.NewMDBDefault(); err != nil {
		return role, err
	} else {
		table := mdb.DB.Collection("role")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		err = table.FindOne(ctx, bson.M{"name": r.Name}).Decode(&role)
		return role, err
	}
}

func (r *Role) AddOne() (err error) {
	if mdb, err := database.NewMDBDefault(); err != nil {
		return err
	} else {
		table := mdb.DB.Collection("role")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		var role Role
		if err = table.FindOne(ctx, bson.M{"name": r.Name}).Decode(&role); err != nil {
			if err == mongo.ErrNoDocuments {
				if _, err = table.InsertOne(ctx, r); err != nil {
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

func (r *Role) UpdateOneFromName(update Role) (err error) {
	if mdb, err := database.NewMDBDefault(); err != nil {
		return err
	} else {
		table := mdb.DB.Collection("role")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		if doc, err := database.ObjToBson(update); err != nil {
			return err
		} else {
			err = table.FindOneAndUpdate(ctx, bson.M{"name": r.Name}, bson.M{"$set": doc}).Decode(&r)
			return err
		}
	}
}

func (r *Role) DeleteFromName() (err error) {
	if mdb, err := database.NewMDBDefault(); err != nil {
		return err
	} else {
		table := mdb.DB.Collection("role")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		_, err = table.DeleteMany(ctx, bson.M{"name": r.Name})
		return err
	}
}
