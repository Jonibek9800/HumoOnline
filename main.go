package main

import (
	"HumoLab/OnlineBank26/db"
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
	//user := models.User{
	//	Id:       1,
	//	Name:     "Jonibek",
	//	Surname:  "Mahmudov",
	//	Age:      22,
	//	Gender:   "man",
	//	Login:    "124579",
	//	Password: "987654321",
	//	Remove:   true,
	//}
	////struct 1  users
	//open, err3 := sql.Open("sqlite3", "Test")
	//if err3 != nil {
	//	fmt.Println(err3)
	//}
	//err3 = models.DBuser(open)
	//if err3 != nil {
	//	fmt.Println(err3)
	//}
	//err3 = models.AddNewUsers(open, user)
	//if err3 != nil {
	//	fmt.Println(err3)
	//}
	//fmt.Println(user)
	//
	////account sozdani dlya polzovatelya
	//account1 := models.Accounts{
	//	Id:       1,
	//	UsersId:  1,
	//	Number:   2222444456549874,
	//	Amounte:  50,
	//	Currency: "Tjs",
	//}
	//
	////Struct 2 account
	//d, err2 := sql.Open("sqlite3", "Test")
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	//err2 = models.DBacount(d)
	//fmt.Println(err2)
	//err2 = models.AddNewAcount(d, account1)
	//if err2 != nil {
	//	fmt.Println("For give is not added", err2)
	//}
	//fmt.Println(account1)
	Star(database)
}



func Star(database *sql.DB)  {
	for  {
		services.Authorization(database)

	}
}






