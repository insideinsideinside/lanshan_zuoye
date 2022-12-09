package lv1_lv3

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"time"
)

var Rdb *redis.Client //一个获取用户账号密码的redis接口

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}
func GetPassword(ctx context.Context, username string) (string, error) {
	GetKey := Rdb.Get(ctx, username)
	if GetKey.Err() != nil {
		return "", GetKey.Err()
	}
	return GetKey.Val(), nil
}
func CheckPassword(ctx context.Context, username string, password string) bool {
	GetKey := Rdb.Get(ctx, username)
	if GetKey.Val() == password {
		return true
	} else {
		return false
	}
}

func SetPassWord(ctx context.Context, username string, password string, expiration time.Duration) error {
	SetKV := Rdb.Set(ctx, username, password, expiration)
	return SetKV.Err()
} // 以下是router包的内容，方便起见我写在一起

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag, _ := GetPassword(c, username)
	if flag != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user already exists",
		})
		return
	} else {
		SetPassWord(c, username, password, -1)
		SetSql(username, password) //写入数据库，略
	}

}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag, _ := GetPassword(c, username)
	if flag == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exists",
		})
		return
	} else {
		if !CheckPassword(c, username, password) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": "wrong password",
			})
			return
		} else {
			c.SetCookie("gin_demo_cookie", "test", 3600, "/", "localhost", false, true)
			c.JSON(http.StatusOK, gin.H{
				"status":  200,
				"message": "login successful",
			})
		}
	}
}

func ChangePassword(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	newPassword := c.PostForm("newPassword")
	flag, _ := GetPassword(c, username)
	if flag == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exists",
		})
		return
	} else {
		if !CheckPassword(c, username, password) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": "wrong password",
			})
			return
		} else {
			ChangeSql(username, password)             //更新数据库信息，略
			Rdb.Del(c, username, password)            //删除缓存中的键值对
			SetPassWord(c, username, newPassword, -1) //重新写入到redis

		}
	}
}
