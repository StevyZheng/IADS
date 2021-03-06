package sys

import (
	basemodel2 "iads/server/iads/internals/pkg/models/basemodel"
	database2 "iads/server/iads/internals/pkg/models/database"
)

/*.
admin
tenant
ordinary
*/
type Role struct {
	basemodel2.OrmModel
	RoleName    string `json:"role_name" gorm:"type:varchar(32);unique_index;"`
	RoleDetails string `json:"role_details"`
}

func (r Role) RoleGetFromID(id uint64) (role Role, err error) {
	if err = database2.DBE.Where("id = ?", id).First(&role).Error; err != nil {
		return
	}
	return
}

func (r Role) RoleGetFromName(roleName string) (role Role, err error) {
	if err = database2.DBE.Where("role_name = ?", roleName).First(&role).Error; err != nil {
		return
	}
	return
}

func (r Role) RoleList() (roles []Role, err error) {
	if err = database2.DBE.Find(&roles).Error; err != nil {
		return
	}
	return
}

func (r Role) RoleInsert() (id uint64, err error) {
	//添加数据
	if err = database2.DBE.Create(&r).Error; err != nil {
		return
	}
	id = r.ID
	return
}

//修改成r
func (r Role) RoleUpdate(id uint64) (updateRole Role, err error) {
	if err = database2.DBE.Where("id = ?", id).First(&updateRole).Error; err != nil {
		return
	}
	if err = database2.DBE.Model(&updateRole).Updates(&r).Error; err != nil {
		return
	}
	updateRole = r
	return
}

//删除数据
func (r Role) RoleDestroy(id uint64) (role Role, err error) {
	r.ID = id
	if err = database2.DBE.Delete(&r).Error; err != nil {
		return
	}
	role = r
	return
}
