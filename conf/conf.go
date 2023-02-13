package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

var Config *Conf

type Conf struct {
	path             string
	LogLevel         string `json:"LogLevel"`
	Addr             string `json:"Addr"`
	DbPath           string `json:"DbPath"`
	EnableLoadStatic bool   `json:"LoadStatic"`
	StaticPath       string `json:"StaticPath"`
	EnableMailVerify bool   `json:"EnableMailVerify"`
	MailConfig       *Mail  `json:"MailConfig"`
}
type Mail struct {
	Host     string `json:"Host"`
	Port     int    `json:"Port"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func New(path string) *Conf {
	return &Conf{
		path:             path,
		LogLevel:         "error",
		Addr:             ":8011",
		DbPath:           "./data.db",
		EnableLoadStatic: true,
		StaticPath:       "./web",
		EnableMailVerify: false,
		MailConfig: &Mail{
			Host:     "smtp.qq.com",
			Port:     465,
			Email:    "a@qq.com",
			Password: "12321",
		},
	}
}

func Init(path string) error {
	Config = New(path)
	err := Config.LoadConfig()
	if err != nil {
		return err
	}
	return nil
}

func (c *Conf) LoadConfig() error {
	r, err := os.Open(c.path)
	if err != nil {
		return fmt.Errorf("read config file error: %s", err)
	}
	err = json.NewDecoder(r).Decode(c)
	if err != nil {
		return fmt.Errorf("unmarshal config file error: %w", err)
	}
	return nil
}

func (c *Conf) SaveConfig() error {
	f, err := os.OpenFile(c.path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("open file error: %s", err)
	}
	err = json.NewEncoder(f).Encode(c)
	if err != nil {
		return fmt.Errorf("unmarshal config file error: %s", err)
	}
	return nil
}
