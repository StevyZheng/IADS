package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	. "iads/server/manager/internals/pkg/models/sys"
)

func RoleList(c *gin.Context) {
	roles, err := Role{}.List()
	if err != nil {
		JsonResult(c, 400, err, nil)
	} else {
		JsonResult(c, 200, nil, roles)
	}
}

func RoleAddOne(c *gin.Context) {
	var role = Role{}
	if err := c.ShouldBind(&role); err != nil {
		JsonResult(c, 402, err, nil)
	} else {
		err = role.AddOne()
		if err != nil {
			JsonResult(c, 401, err, nil)
		} else {
			JsonResult(c, 200, err, nil)
		}
	}
}

type RoleUpdate struct {
	Before Role `json:"before"`
	After  Role `json:"after"`
}

func RoleUpdateOneFromName(c *gin.Context) {
	var update = RoleUpdate{}
	if err := c.ShouldBind(&update); err != nil {
		JsonResult(c, 402, err, nil)
	} else {
		if err := update.Before.UpdateOneFromName(update.After); err != nil {
			if err == mongo.ErrNoDocuments {
				JsonResult(c, 403, errors.New("role not exist can not update"), nil)
			} else {
				JsonResult(c, 401, err, nil)
			}
		} else {
			JsonResult(c, 200, err, update.Before)
		}
	}
}

func RoleDeleteFromName(c *gin.Context) {
	var role = Role{}
	if err := c.ShouldBind(&role); err != nil {
		JsonResult(c, 402, err, nil)
	} else {
		if err = role.DeleteFromName(); err != nil {
			JsonResult(c, 401, err, nil)
		} else {
			JsonResult(c, 200, nil, nil)
		}
	}
}
