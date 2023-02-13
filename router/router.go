package router

import (
	"github.com/Yuzuki616/Aws-Panel/conf"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func Init() {
	gin.SetMode(gin.ReleaseMode)
	engine = gin.Default()
}

func LoadRoute() {
	store := cookie.NewStore([]byte("loginSession"))
	/*store, err := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("loginuser"))
	if err != nil {
		log.Error("Create session error: ", err)
	}*/
	engine.Use(sessions.Sessions("loginSession", store))
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:8080"}
	config.AllowCredentials = true
	engine.Use(cors.New(config))
	LoadNormalRoute()
	LoadUserRoute()
	LoadAdminRoute()
}

func Start() error {
	LoadRoute()
	runErr := engine.Run(conf.Config.Addr)
	if runErr != nil {
		return runErr
	}
	return nil
}
