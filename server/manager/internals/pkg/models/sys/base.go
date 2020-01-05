package sys

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BaseModel struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Created time.Time          `json:"created" bson:"created,omitempty"`
	Updated time.Time          `json:"updated" bson:"updated,omitempty"`
	Deleted time.Time          `json:"deleted" bson:"deleted,omitempty"`
}
