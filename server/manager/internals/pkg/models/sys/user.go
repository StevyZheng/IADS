package sys

import (
	"iads/server/manager/internals/pkg/models/database"
)

type User struct {
	BaseModel
	database.MDatabase
	Name     string `bson:"name,omitempty" json:"name"`
	Password string `bson:"password,omitempty" json:"password"`
	Role     Role
}
