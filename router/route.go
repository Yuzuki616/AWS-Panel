package router

import (
	"github.com/Yuzuki616/Aws-Panel/conf"
	"github.com/Yuzuki616/Aws-Panel/controller"
	"github.com/Yuzuki616/Aws-Panel/middleware"
	"github.com/gin-gonic/gin"
	"path"
)

func LoadNormalRoute() {
	if conf.Config.EnableLoadStatic {
		//Page
		engine.Static("/js", path.Join(conf.Config.StaticPath, "js"))
		engine.Static("/css", path.Join(conf.Config.StaticPath, "css"))
		engine.Static("/img", path.Join(conf.Config.StaticPath, "img"))
		//engine.StaticFile("/","./web/index.html")
		engine.StaticFile("/favicon.ico", path.Join(conf.Config.StaticPath, "favicon.ico"))
		engine.NoRoute(func(c *gin.Context) {
			c.File(path.Join(conf.Config.StaticPath, "index.html"))
		})
	}
	engine.GET("/api/v1/Config/IsEnableEmailVerify", controller.IsEnableEmailVerify)
	engine.POST("/api/v1/User/SendMailVerify", controller.SendMailVerify)
	engine.POST("/api/v1/User/Login", controller.Login)
	engine.POST("/api/v1/User/Register", controller.Register)
}

func LoadUserRoute() {
	user := engine.Group("/api/v1/")
	user.Use(middleware.UserCheck)
	user.POST("User/ChangeEmail", controller.ChangeEmail)
	user.POST("User/ChangePassword", controller.ChangePassword)
	user.GET("User/Info", controller.GetUserInfo)
	user.GET("User/Logout", controller.Logout)
	user.GET("User/IsAdmin", controller.IsAdmin)
	//Secret
	user.POST("Secret/Add", controller.AddSecret)
	user.GET("Secret/List", controller.ListSecret)
	user.POST("Secret/Delete", controller.DelSecret)
	user.POST("Secret/Info", controller.GetSecret)
	//Ec2
	user.POST("Ec2/Create", controller.CreateEc2)
	user.POST("Ec2/List", controller.ListEc2)
	user.POST("Ec2/Info", controller.GetEc2Info)
	user.POST("Ec2/ChangeIp", controller.ChangeEc2Ip)
	user.POST("Ec2/Stop", controller.StopEc2)
	user.POST("Ec2/Start", controller.StartEc2)
	user.POST("Ec2/Reboot", controller.RebootEc2)
	user.POST("Ec2/Delete", controller.DeleteEc2)
	user.POST("Ec2/CreateSshKey", controller.CreateEc2SshKey)
	user.POST("Ec2/ListSshKey", controller.ListEc2SshKey)
	user.POST("Ec2/DeleteSshKey", controller.DeleteEc2SshKey)
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

func LoadAdminRoute() {
	admin := engine.Group("/api/v1/")
	admin.Use(middleware.UserCheck, middleware.AdminCheck)
	//Admin Only
	admin.GET("User/List", controller.GetUserList)
	admin.POST("User/Delete", controller.DeleteUser)
	admin.POST("User/Ban", controller.BanUser)
	admin.POST("User/UnBan", controller.UnBanUser)
}
