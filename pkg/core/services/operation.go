package services

import (
	"HumoLab/OnlineBank26/db"
	"HumoLab/OnlineBank26/models"
	"database/sql"
	"fmt"
	"log"
)

func ShowAccountAmount(Db *sql.DB, account models.Users) {
	rows, err := Db.Query(db.SelAccountByUserId, account.Id)
	if err != nil {
		panic(err)
	}
	accounts := []models.Accounts{}
	for rows.Next(){
		p := models.Accounts{}
		err := rows.Scan(
			&p.Id,
			&p.UsersId,
			&p.Name,
			&p.Number,
			&p.Amounte,
			&p.Currency)
		if err != nil{
			fmt.Println(err)
			continue
		}
		accounts = append(accounts, p)
	}
	if len(accounts) > 1 {
		fmt.Println("Ваш баланс на карте...")
		var total int64
		for _, p := range accounts{
			fmt.Printf("%s составляет %d %s \n", p.Name, p.Amounte, p.Currency)
			total += p.Amounte
		}
		fmt.Println("В сумме у Вас %d %s", total, accounts[0].Currency)
	} else {
		fmt.Println("Ваш баланс на карте  составляет  \n", accounts[0].Name, accounts[0].Amounte, accounts[0].Currency)
	}
	//if account.Admin == true {
	//	AdminOperation(Db, account)
	//}else {
	//	UserOperation(Db, account)
	//}
	UserAdminExit(Db, account)
}



func UserAdminExit(db *sql.DB, user models.Users) {
	fmt.Println("\n Нажмите 1 чтобы выйти")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Println("Неверный ввод", err)
	}
	switch input {
	case 1:
		if user.Admin == true {
			AdminOperation(db, user)
		} else {
			UserOperation(db, user)
		}
	default:
		log.Println("Неверный ввод", err)
	}
}


func ShowHistoryOfTransaction(Db *sql.DB, user models.Users) {
	rows, err := Db.Query(db.SelAccountByUserId, user.Id)
	if err != nil {
		panic(err)
	}
	acc := []models.Accounts{}
	for rows.Next(){
		p := models.Accounts{}
		err := rows.Scan(
			&p.Id,
			&p.UsersId,
			&p.Name,
			&p.Number,
			&p.Amounte,
			&p.Currency)
		if err != nil{
			fmt.Println(err)
			continue
		}
		acc = append(acc, p)
	}
	var t int
	for i := 0; i < len(acc); i++ {
		t = i
	}
	if t < 1 {
		rows, err := Db.Query(db.SelHistoryOfTransaction, acc[0].Number)
		if err != nil {
			panic(err)
		}
		arr := []models.HistoryOfTransaction{}
		for rows.Next(){
			p := models.HistoryOfTransaction{}
			err := rows.Scan(
				&p.ID,
				&p.Date,
				&p.OperationAmount,
				&p.AccountNumber,
				&p.ReceiverAccountNumber,
				&p.TransactionLimit)
			if err != nil{
				fmt.Println(err)
				continue
			}
			arr = append(arr, p)
		}
		for i := 0; i < len(arr); i++ {
			fmt.Println("    ОАО ХУМО БАНК   ")
			fmt.Println("Номер банкомата:",arr[i].ID)
			fmt.Println("Дата:",arr[i].Date)
			fmt.Println("Сумма операции:",arr[i].OperationAmount)
			fmt.Println("Номер вашей карты:",arr[i].AccountNumber)
			fmt.Println("Номер карты получателя:",arr[i].ReceiverAccountNumber)
			fmt.Println("Доступный лимит:",arr[i].TransactionLimit, "\n")
		}
	} else {
		fmt.Println("Введите номер вашей карты:")
		var accountNumber int64
		_, err := fmt.Scan(&accountNumber)
		if err != nil {
			panic(err)
		}
		rows, err := Db.Query(db.SelHistoryOfTransaction, accountNumber)
		if err != nil {
			panic(err)
		}
		arr := []models.HistoryOfTransaction{}
		for rows.Next(){
			p := models.HistoryOfTransaction{}
			err := rows.Scan(
				&p.ID,
				&p.Date,
				&p.OperationAmount,
				&p.AccountNumber,
				&p.ReceiverAccountNumber,
				&p.TransactionLimit)
			if err != nil{
				fmt.Println(err)
				continue
			}
			arr = append(arr, p)
		}
		for i := 0; i < len(arr); i++ {
			fmt.Println("    ОАО ХУМО БАНК   ")
			fmt.Println("Номер банкомата:",arr[i].ID)
			fmt.Println("Дата:",arr[i].Date)
			fmt.Println("Сумма операции:",arr[i].OperationAmount)
			fmt.Println("Номер вашей карты:",arr[i].AccountNumber)
			fmt.Println("Номер карты получателя:",arr[i].ReceiverAccountNumber)
			fmt.Println("Доступный лимит:",arr[i].TransactionLimit, "\n")
		}
	}
	UserAdminExit(Db, user)
}
