package sys

import (
	"time"
)

type BaseModel struct {
	Created time.Time `json:"created" bson:"created,omitempty"`
	Updated time.Time `json:"updated" bson:"updated,omitempty"`
	Deleted time.Time `json:"deleted" bson:"deleted,omitempty"`
}
