package main

import (
	"HumoLab/OnlineBank26/db"
	"HumoLab/OnlineBank26/models"
	"HumoLab/OnlineBank26/pkg/core/services"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)


func main() {
	database, err := sql.Open("sqlite3", "Test")
	db.DbInit(database)
	if err != nil {
		fmt.Println("not found", err)
	}else {
		fmt.Println("CONNECTION TO DB IS SUCCESS")
	}
	models.JsonFunk(database)
	Start(database)
}



func Start(database *sql.DB)  {

	for  {
		services.Authorization(database)

	}
}






