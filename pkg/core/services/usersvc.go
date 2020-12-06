package services

import (
	"HumoLab/OnlineBank26/models"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
)

//const LoginOperation = `Введите логин и пароль:`
//const AuthorizedOperation = `1. Показать Баланс
//2.Перевод Денег
//3.Оплата Услуг
//4.История транзакций
//5.Добавить Адресс банкомата
//6.Показат доступные банкоматы
//7.Выход`
//
//func Authorized(database *sql.DB, Id int64){
//	for {
//		fmt.Println(AuthorizedOperation)
//		fmt.Println(`Выбери команду:`)
//		var number int64
//		fmt.Scan(&number)
//		switch number {
//		case 1:
//			fmt.Println(`Показываю Баланс`)
//			CheckBalance(database, Id)
//		case 2:
//			//TODO: Сделать функцию для перевода денег
//			fmt.Println(`Перевод денег`)
//			Tranclation(database,)
//		case 3:
//			//TODO: СДелать функцию для Оплаты услуги
//			fmt.Println(`Оплачиваю услугу`)
//		case 4:
//			//TODO:Сделать функцию для истории транзакций
//			fmt.Println(`Показываю историю транзакций`)
//		case 5:
//			AddAtm(database)
//		case 6:
//			ShowATMsAddresses(database)
//		case 7:
//			return
//		default:
//			fmt.Println("Некорректный ввод попробуйте еще раз")
//		}
//	}
//	return
//}

////Summa Balansa
//func CheckBalance(database *sql.DB, Id int64){
//	var Amounte int
//	err := database.QueryRow(`select Amounte from Accounts
//where UsersId = ($1)`, Id).Scan(
//		&Amounte,
//		)
//	fmt.Println(err)
//	fmt.Println(Amounte)
//}
//

//Dobavlenie addres bancomata
func AddAtm(db *sql.DB) (ok bool) {
	fmt.Println("Enter ATMs address")
	var s string
	fmt.Scan(&s)
	reader := bufio.NewReader(os.Stdin)
	Address, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Can't read command: %v", err)
		return false
	}
	fmt.Println(s)
	sprintf := fmt.Sprintf("%s %s", s, Address)
	fmt.Println(sprintf)
	_, err = models.AddATM(db, sprintf)
	if err != nil {
		fmt.Println("vse ploxo")
		return false
	}
	fmt.Printf("Был добавлен АТМ по адрессу: %s %s\n",s, Address)
	//_, err := models.AddATM(database, text.Text())
	return true
}

//
//func Login(database *sql.DB) (ok bool, id int64){
//	fmt.Println(LoginOperation)
//	var login, password string
//	fmt.Println("login:")
//	fmt.Scan(&login)
//	fmt.Println("password:")
//	fmt.Scan(&password)
//
//	var User models.User
//	_ = database.QueryRow( `Select *from users where login = ($1)
//and password = ($2)`, login, password).Scan(
//		&User.Id,
//		&User.Name,
//		&User.Surname,
//		&User.Age,
//		&User.Gender,
//		&User.Login,
//		&User.Password,
//		&User.Remove)
//	if User.Id > 0{
//		return false, User.Id
//	}
//	return true, User.Id
//}
//
//const transfer = "Введите номер карты на которую переводите"
//
//func Tranclation(database *sql.DB) (ok bool, amount string)  {
//		var account models.Accounts
//		fmt.Println(transfer)
//		var card string
//		fmt.Scan(&card)
//		fmt.Println("ведите сумму ")
//		fmt.Scan(&amount)
//		fmt.Println("Введите номер вашей карты")
//		var namberCard =  account.Number
//		fmt.Scan(&namberCard)
//		fmt.Println("Успешно переведено")
//
// database.QueryRow(`select *from Accounts where Amounte = ($1)`, card, amount ).Scan(
// 	&account.Id,
// 	&account.UsersId,
// 	&account.Number,
// 	&account.Amounte,
// 	&account.Currency)
//	if account.Amounte > 0 {
//		return false, string(account.Amounte)
//	}
//	return true, string(account.Amounte)
//}




//
//func Authorization(database *sql.DB) (Login, password string) {
//	fmt.Println(AuthorizationOperation)
//		var number int64
//	fmt.Scan(&number)
//		switch number {
//	case 1:
//		fmt.Println(LoginOperation)
//			fmt.Println("Login:")
//		fmt.Scan(&Login)
//			fmt.Println("password:")
//		fmt.Scan(&password)
//			fmt.Println("Добро пожаловать")
//		Facilitiec(database)
//
//	case 2:
//		fmt.Println("Enter bancomat Address")
//			var address string
//		fmt.Scan(&address)
//		reader := bufio.NewReader(os.Stdin)
//		readString, err := reader.ReadString('\n')
//		if err != nil {
//			log.Fatal("owibka")
//			return
//		}
//		fmt.Scan(readString)
//		_, err1 := models.AddATM(database, readString)
//		if err1 != nil {
//			fmt.Println(err1)
//		}
//		fmt.Println(address)
//		//fmt.Scan(&address)
//		//_, err := models.AddATM(database, address)
//		//if err != nil {
//		//	fmt.Println("Sorry acsec to not", err)
//		//}
//	case 3:
//
//	case 4:
//		fmt.Println("Good Bye")
//		log.Fatal("Exit")
//
//		default:
//		fmt.Println("Некорректный ввод попробуйте ещё раз")
//	}
//	return
//}
//
//func UserLogin(database *sql.DB, login, password string)  {
//	var User models.User
//	 _ = database.QueryRow(`select *from  users where login = ($1)
//	and  password = ($2)`, login, password).Scan(
//		&User.Id,
//		&User.Name,
//		&User.Surname,
//		&User.Age,
//		&User.Gender,
//		&User.Login,
//		&User.Password,
//		&User.Remove,
//		)
//	fmt.Println(User)
//}

