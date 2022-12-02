package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
	"zuoye3/api/middleware"
	"zuoye3/dao"
	"zuoye3/model"
	"zuoye3/utils"
)

func register(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.Responsefail(c, "verification failed")
	}
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	flag := dao.SelectUser(username)
	if flag {
		utils.Responsefail(c, "user already exists")
		return
	}
	dao.AddUser(username, password)
	utils.Responsesucess(c, "register success")
}

func login(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.Responsefail(c, "verification failed")
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag := dao.SelectUser(username)
	if !flag {
		utils.Responsefail(c, "user doesn't exists")
		return
	}
	selectpassword := dao.SelectPassword(username)
	if selectpassword != password {
		utils.Responsefail(c, "wrong password")
		return
	}
	claim := model.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "Wzy",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.Responsesucess(c, tokenString)
}
func Changepassword(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	newpassword := c.PostForm("newpassword")
	flag := dao.SelectUser(username)
	if !flag {
		utils.Responsefail(c, "user doesn't exists")
		return
	}
	if password != dao.SelectPassword(username) {
		utils.Responsefail(c, "wrong password")
		return
	}
	u := model.User{
		Username: username,
		Password: newpassword,
		ID:       dao.SelectID(username),
	}
	dao.ChangePassword(u)
	flag1 := dao.SelectUser(username)
	if flag1 {
		utils.Responsesucess(c, "password change success")
		return
	}
}
func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.Responsesucess(c, username.(string))
}

func FindPassword(c *gin.Context) {
	username := c.PostForm("username")
	flag := dao.SelectUser(username)
	if !flag {
		utils.Responsefail(c, "user doesn't exists")
		return
	}
	utils.Responsesucess(c, dao.SelectPassword(username))
}
