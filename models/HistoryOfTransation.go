package models

import (
	"HumoLab/OnlineBank26/db"
	"database/sql"
	"fmt"
	"time"
)

type HistoryOfTransaction struct {
	ID int64
	Date string
	OperationAmount int64
	AccountNumber int64
	ReceiverAccountNumber int64
	TransactionLimit int64
}

func AddTransaction(Db *sql.DB, myAccount, receiverAccount  Accounts, operationAmount int64) (err error) {
	var check HistoryOfTransaction
	data :=time.Now()
	check.Date = data.Format("02-Jan-2006")
	check.OperationAmount = operationAmount
	check.AccountNumber = myAccount.Number
	check.TransactionLimit = myAccount.Amounte - operationAmount
	check.ReceiverAccountNumber = receiverAccount.Number
	_, err = Db.Exec(db.AddOfTransaction, check.Date, check.OperationAmount, check.AccountNumber, check.ReceiverAccountNumber, check.TransactionLimit)
	if err != nil {
		fmt.Println(err)
	}
	return
}
