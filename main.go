package main

import (
	// "MemberSystem/database"
	// "MemberSystem/routes"

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

	// dbconfig := database.DBConfig{
	// 	User:     "root",
	// 	Password: "",
	// 	Host:     "127.0.0.1",
	// 	Port:     "3306",
	// 	DBName:   "arthur",
	// }

	//database.ConnectDB(dbconfig)
	//defer database.CloseDB()

	//routes.InitRoutes()
}
