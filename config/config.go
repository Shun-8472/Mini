package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

const (
	configName = "config"
	fileType   = "yaml"
	configPath = "./config"
)

var C = new(Configurations)

type Configurations struct {
	Server Server
	Mysql  Mysql
}

type Server struct {
	Host string
	Port int
}

type Mysql struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     int
}

func InitConfigs() {
	initViper()
	fromConfig()
}

func initViper() {
	viper.SetConfigName(configName)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(configPath)
}

func fromConfig() {
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic("Read configs error，reason：" + err.Error())
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		log.Panic("Unmarshal error: " + err.Error())
	}
}

func GetServerAddress() string {
	return fmt.Sprintf("%s:%d", C.Server.Host, C.Server.Port)
}

func GetMySqlAddress() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", C.Mysql.DBUser, C.Mysql.DBPassword, C.Mysql.DBHost, C.Mysql.DBPort, C.Mysql.DBName)
}
