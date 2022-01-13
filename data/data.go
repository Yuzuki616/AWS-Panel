package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DbInit(path string) error {
	db, openErr := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if openErr != nil {
		return openErr
	}
	err := db.AutoMigrate(UserData{}, AwsSecret{})
	if err != nil {
		return err
	}
	Db = db
	return nil
}
