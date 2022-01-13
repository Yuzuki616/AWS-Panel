package data

import (
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

type UserData struct {
	gorm.Model
	//IsAdmin bool `gorm:"IsAdmin"`
	Status   int    `gorm:"Status"` //0 正常,1 封禁
	Username string `gorm:"Username"`
	Password string `gorm:"Password"`
}

func LoginVerify(Username string, PasswordMd5 string) error {
	var user UserData
	Db.Where("Username = ? and password = ?", Username, PasswordMd5).First(&user)
	if user.ID == 0 {
		return errors.New("用户名或密码错误")
	}
	if user.Status == 1 {
		return errors.New("用户已封禁")
	}
	return nil
}

func Register(Username string, PasswordMd5 string) error {
	var user UserData
	Db.Where("Username = ?", Username).First(&user)
	if user.ID != 0 {
		return errors.New("用户已存在")
	}
	user.Username = Username
	user.Password = PasswordMd5
	user.Status = 0
	Db.Create(&user)
	return nil
}

func ChangePassword(username string, OldPasswordMd5, NewPasswordMd5 string) error {
	var user UserData
	Db.Where("Username = ? and password = ?", username, OldPasswordMd5).First(&user)
	if user.ID == 0 {
		return errors.New("用户不存在或密码错误")
	}
	user.Password = NewPasswordMd5
	Db.Save(&user)
	return nil
}
