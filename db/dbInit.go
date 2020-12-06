package db

import (
	"database/sql"
	"log"
)

func DbInit(db *sql.DB){
	DDLs := []string{CreateUsersAccount, CreateTransactionTable, CreateAccountTable, CreateATMsTable}
	for _, ddl := range DDLs{
		_, err := db.Exec(ddl)
		if err != nil {
			log.Fatalf("Can't init %s err is %e", ddl, err)
		}
	}
	return
}

//func DBInit(db *sql.DB)  {
//	DDLsk := []string{CreateATMsTable}
//	for _, ddls := range DDLsk{
//		_, err := db.Exec(ddls)
//		if err != nil {
//			log.Fatal(ddls, err)
//		}
//		fmt.Println(ddls)
//	}
//	return
//}
//
