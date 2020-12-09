## Hello world
My first program online bank wich 
***
consists of poles of the tasks
***
#####Piece of the code from my list
```markdown
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

```
###### My first md file