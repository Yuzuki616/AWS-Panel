package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

func New() *Router {
	gin.SetMode(gin.ReleaseMode)
	return &Router{router: gin.Default()}
}

func (p *Router) LoadRoute() {
	store := cookie.NewStore([]byte("loginSession"))
	/*store, err := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("loginuser"))
	if err != nil {
		log.Error("Create session error: ", err)
	}*/
	p.router.Use(sessions.Sessions("loginSession", store))
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:8080"}
	config.AllowCredentials = true
	p.router.Use(cors.New(config))
	p.LoadNormalRoute()
	p.LoadUserRoute()
	p.LoadAdminRoute()
}

func (p *Router) Start() error {
	runErr := p.router.Run(":8011")
	if runErr != nil {
		return runErr
	}
	return nil
}
