package data

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/yuzuki999/Aws-Panel/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DbInit(path string) error {
	if utils.IsNotFound(path) {
		defer func() {
			CreateAdminUser()
			log.Info("Done.Default account: admin password: admin123456")
		}()
	}
	db, openErr := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if openErr != nil {
		return fmt.Errorf("open db error: %v", openErr)
	}
	err := db.AutoMigrate(UserData{}, AwsSecret{})
	if err != nil {
		return fmt.Errorf("AutoMigrate error: %v", err)
	}
	Db = db
	return nil
}
