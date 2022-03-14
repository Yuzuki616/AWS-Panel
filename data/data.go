package data

import (
	log "github.com/sirupsen/logrus"
	"github.com/yuzuki999/Aws-Panel/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DbInit(path string) error {
	db, openErr := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if openErr != nil {
		return openErr
	}
	Db = db
	if utils.IsNotFound(path) {
		log.Info("Init database")
		err := db.AutoMigrate(UserData{}, AwsSecret{})
		if err != nil {
			return err
		}
		CreateAdminUser()
		log.Info("Default account: admin password: admin123456")
	}
	return nil
}
