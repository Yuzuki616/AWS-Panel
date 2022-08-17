package main

import (
	"github.com/Yuzuki616/Aws-Panel/data"
	"github.com/Yuzuki616/Aws-Panel/router"
	"github.com/Yuzuki616/Aws-Panel/utils"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"time"
)

const version = "0.3.6"

func printVersion() {
	log.Info("Aws Panel")
	log.Info("Version: ", version)
	log.Info("Github: https://github.com/Yuzuki616/AWS-Panel")
}

func main() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "01-02 15:04:05",
		LogFormat:       "Aws-Panel | %time% | %lvl% >> %msg% \n",
	})
	printVersion()
	if utils.RunningByDoubleClick() {
		log.Warning("不建议直接双击运行本程序，这将导致一些非可预料后果，请通过控制台启动本程序")
		log.Warning("将等待10秒后启动")
		time.Sleep(time.Second * 10)
	}
	dbErr := data.DbInit("./data.db")
	if dbErr != nil {
		log.Error("Database init error: ", dbErr)
	}
	route := router.New()
	route.LoadRoute()
	startErr := route.Start()
	if startErr != nil {
		log.Error(startErr)
	}
}
