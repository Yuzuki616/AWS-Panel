package main

import (
	"flag"
	"github.com/Yuzuki616/Aws-Panel/conf"
	"github.com/Yuzuki616/Aws-Panel/data"
	"github.com/Yuzuki616/Aws-Panel/mail"
	"github.com/Yuzuki616/Aws-Panel/router"
	"github.com/Yuzuki616/Aws-Panel/utils"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"time"
)

var config = flag.String("config", ".config.json", "config file path")
var ver = flag.Bool("version", false, "print version message")

var (
	// use ld flags replace
	version   = "master"
	commit    = "none"
	buildDate = "none"
)

func printVersion() {
	log.Info("Aws Panel")
	log.Info("Version: ", version)
	log.Info("Commit: ", commit)
	log.Info("Build Date: ", buildDate)
	log.Info("Github: https://github.com/Yuzuki616/AWS-Panel")
}

func main() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "01-02 15:04:05",
		LogFormat:       "Aws-Panel | %time% | %lvl% >> %msg% \n",
	})
	printVersion()
	if *ver {
		return
	}
	err := conf.Init("./config.json")
	if err != nil {
		log.Error("Init config error: ", err)
	}
	switch conf.Config.LogLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	}
	if utils.RunningByDoubleClick() {
		log.Warning("不建议直接双击运行本程序，这将导致一些非可预料后果，请通过控制台启动本程序")
		log.Warning("将等待10秒后启动")
		time.Sleep(time.Second * 10)
	}
	err = data.Init(conf.Config.DbPath)
	if err != nil {
		log.Error("Database init error: ", err)
	}
	mail.Init()
	router.Init()
	err = router.Start()
	if err != nil {
		log.Error(err)
	}
}
