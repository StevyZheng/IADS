package v1

import (
	"errors"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	. "iads/server/manager/internals/pkg/models/sys"
	"iads/server/manager/pkg/jwt"
	"log"
	"net/http"
	"time"
)

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResult struct {
	Token string `json:"token"`
	User
}

func LoginCheck(info LoginInfo) (flag bool, u User, err error) {
	var user User
	if len(info.Username) == 0 || len(info.Password) == 0 {
		return false, user, nil
	}
	user.Name = info.Username
	userFind, err := user.FindOne()
	if err != nil {
		return false, userFind, err
	} else {
		if info.Password == userFind.Password {
			return true, userFind, nil
		} else {
			return false, userFind, nil
		}
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
			JsonResult(c, 405, err, nil)
		}
	} else {
		println(err.Error())
		JsonResult(c, 406, err, nil)
	}
}

// 生成令牌
func generateToken(c *gin.Context, user User) {
	j := &jwt.JWT{
		SigningKey: []byte("newtrekWang"),
	}
	claims := jwt.CustomClaims{
		Username: user.Name,
		//RoleName:   user.RoleName,
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
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}
}

func UserList(c *gin.Context) {
	users, err := User{}.List()
	if err != nil {
		JsonResult(c, 400, err, nil)
	} else {
		JsonResult(c, 200, nil, users)
	}
}

func UserAddOne(c *gin.Context) {
	var user = User{}
	if err := c.ShouldBind(&user); err != nil {
		JsonResult(c, 402, err, nil)
	} else {
		err = user.AddOne()
		if err != nil {
			JsonResult(c, 401, err, nil)
		} else {
			JsonResult(c, 200, err, nil)
		}
	}
}

type UserUpdate struct {
	Before User `json:"before"`
	After  User `json:"after"`
}

func UserUpdateOneFromName(c *gin.Context) {
	var update = UserUpdate{}
	if err := c.ShouldBind(&update); err != nil {
		JsonResult(c, 402, err, nil)
	} else {
		if err = update.Before.UpdateOneFromName(update.After); err != nil {
			if err == mongo.ErrNoDocuments {
				JsonResult(c, 403, errors.New("user not exist can not update"), nil)
			} else {
				JsonResult(c, 401, err, nil)
			}
		} else {
			JsonResult(c, 200, err, update.Before)
		}
	}
}

func UserDeleteFromName(c *gin.Context) {
	var user = User{}
	if err := c.ShouldBind(&user); err != nil {
		JsonResult(c, 402, err, nil)
	} else {
		if err = user.DeleteFromName(); err != nil {
			JsonResult(c, 401, err, nil)
		} else {
			JsonResult(c, 200, nil, nil)
		}
	}
}
