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
	user.POST("/api/v1/user/SendMailVerify", controller.SendMailVerify)
	user.POST("/api/v1/User/ChangeUsername", controller.ChangeUsername)
	user.POST("/api/v1/User/ChangePassword", controller.ChangePassword)
	user.GET("/api/v1/User/Info", controller.GetUserInfo)
	user.GET("/api/v1/User/Logout", controller.Logout)
	user.GET("/api/v1/User/IsAdmin", controller.IsAdmin)
	//Secret
	user.POST("/api/v1/Secret/Add", controller.AddSecret)
	user.GET("/api/v1/Secret/List", controller.ListSecret)
	user.POST("/api/v1/Secret/Delete", controller.DelSecret)
	user.POST("/api/v1/Secret/Info", controller.GetSecretInfo)
	//Ec2
	user.POST("/api/v1/Ec2/Create", controller.CreateEc2)
	user.POST("/api/v1/Ec2/List", controller.ListEc2)
	user.POST("/api/v1/Ec2/Info", controller.GetEc2Info)
	user.POST("/api/v1/Ec2/ChangeIp", controller.ChangeEc2Ip)
	user.POST("/api/v1/Ec2/Stop", controller.StopEc2)
	user.POST("/api/v1/Ec2/Start", controller.StartEc2)
	user.POST("/api/v1/Ec2/Reboot", controller.RebootEc2)
	user.POST("/api/v1/Ec2/Delete", controller.DeleteEc2)
	//Lightsail
	user.POST("/api/v1/LightSail/GetRegions", controller.GetRegions)
	user.POST("/api/v1/LightSail/Create", controller.CreateLightsail)
	user.POST("/api/v1/LightSail/OpenPorts", controller.OpenLightsailPorts)
	user.POST("/api/v1/LightSail/GetBlueprintId")
	user.POST("/api/v1/LightSail/Info", controller.GetLightsailInfo)
	user.POST("/api/v1/LightSail/List", controller.ListLightsail)
	user.POST("/api/v1/LightSail/ChangeIp", controller.ChangeLightsailIp)
	user.POST("/api/v1/LightSail/Stop", controller.StopLightsail)
	user.POST("/api/v1/LightSail/Start", controller.StartLightsail)
	user.POST("/api/v1/LightSail/Reboot", controller.RebootLightsail)
	user.POST("/api/v1/LightSail/Delete", controller.DeleteLightsail)
	//Quota
	user.POST("/api/v1/Quota/Get")
	user.GET("/api/v1/Quota/ListChangeRequest")
	user.POST("/api/v1/Quota/ChangeQuota")
}

func (p *Router) LoadAdminRoute() {
	admin := p.router.Group("/api/v1/")
	admin.Use(middleware.UserCheck, middleware.AdminCheck)
	//Admin Only
	admin.GET("/api/v1/User/List", controller.GetUserList)
	admin.POST("/api/v1/User/Delete", controller.DeleteUser)
	admin.POST("/api/v1/User/Ban", controller.BanUser)
	admin.POST("/api/v1/User/UnBan", controller.UnBanUser)
}
