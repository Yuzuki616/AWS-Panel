package router

import (
	"github.com/Yuzuki616/Aws-Panel/conf"
	"github.com/Yuzuki616/Aws-Panel/controller"
	"github.com/Yuzuki616/Aws-Panel/middleware"
	"github.com/gin-gonic/gin"
	"path"
)

func (p *Router) LoadNormalRoute() {
	if conf.Config.EnableLoadStatic {
		//Page
		p.router.Static("/js", path.Join(conf.Config.StaticPath, "js"))
		p.router.Static("/css", path.Join(conf.Config.StaticPath, "css"))
		p.router.Static("/img", path.Join(conf.Config.StaticPath, "img"))
		//p.router.StaticFile("/","./web/index.html")
		p.router.StaticFile("/favicon.ico", path.Join(conf.Config.StaticPath, "favicon.ico"))
		p.router.NoRoute(func(c *gin.Context) {
			c.File(path.Join(conf.Config.StaticPath, "index.html"))
		})
	}
	p.router.GET("/api/v1/Config/IsEnableEmailVerify", controller.IsEnableEmailVerify)
	p.router.POST("/api/v1/User/Login", controller.LoginVerify)
	p.router.POST("/api/v1/User/Register", controller.Register)
}

func (p *Router) LoadUserRoute() {
	user := p.router.Group("/api/v1/")
	user.Use(middleware.UserCheck)
	user.POST("User/SendMailVerify", controller.SendMailVerify)
	user.POST("User/ChangeUsername", controller.ChangeUsername)
	user.POST("User/ChangePassword", controller.ChangePassword)
	user.GET("User/Info", controller.GetUserInfo)
	user.GET("User/Logout", controller.Logout)
	user.GET("User/IsAdmin", controller.IsAdmin)
	//Secret
	user.POST("Secret/Add", controller.AddSecret)
	user.GET("Secret/List", controller.ListSecret)
	user.POST("Secret/Delete", controller.DelSecret)
	user.POST("Secret/Info", controller.GetSecretInfo)
	//Ec2
	user.POST("Ec2/Create", controller.CreateEc2)
	user.POST("Ec2/List", controller.ListEc2)
	user.POST("Ec2/Info", controller.GetEc2Info)
	user.POST("Ec2/ChangeIp", controller.ChangeEc2Ip)
	user.POST("Ec2/Stop", controller.StopEc2)
	user.POST("Ec2/Start", controller.StartEc2)
	user.POST("Ec2/Reboot", controller.RebootEc2)
	user.POST("Ec2/Delete", controller.DeleteEc2)
	//Lightsail
	user.POST("LightSail/GetRegions", controller.GetRegions)
	user.POST("LightSail/Create", controller.CreateLightsail)
	user.POST("LightSail/OpenPorts", controller.OpenLightsailPorts)
	user.POST("LightSail/GetBlueprintId")
	user.POST("LightSail/Info", controller.GetLightsailInfo)
	user.POST("LightSail/List", controller.ListLightsail)
	user.POST("LightSail/ChangeIp", controller.ChangeLightsailIp)
	user.POST("LightSail/Stop", controller.StopLightsail)
	user.POST("LightSail/Start", controller.StartLightsail)
	user.POST("LightSail/Reboot", controller.RebootLightsail)
	user.POST("LightSail/Delete", controller.DeleteLightsail)
	//Quota
	user.POST("Quota/Get")
	user.GET("Quota/ListChangeRequest")
	user.POST("Quota/ChangeQuota")
}

func (p *Router) LoadAdminRoute() {
	admin := p.router.Group("/api/v1/")
	admin.Use(middleware.UserCheck, middleware.AdminCheck)
	//Admin Only
	admin.GET("User/List", controller.GetUserList)
	admin.POST("User/Delete", controller.DeleteUser)
	admin.POST("User/Ban", controller.BanUser)
	admin.POST("User/UnBan", controller.UnBanUser)
}
