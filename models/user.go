package models

import (
	"database/sql"
	"fmt"
)

type Users struct {
	Id int64
	Admin bool
	Name string
	Surname string
	Age  int64
	Gender string
	Login string
	Password string
	Remove bool
}

//func DBuser(db *sql.DB) (err error) {
//	const userDDL = `create table if not exists Users
//(
//    Id integer primary key autoincrement,
//    Admin boolean,
//    Name text not null,
//    Surname text not null,
//    Age integer not null,
//    Gender text not null,
//    Login text not null,
//    Password text not null,
//    Remove boolean not null default false
//)
//`
//	_, err = db.Exec(userDDL)
//	if err != nil {
//		return err
//	}
//	return nil
//}


func AddNewUsers(db *sql.DB, user Users) (err3 error) {

	_, err3 = db.Exec("Insert Into Users(Admin, Name, Surname, Age, Gender, Login, Password, Remove) values (($1), ($2), ($3), ($4), ($5), ($6), ($7), ($8))",user.Admin, user.Name, user.Surname, user.Age, user.Gender, user.Login, user.Password, user.Remove)
	if err3 != nil {
		fmt.Println(err3)
	}
	return err3
}