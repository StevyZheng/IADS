package v1

import (
	"github.com/gin-gonic/gin"
	"iads/server/internals/pkg/models/sys"
	"iads/server/pkg/config"
)

func RoleGetFromName(c *gin.Context) {
	var role sys.Role
	roleName := c.Param("role_name")
	role.RoleName = roleName
	role, err := role.RoleGetFromName(roleName)
	if err != nil {
		config.JsonRequest(c, -1, nil, err)
		return
	}
	config.JsonRequest(c, 1, role, err)
}

func RoleList(c *gin.Context) {
	var role sys.Role
	result, err := role.RoleList()
	if err != nil {
		config.JsonRequest(c, -1, nil, err)
		return
	}
	config.JsonRequest(c, 1, result, err)
}

//添加角色
func RoleCreate(c *gin.Context) {
	var role sys.Role
	err := c.ShouldBindJSON(&role)
	//role.RoleName = c.Request.FormValue("rolename")
	id, err := role.RoleInsert()
	if err != nil {
		config.JsonRequest(c, -1, nil, err)
		return
	}
	config.JsonRequest(c, 1, id, nil)
}

func RoleDestroyFromID(c *gin.Context) {
	var role sys.Role
	err := c.ShouldBindJSON(&role.ID)
	//roleId, err := strconv.ParseInt(c.Request.FormValue("role_id"), 10, 64)
	if err != nil {
		config.JsonRequest(c, -1, nil, err)
		return
	}
	_, err = role.RoleDestroy(role.ID)
	if err != nil {
		config.JsonRequest(c, -1, nil, err)
		return
	}
	config.JsonRequest(c, 1, role, nil)
}

func RoleDestroyFromRoleName(c *gin.Context) {
	var role sys.Role
	err := c.ShouldBindJSON(&role.RoleName)
	//role.RoleName = c.Request.FormValue("role_name")
	if role, err = role.RoleGetFromName(role.RoleName); err != nil {
		return
	}
	_, err = role.RoleDestroy(role.ID)
	if err != nil {
		config.JsonRequest(c, -1, nil, err)
		return
	}
	config.JsonRequest(c, 1, role, nil)
}
