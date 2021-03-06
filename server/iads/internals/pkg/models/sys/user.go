package sys

import (
	basemodel2 "iads/server/iads/internals/pkg/models/basemodel"
	database2 "iads/server/iads/internals/pkg/models/database"
)

type User struct {
	basemodel2.OrmModel
	Username string `json:"username" gorm:"type:varchar(64);unique_index" `
	Password string `json:"password" gorm:"type:varchar(256)"`
	Email    string `json:"email" gorm:"type:varchar(128)"`
	Role     Role   `json:"role" gorm:"auto_preload;foreignkey:RoleID"`
	RoleID   uint64 `json:"role_id"`
}

func (u *User) UserGetFromName() (user User, err error) {
	if err = database2.DBE.Where("username = ?", u.Username).First(&user).Error; err != nil {
		return
	}
	user.Role.ID = user.RoleID
	var r Role
	database2.DBE.Where("id = ?", user.RoleID).First(&r)
	user.Role.RoleName = r.RoleName
	user.Role.RoleDetails = r.RoleDetails
	user.Role.CreatedAt = r.CreatedAt
	user.Role.UpdatedAt = r.UpdatedAt
	return
}

//列表
func (u *User) UserList() (users []User, err error) {
	//orm.Eloquent.Model(&user).Related(&user.Role).Find(&user.Role)
	if err = database2.DBE.Find(&users).Error; err != nil {
		return
	}
	for i := range users {
		var r Role
		users[i].Role.ID = users[i].RoleID
		database2.DBE.Where("id = ?", users[i].RoleID).First(&r)
		users[i].Role.RoleName = r.RoleName
		users[i].Role.RoleDetails = r.RoleDetails
		users[i].Role.CreatedAt = r.CreatedAt
		users[i].Role.UpdatedAt = r.UpdatedAt
	}
	return
}

//添加
func (u *User) UserInsert() (id uint64, err error) {
	//添加数据
	if 0 == u.RoleID {
		if u.Role, err = u.Role.RoleGetFromName(u.Role.RoleName); err != nil {
			return
		}
		u.RoleID = u.Role.ID
		u.Role = Role{}
	}
	result := database2.DBE.Create(&u)
	id = u.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//修改
func (u *User) UserUpdate(id uint64) (updateUser User, err error) {
	if err = database2.DBE.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
		return
	}
	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = database2.DBE.Model(&updateUser).Updates(&u).Error; err != nil {
		return
	}
	return
}

//删除数据
func (u *User) UserDestroyFromID(id uint64) (Result User, err error) {
	if err = database2.DBE.Select([]string{"id"}).First(&u, id).Error; err != nil {
		return
	}
	if err = database2.DBE.Delete(&u).Error; err != nil {
		return
	}
	Result = *u
	return
}

func (u *User) UserDestroyFromName(userName string) (Result User, err error) {
	if err = database2.DBE.Where("username = ?", userName).First(&u).Error; err != nil {
		return
	}
	if err = database2.DBE.Delete(&u).Error; err != nil {
		return
	}
	Result = *u
	return
}
