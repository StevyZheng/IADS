package v1

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	database2 "iads/server/iads/internals/pkg/models/database"
	sys2 "iads/server/iads/internals/pkg/models/sys"
	config2 "iads/server/iads/pkg/config"
	jwt2 "iads/server/iads/pkg/jwt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResult struct {
	Token string `json:"token"`
	sys2.User
}

func LoginCheck(info LoginInfo) (flag bool, u sys2.User, err error) {
	var user sys2.User
	if len(info.Username) == 0 || len(info.Password) == 0 {
		return false, user, nil
	}
	err = database2.DBE.Where("username = ?", info.Username).Preload("Role").First(&user).Error
	if err != nil {
		return false, user, err
	}
	if info.Password == user.Password {
		return true, user, nil
	} else {
		return false, user, nil
	}
}

func Login(c *gin.Context) {
	var login LoginInfo
	err := c.ShouldBindJSON(&login)
	if err == nil {
		isPass, user, err := LoginCheck(login)
		if isPass {
			generateToken(c, user)
		} else {
			config2.JsonRequest(c, -1, nil, err)
		}
	} else {
		println(err.Error())
		config2.JsonRequest(c, -3, nil, nil)
	}
}

// 生成令牌
func generateToken(c *gin.Context, user sys2.User) {
	j := &jwt2.JWT{
		SigningKey: []byte("newtrekWang"),
	}
	claims := jwt2.CustomClaims{
		UserID:   user.ID,
		UserName: user.Username,
		Email:    user.Email,
		RoleID:   user.RoleID,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "newtrekWang",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	log.Println(token)

	data := LoginResult{
		User:  user,
		Token: token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   data,
	})
	return
}

// GetDataByTime 一个需要token认证的测试接口
func GetDataByTime(c *gin.Context) {
	claims := c.MustGet("claims").(*jwt2.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}
}

func UserGetFromName(c *gin.Context) {
	var user sys2.User
	userName := c.Param("username")
	user.Username = userName
	role, err := user.UserGetFromName()
	if err != nil {
		config2.JsonRequest(c, -1, nil, err)
		return
	}
	config2.JsonRequest(c, 1, role, err)
}

//列表数据
func UserList(c *gin.Context) {
	var user sys2.User
	result, err := user.UserList()
	if err != nil {
		config2.JsonRequest(c, -2, nil, err)
		return
	}
	config2.JsonRequest(c, 1, result, nil)
}

type UserStoreInfo struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RoleName string `json:"role_name"`
}

//添加用户
func UserCreate(c *gin.Context) {
	var userInfo UserStoreInfo
	err := c.ShouldBindJSON(&userInfo)
	var user sys2.User
	user.Role.RoleName = userInfo.RoleName
	user.Username = userInfo.UserName
	user.Password = userInfo.Password
	user.Email = userInfo.Email
	id, err := user.UserInsert()
	if err != nil {
		config2.JsonRequest(c, -1, nil, err)
		return
	}
	config2.JsonRequest(c, 1, id, nil)
}

type UserUpdateInfo struct {
	UserID int64 `json:"user_id"`
	UserStoreInfo
}

//修改数据
func UserUpdate(c *gin.Context) {
	var user sys2.User
	var userUpdateInfo UserUpdateInfo
	err := c.ShouldBind(&userUpdateInfo)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	user.Password = c.Request.FormValue("password")
	result, err := user.UserUpdate(uint64(id))
	if err != nil || result.ID == 0 {
		config2.JsonRequest(c, -1, nil, err)
		return
	}
	config2.JsonRequest(c, 1, nil, nil)
}

func UserDestroyFromUserName(c *gin.Context) {
	var user sys2.User
	err := c.ShouldBindJSON(&user)
	result, err := user.UserDestroyFromName(user.Username)
	if err != nil || result.ID == 0 {
		config2.JsonRequest(c, -1, nil, err)
		return
	}
	config2.JsonRequest(c, 1, nil, nil)
}

func UserDestroy(c *gin.Context) {
	var user sys2.User
	user.Username = c.Param("username")
	result, err := user.UserDestroyFromName(user.Username)
	if err != nil || result.ID == 0 {
		config2.JsonRequest(c, -1, nil, err)
		return
	}
	config2.JsonRequest(c, 1, nil, nil)
}
