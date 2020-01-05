package v1

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"iads/server/manager/internals/pkg/models/database"
	. "iads/server/manager/internals/pkg/models/sys"
)

func RoleList(c *gin.Context) {
	db, err := database.NewMDBDefault()
	if err != nil {
		JsonResult(c, 400, err, nil)
	} else {
		result, err := db.FindMore("role", bson.M{})
		if err != nil {
			JsonResult(c, 401, err, nil)
		}
		JsonResult(c, 200, nil, result)
	}
}

func RoleAddOne(c *gin.Context) {
	db, err := database.NewMDBDefault()
	if err != nil {
		JsonResult(c, 400, err, nil)
	}
	var role = Role{}
	if err = c.ShouldBind(&role); err != nil {
		JsonResult(c, 402, err, nil)
	} else {
		if err = db.InsertOne("role", role); err != nil {
			JsonResult(c, 401, err, nil)
		} else {
			JsonResult(c, 200, nil, nil)
		}
	}
}

type RoleUpdate struct {
	Before Role `json:"before"`
	After  Role `json:"after"`
}

func RoleUpdateFromName(c *gin.Context) {
	db, err := database.NewMDBDefault()
	if err != nil {
		JsonResult(c, 400, err, nil)
	}
	var update = RoleUpdate{}
	if err = c.ShouldBind(&update); err != nil {
		JsonResult(c, 402, err, nil)
	} else {
		afterBson, err := database.ObjToBson(update.After)
		if err = db.UpdateOne("role", bson.M{"name": update.Before.Name}, afterBson); err != nil {
			JsonResult(c, 401, err, nil)
		} else {
			JsonResult(c, 200, nil, nil)
		}
	}
}

func RoleDeleteFromName(c *gin.Context) {
	db, err := database.NewMDBDefault()
	if err != nil {
		JsonResult(c, 400, err, nil)
	}
	var role = Role{}
	if err = c.ShouldBind(&role); err != nil {
		JsonResult(c, 402, err, nil)
	} else {
		if err = db.DeleteMany("role", bson.M{"name": role.Name}); err != nil {
			JsonResult(c, 401, err, nil)
		} else {
			JsonResult(c, 200, nil, nil)
		}
	}
}
