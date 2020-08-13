package tool

import (
	"github.com/gin-contrib/sessions/redis"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//初始化session操作
func InitSession(engine *gin.Engine) {
	config := GetConfig().RedisConfig
	store, err := redis.NewStore(10, "tcp", config.Addr+":"+config.Port, "", []byte("secret"))
	if err != nil {
		fmt.Println(err.Error())
	}
	engine.Use(sessions.Sessions("mysession", store))
}

// set操作
func SetSess(context *gin.Context, key interface{}, value interface{}) error {
	session := sessions.Default(context)
	if session == nil {
		return nil
	}
	session.Set(key, value)
	return session.Save()
}

// get操作
func GetSess(context *gin.Context, key interface{}) interface{} {
	session := sessions.Default(context)
	return session.Get(key)
}
