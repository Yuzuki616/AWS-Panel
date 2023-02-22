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
	engine.POST("/api/v1/User/SendMailVerify", controller.ErrorHandle(controller.SendMailVerify))
	engine.POST("/api/v1/User/Login", controller.ErrorHandle(controller.Login))
	engine.POST("/api/v1/User/Register", controller.ErrorHandle(controller.Register))
}

func LoadUserRoute() {
	user := engine.Group("/api/v1/")
	user.Use(middleware.UserCheck)
	user.POST("User/ChangeEmail", controller.ErrorHandle(controller.ChangeEmail))
	user.POST("User/ChangePassword", controller.ErrorHandle(controller.ChangePassword))
	user.GET("User/Info", controller.GetUserInfo)
	user.GET("User/Logout", controller.ErrorHandle(controller.Logout))
	user.GET("User/IsAdmin", controller.IsAdmin)
	//Secret
	user.POST("Secret/Add", controller.ErrorHandle(controller.AddSecret))
	user.GET("Secret/List", controller.ErrorHandle(controller.ListSecret))
	user.POST("Secret/Delete", controller.ErrorHandle(controller.DelSecret))
	user.POST("Secret/Info", controller.ErrorHandle(controller.GetSecret))
	//Ec2
	user.POST("Ec2/Create", controller.ErrorHandle(controller.CreateEc2))
	user.POST("Ec2/List", controller.ErrorHandle(controller.ListEc2))
	user.POST("Ec2/Info", controller.ErrorHandle(controller.GetEc2Info))
	user.POST("Ec2/ChangeIp", controller.ErrorHandle(controller.ChangeEc2Ip))
	user.POST("Ec2/Stop", controller.ErrorHandle(controller.StopEc2))
	user.POST("Ec2/Start", controller.ErrorHandle(controller.StartEc2))
	user.POST("Ec2/Reboot", controller.ErrorHandle(controller.RebootEc2))
	user.POST("Ec2/Delete", controller.ErrorHandle(controller.DeleteEc2))
	user.POST("Ec2/CreateSshKey", controller.ErrorHandle(controller.CreateEc2SshKey))
	user.POST("Ec2/ListSshKey", controller.ErrorHandle(controller.ListEc2SshKey))
	user.POST("Ec2/DeleteSshKey", controller.ErrorHandle(controller.DeleteEc2SshKey))
	//Lightsail
	user.POST("LightSail/GetRegions", controller.ErrorHandle(controller.GetRegions))
	user.POST("LightSail/Create", controller.ErrorHandle(controller.CreateLightsail))
	user.POST("LightSail/OpenPorts", controller.ErrorHandle(controller.OpenLightsailPorts))
	user.POST("LightSail/GetBlueprintId")
	user.POST("LightSail/Info", controller.ErrorHandle(controller.GetLightsailInfo))
	user.POST("LightSail/List", controller.ErrorHandle(controller.ListLightsail))
	user.POST("LightSail/ChangeIp", controller.ErrorHandle(controller.ChangeLightsailIp))
	user.POST("LightSail/Stop", controller.ErrorHandle(controller.StopLightsail))
	user.POST("LightSail/Start", controller.ErrorHandle(controller.StartLightsail))
	user.POST("LightSail/Reboot", controller.ErrorHandle(controller.RebootLightsail))
	user.POST("LightSail/Delete", controller.ErrorHandle(controller.DeleteLightsail))
	//Quota
	user.POST("Quota/Get")
	user.GET("Quota/ListChangeRequest")
	user.POST("Quota/ChangeQuota")
}

func LoadAdminRoute() {
	admin := engine.Group("/api/v1/")
	admin.Use(middleware.UserCheck, middleware.AdminCheck)
	//Admin Only
	admin.GET("User/List", controller.ErrorHandle(controller.GetUserList))
	admin.POST("User/Delete", controller.ErrorHandle(controller.DeleteUser))
	admin.POST("User/Ban", controller.ErrorHandle(controller.BanUser))
	admin.POST("User/UnBan", controller.ErrorHandle(controller.UnBanUser))
}
