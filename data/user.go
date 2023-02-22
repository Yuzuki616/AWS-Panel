package data

import (
	"errors"
	"github.com/Yuzuki616/Aws-Panel/utils"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

type UserData struct {
	gorm.Model
	Status   int    `gorm:"column:Status"` //0 正常,1 封禁
	Email    string `gorm:"column:Email"`
	Password string `gorm:"column:Password"`
	IsAdmin  int    `gorm:"column:IsAdmin"` //0 否,1 是
}

func VerifyUser(email string, Password string) error {
	var user UserData
	client.Where("Email = ?", email).First(&user)
	if user.ID == 0 || !utils.VerifyPasswordHash(Password, user.Password) {
		return errors.New("用户名或密码错误")
	}
	if user.Status == 1 {
		return errors.New("用户已封禁")
	}
	return nil
}

func CreateUser(email, Email, Password string, IsAdmin int) error {
	var user UserData
	client.Where("Email = ?", email).First(&user)
	if user.ID != 0 {
		return errors.New("用户已存在")
	}
	user.Email = email
	user.Email = Email
	Password, err := utils.GenPasswordHash(Password)
	if err != nil {
		return err
	}
	user.Password = Password
	user.Status = 0
	user.IsAdmin = IsAdmin
	client.Create(&user)
	return nil
}

func ChangeEmail(oldEmail, newEmail, Password string) error {
	var user UserData
	client.Where("Email = ?", oldEmail).First(&user)
	if user.ID == 0 || !utils.VerifyPasswordHash(Password, user.Password) {
		return errors.New("用户名或密码错误")
	}
	user.Email = newEmail
	client.Save(&user)
	return nil
}

func ChangeUserPassword(email string, oldPassword, newPassword string) error {
	var user UserData
	client.Where("Email = ?", email).First(&user)
	if user.ID == 0 {
		return errors.New("用户不存在或密码错误")
	}
	if !utils.VerifyPasswordHash(oldPassword, user.Password) {
		return errors.New("用户不存在或密码错误")
	}
	hash, _ := utils.GenPasswordHash(newPassword)
	user.Password = hash
	client.Save(&user)
	return nil
}

func IsAdmin(email string) bool {
	var user UserData
	client.Where("Email = ?", email).First(&user)
	if user.ID == 0 {
		return false
	}
	if user.IsAdmin == 1 {
		return true
	}
	return false
}

func DeleteUser(email string) error {
	var user UserData
	client.Where("Email = ?", email).First(&user)
	if user.ID == 0 {
		return errors.New("用户不存在")
	}
	if user.IsAdmin == 1 {
		return errors.New("不能删除管理员账户")
	}
	client.Delete(&user)
	return nil
}

func BanUser(email string) error {
	var user UserData
	client.Where("Email = ?", email).First(&user)
	if user.ID == 0 {
		return errors.New("用户不存在")
	}
	if user.IsAdmin == 1 {
		return errors.New("不能封禁管理员账户")
	}
	if user.Status == 0 {
		user.Status = 1
		client.Save(&user)
	} else {
		return errors.New("该用户已被封禁")
	}
	return nil
}

func UnBanUser(email string) error {
	var user UserData
	client.Where("Email = ?", email).First(&user)
	if user.ID == 0 {
		return errors.New("用户不存在")
	}
	if user.IsAdmin == 1 {
		return errors.New("不能解封管理员账户")
	}
	if user.Status == 1 {
		user.Status = 0
		client.Save(&user)
	} else {
		return errors.New("该用户未被封禁")
	}
	return nil
}

type UserInfo struct {
	Email   string
	Status  int
	IsAdmin int
}

func GetUserList() ([]UserInfo, error) {
	var u []UserData
	client.Find(&u)
	if len(u) == 0 {
		return nil, errors.New("未找到任何用户")
	}
	var users []UserInfo
	for i := range u {
		users = append(users, UserInfo{
			Email:   u[i].Email,
			Status:  u[i].Status,
			IsAdmin: u[i].IsAdmin,
		})
	}
	return users, nil
}
