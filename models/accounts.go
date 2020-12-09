package models

import (
	"database/sql"
	"fmt"
)

type Accounts struct {
	Id int64
	UsersId  int64
	Name string
	Number int64
	Amounte int64
	Currency string
}

//func DBacount(db *sql.DB) (err error)  {
//	const acountDDL = `create table if not exists Accounts (
//    Id integer primary key autoincrement ,
//    UsersId integer not null ,
//    Name text not null ,
//    Number text not null ,
//    Amounte integer ,
//    Currency text
//)
//`
//
//	_, err2 := db.Exec(acountDDL)
//	if err2 != nil {
//		return err2
//	}
//	return nil
//}


func AddNewAcount(db *sql.DB, Accounts Accounts) (err error)  {
	_, err = db.Exec("Insert Into Accounts(UsersId, Number, Name, Amount, Currency) values (($1), ($2), ($3), ($4), ($5))", Accounts.UsersId,Accounts.Name, Accounts.Number, Accounts.Amounte, Accounts.Currency)
	if err != nil {
		fmt.Println(err)
	}
	return err
}