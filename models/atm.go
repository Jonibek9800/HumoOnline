package models

import (
	"HumoLab/OnlineBank26/db"
	"database/sql"
	"fmt"
	"log"
)

type ATMs struct {
	ID int64
	Name string
	Status bool
}

func AddATM(database *sql.DB, address string) (ok bool, err error) {
	_, err = database.Exec(db.AddNewATM, address)
	if err != nil {
		log.Println("Can't insert to ATM is not fount", err)
		return false, err
	}
	fmt.Println(err)
	return true, nil
}
