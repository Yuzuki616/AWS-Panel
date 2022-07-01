package data

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
)

type AwsSecret struct {
	gorm.Model
	UserId   string
	Name     string
	SecretId string
	Secret   string
}

func AddSecret(username string, name, Id, Secret string) error {
	user := UserData{}
	secret := AwsSecret{}
	Db.Where("Username = ?", username).First(&user)
	if user.ID == 0 {
		return errors.New("用户不存在")
	}
	Db.Where("user_id = ? and Name = ?", user.ID, name).First(&secret)
	if secret.ID != 0 {
		return errors.New("密钥已存在")
	}
	Db.Create(&AwsSecret{
		Name:     name,
		SecretId: Id,
		Secret:   Secret,
		UserId:   strconv.Itoa(int(user.ID)),
	})
	return nil
}

func ListSecret(username string) ([]AwsSecret, error) {
	var user UserData
	var secrets []AwsSecret
	Db.Where("Username = ?", username).First(&user)
	if user.ID == 0 {
		return nil, errors.New("用户不存在")
	}
	Db.Where("user_id = ?", user.ID).Find(&secrets)
	return secrets, nil
}

func DelSecret(username string, name string) error {
	var user UserData
	var secret AwsSecret
	Db.Where("Username = ?", username).First(&user)
	if user.ID == 0 {
		return errors.New("用户不存在")
	}
	Db.Where("user_id = ? and Name = ?", user.ID, name).First(&secret)
	if secret.ID == 0 {
		return errors.New("密钥不存在")
	}
	Db.Delete(&secret)
	return nil
}

func GetSecret(username string, name string) (*AwsSecret, error) {
	var user UserData
	var secret AwsSecret
	Db.Where("Username = ?", username).First(&user)
	if user.ID == 0 {
		return nil, errors.New("用户不存在")
	}
	Db.Where("user_id = ? and Name = ?", user.ID, name).First(&secret)
	if secret.ID == 0 {
		return nil, errors.New("密钥不存在")
	}
	return &secret, nil
}
