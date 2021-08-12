package main

import (
	"MemberSystem/database"
	"MemberSystem/routes"

	"github.com/spf13/viper"
)

func init() {

	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {

	dbconfig := database.DBConfig{
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		DBName:   viper.GetString("database.name"),
	}

	database.ConnectDB(dbconfig)
	defer database.CloseDB()

	routes.InitRoutes()
}
