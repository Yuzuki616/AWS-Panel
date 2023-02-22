package data

import (
	"errors"
	"gorm.io/gorm"
)

type AwsSecret struct {
	gorm.Model
	UserId   uint   `gorm:"column:UserId"`
	Name     string `gorm:"column:Name"`
	SecretId string `gorm:"column:SecretId"`
	Secret   string `gorm:"column:Secret"`
}

func AddSecret(email string, name, Id, Secret string) error {
	user := UserData{}
	secret := AwsSecret{}
	client.Where("Email = ?", email).First(&user)
	if user.ID == 0 {
		return errors.New("用户不存在")
	}
	client.Where("UserId = ? and Name = ?", user.ID, name).First(&secret)
	if secret.ID != 0 {
		return errors.New("密钥已存在")
	}
	client.Create(&AwsSecret{
		Name:     name,
		SecretId: Id,
		Secret:   Secret,
		UserId:   user.ID,
	})
	return nil
}

func ListSecret(email string) ([]AwsSecret, error) {
	var user UserData
	var secrets []AwsSecret
	client.Where("Email = ?", email).First(&user)
	if user.ID == 0 {
		return nil, errors.New("用户不存在")
	}
	client.Where("UserId = ?", user.ID).Find(&secrets)
	return secrets, nil
}

func DelSecret(email string, name string) error {
	var user UserData
	var secret AwsSecret
	client.Where("Email = ?", email).First(&user)
	if user.ID == 0 {
		return errors.New("用户不存在")
	}
	client.Where("UserId = ? and Name = ?", user.ID, name).First(&secret)
	if secret.ID == 0 {
		return errors.New("密钥不存在")
	}
	client.Delete(&secret)
	return nil
}

func GetSecret(email string, name string) (*AwsSecret, error) {
	var user UserData
	var secret AwsSecret
	client.Where("Email = ?", email).First(&user)
	if user.ID == 0 {
		return nil, errors.New("用户不存在")
	}
	client.Where("UserId = ? and Name = ?", user.ID, name).First(&secret)
	if secret.ID == 0 {
		return nil, errors.New("密钥不存在")
	}
	return &secret, nil
}
