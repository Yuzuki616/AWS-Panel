package data

import "gorm.io/gorm"

type Config struct {
	gorm.Model
	EmailVerity int // 0 off,1 on
	MailHost    string
	Port        int
	Email       string
	Password    string
}

func SaveMailConfig(host, email, password string, enable, port int) {
	config := Config{}
	config.EmailVerity = enable
	config.MailHost = host
	config.Port = port
	config.Email = email
	config.Password = password
	Db.Save(&config)
}

func GetMailConfig() *Config {
	config := &Config{}
	Db.First(config)
	return config
}

func IsEmailVerity() bool {
	config := &Config{}
	Db.First(config)
	if config.EmailVerity == 1 {
		return true
	}
	return false
}
