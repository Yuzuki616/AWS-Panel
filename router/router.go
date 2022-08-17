package router

import (
	"github.com/Yuzuki616/Aws-Panel/controller"
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

	//User
	p.router.POST("/api/v1/User/Login", controller.LoginVerify)
	p.router.POST("/api/v1/User/Register", controller.Register)
	//p.router.POST("/api/v1/user/SendMailVerify", controller.SendMailVerify)
	p.router.POST("/api/v1/User/ChangeUsername", controller.ChangeUsername)
	p.router.POST("/api/v1/User/ChangePassword", controller.ChangePassword)
	p.router.GET("/api/v1/User/Info", controller.GetUserInfo)
	p.router.GET("/api/v1/User/Logout", controller.Logout)
	p.router.GET("/api/v1/User/IsAdmin", controller.IsAdmin)
	//Admin Only
	p.router.GET("/api/v1/User/List", controller.GetUserList)
	p.router.POST("/api/v1/User/Delete", controller.DeleteUser)
	p.router.POST("/api/v1/User/Ban", controller.BanUser)
	p.router.POST("/api/v1/User/UnBan", controller.UnBanUser)

	//Secret
	p.router.POST("/api/v1/Secret/Add", controller.AddSecret)
	p.router.GET("/api/v1/Secret/List", controller.ListSecret)
	p.router.POST("/api/v1/Secret/Delete", controller.DelSecret)
	p.router.POST("/api/v1/Secret/Info", controller.GetSecretInfo)

	//Ec2
	p.router.POST("/api/v1/Ec2/Create", controller.CreateEc2)
	p.router.POST("/api/v1/Ec2/List", controller.ListEc2)
	p.router.POST("/api/v1/Ec2/Info", controller.GetEc2Info)
	p.router.POST("/api/v1/Ec2/ChangeIp", controller.ChangeEc2Ip)
	p.router.POST("/api/v1/Ec2/Stop", controller.StopEc2)
	p.router.POST("/api/v1/Ec2/Start", controller.StartEc2)
	p.router.POST("/api/v1/Ec2/Reboot", controller.RebootEc2)
	p.router.POST("/api/v1/Ec2/Delete", controller.DeleteEc2)

	//Lightsail
	p.router.POST("/api/v1/LightSail/GetRegions", controller.GetRegions)
	p.router.POST("/api/v1/LightSail/Create", controller.CreateLightsail)
	p.router.POST("/api/v1/LightSail/OpenPorts", controller.OpenLightsailPorts)
	p.router.POST("/api/v1/LightSail/GetBlueprintId")
	p.router.POST("/api/v1/LightSail/Info", controller.GetLightsailInfo)
	p.router.POST("/api/v1/LightSail/List", controller.ListLightsail)
	p.router.POST("/api/v1/LightSail/ChangeIp", controller.ChangeLightsailIp)
	p.router.POST("/api/v1/LightSail/Stop", controller.StopLightsail)
	p.router.POST("/api/v1/LightSail/Start", controller.StartLightsail)
	p.router.POST("/api/v1/LightSail/Reboot", controller.RebootLightsail)
	p.router.POST("/api/v1/LightSail/Delete", controller.DeleteLightsail)

	//Quota
	p.router.POST("/api/v1/Quota/Get")
	p.router.GET("/api/v1/Quota/ListChangeRequest")
	p.router.POST("/api/v1/Quota/ChangeQuota")

	//Page
	p.router.Static("/js", "./web/js")
	p.router.Static("/css", "./web/css")
	p.router.Static("/img", "./web/img")
	//p.router.StaticFile("/","./web/index.html")
	p.router.StaticFile("/favicon.ico", "./web/favicon.ico")
	p.router.NoRoute(func(c *gin.Context) {
		c.File("./web/index.html")
	})
}

func (p *Router) Start() error {
	runErr := p.router.Run(":8011")
	if runErr != nil {
		return runErr
	}
	return nil
}
