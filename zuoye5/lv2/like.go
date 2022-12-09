package lv2

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
)

var Rdb *redis.Client //一个获取用户账号密码的redis接口

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}

func Like(ctx context.Context, username string, c *gin.Context) {
	var like string
	if Rdb.SIsMember(ctx, like, username).Val() {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exists",
		})
		return
	} else {
		Rdb.SAdd(ctx, like, username)
		fmt.Println("点赞成功")
	}
}
func UnLike(ctx context.Context, username string, c *gin.Context) {
	var like string
	if Rdb.SIsMember(ctx, like, username).Val() {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exists",
		})
		return
	} else {
		Rdb.SRem(ctx, like, username)
		fmt.Println("取消点赞成功")
	}
}
