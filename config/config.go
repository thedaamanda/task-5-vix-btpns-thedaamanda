package config

import (
	"project/config/keys"
	"project/database/postgres"
	"strconv"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type (
	config struct {
	}

	Config interface {
		ServiceName() string
		ServiceHost() string
		ServicePort() int
		ServiceEnvironment() string
		Database() *gorm.DB
	}
)

func NewConfig() Config {
	return &config{}
}

func (c *config) Database() *gorm.DB {
	return postgres.InitGorm()
}

func (c *config) ServiceName() string {
	return viper.GetString(keys.AppName)
}

func (c *config) ServiceHost() string {
	return viper.GetString(keys.HostAddress)
}

func (c *config) ServicePort() int {
	v := viper.GetString(keys.HostPort)
	port, _ := strconv.Atoi(v)

	return port
}

func (c *config) ServiceEnvironment() string {
	return viper.GetString(keys.Environment)
}
