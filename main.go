package main

import (
	"HumoLab/OnlineBank26/db"
	"HumoLab/OnlineBank26/models"
	"HumoLab/OnlineBank26/pkg/core/services"
	"HumoLab/OnlineBank26/pkg/core/zip"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
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
	files := []string{"main.go", "settings.json"}
	output := "done.zip"

	err := zip.ZipFiles(output, files)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Zipped File:", output)
	for  {
		services.Authorization(database)

	}
}






