package main

import (
	"MemberSystem/database"
	"MemberSystem/routes"
)

func main() {
	dbconfig := database.DBConfig{
		User:     "root",
		Password: "",
		Host:     "127.0.0.1",
		Port:     "3306",
		DBName:   "arthur",
	}

	database.ConnectDB(dbconfig)
	defer database.CloseDB(database.DB)

	routes.InitRoutes()
}
