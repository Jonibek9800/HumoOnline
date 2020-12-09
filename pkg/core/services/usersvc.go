package services

import (
	"HumoLab/OnlineBank26/models"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
)


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
	_, err = models.AddAdressATM(db, sprintf)
	if err != nil {
		fmt.Println("vse ploxo")
		return false
	}
	fmt.Printf("Был добавлен АТМ по адрессу: %s %s\n",s, Address)
	//_, err := models.AddATM(database, text.Text())
	return true
}
