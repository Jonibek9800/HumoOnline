package services

import (
	"HumoLab/OnlineBank26/db"
	"HumoLab/OnlineBank26/models"
	"database/sql"
	"fmt"
	"log"
)

func Translations(Db *sql.DB, user models.Users) {
	fmt.Println("Перевод на карты любых банков...")
	var myAccountNumber, receiverAccountNumber, translationAmount int64
	fmt.Println("Введите номер вашей карты:")
	_, err := fmt.Scan(&myAccountNumber)
	if err != nil {
		log.Print("Неверный ввод", err)
	}
	fmt.Println("Введите номер карты получателя:")
	_, err = fmt.Scan(&receiverAccountNumber)
	if err != nil {
		log.Print("Неверный ввод", err)
	}
	fmt.Println("Сумма:")
	_, err = fmt.Scan(&translationAmount)
	if err != nil {
		log.Print("Неверный ввод", err)
	}
	var myAccount, receiverAccount models.Accounts
	row := Db.QueryRow(db.SelAccountAmount, myAccountNumber)
	err = row.Scan(&myAccount.Amounte)
	if err != nil {
		return
	}
	row = Db.QueryRow(db.SelAccountNumber, myAccountNumber)
	err = row.Scan(&myAccount.Number)
	if err != nil {
		log.Println(err)
	}
	row = Db.QueryRow(db.SelAccountAmount, receiverAccountNumber)
	err = row.Scan(&receiverAccount.Amounte)
	if err != nil {
		log.Println(err)
	}
	row = Db.QueryRow(db.SelAccountNumber, receiverAccountNumber)
	err = row.Scan(&receiverAccount.Number)
	if err != nil {
		log.Println(err)
	}
	if myAccount.Amounte < translationAmount {
		fmt.Println("У Вас недостаточно средств!")
		UserAdminExit(Db, user)
		return
	}
	_, err = Db.Exec(db.UpdAccountAmountOfGiver, translationAmount, myAccountNumber)
	if err != nil {
		log.Println(err)
	}
	_, err = Db.Exec(db.UpdAccountAmountOfGainer, translationAmount, receiverAccountNumber)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Ваша сумма успешно переведено!")
	err = models.AddTransaction(Db, myAccount, receiverAccount, translationAmount)
	if err != nil {
		log.Println(err)
	}
	UserAdminExit(Db, user)
}

//Pokazat address bancomata
func ShowATMsAddresses(Db *sql.DB, user models.Users) {
	rows, err := Db.Query(db.SelATM)
	if err != nil {
		panic(err)
	}
	atms := []models.ATMs{}
	for rows.Next(){
		p := models.ATMs{}
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Status)
		if err != nil{
			fmt.Println(err)
			continue
		}
		atms = append(atms, p)
	}
	for _, p := range atms {
		fmt.Println(p)
	}
	UserAdminExit(Db, user)
}