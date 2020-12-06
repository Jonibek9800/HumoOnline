package services

import (
	"HumoLab/OnlineBank26/db"
	"HumoLab/OnlineBank26/models"
	"database/sql"
	"fmt"
	"log"
	"os"
)

func AdminOperation(Db *sql.DB, user models.Users) {
	fmt.Println("Добро пожаловат",user.Surname, user.Name)
	fmt.Println(db.AdminOperationWindow)
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Println("Неверный ввод, введите число", err)
	}
	switch number {
	case 1:
		ShowAccountAmount(Db, user)
	case 2:
		Translations(Db, user)
	case 3:
		ShowATMsAddresses(Db, user)
	case 4:
		AddAtm(Db)
	case 5:
		ShowHistoryOfTransaction(Db, user)
	case 6:
		return
	default:
		fmt.Println("Неверный ввод")
	}
}



func UserOperation(Db *sql.DB, user models.Users) {
	fmt.Printf("Здравствуйте  %s, выбери команду...\n",user.Surname, user.Name)

	fmt.Println(db.UserOperationWindow)
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Println("Неверный ввод, введите число", err)
	}
	switch number {
	case 1:
		ShowAccountAmount(Db, user)
	case 2:
		Translations(Db, user)
	case 3:
		ShowATMsAddresses(Db, user)
	case 4:
		ShowHistoryOfTransaction(Db, user)
	case 5:
		return
	default:
		fmt.Println("Неверный ввод")
	}
}

func Authorization(Db *sql.DB ) {
	fmt.Println("Главное меню")
	fmt.Println("Welcome to system")
	fmt.Println(db.AuthorizationWindow)
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Println("Неверный ввод, введите число", err)
	}
	switch number {
	case 1:
		myUser := AuthorizationOperation(Db)
		if myUser.Admin == true {
			AdminOperation(Db, myUser)
		} else
		{
			UserOperation(Db, myUser)
		}
	case 2:
		fmt.Println("Good bye")
		os.Exit(0 )

	default:
		log.Println("Выберите 1 или 2")
	}
}
func AuthorizationOperation(Db *sql.DB) (user models.Users) {
	fmt.Println("Введите Ваш логин и пароль...")
	var login, password string
	fmt.Println("Логин:")
	_, err := fmt.Scan(&login)
		fmt.Println(err)
	fmt.Println("Пароль:")
	_, err = fmt.Scan(&password)
	fmt.Println(err)
	row := Db.QueryRow(db.SelUserByLoginPassword, login, password)
	err = row.Scan(
		&user.Id,
		&user.Admin,
		&user.Name,
		&user.Surname,
		&user.Age,
		&user.Gender,
		&user.Login,
		&user.Password,
		&user.Remove)
	if err != nil {
		fmt.Errorf("Vash accaunt ne nayden")
	}
	return
}

