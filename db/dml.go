package db

const AddNewATM = 	`Insert Into atms(name) values (($1))`

const AddOfTransaction = `insert into HistoryOfTransaction(date, operationAmount, accountNumber, receiverAccountNumber, TransactionLimit)
values (($1), ($2), ($3), ($4), ($5))`

const SelUserByLoginPassword = `select *from Users where login = ($1) and password = ($2)`

const SelAccountByUserId = `select * from Accounts where Accounts.UsersId = ($1)`

const SelAccountAmount = `select Amount from Accounts where number = ($1)`

const SelAccountNumber = `select Number from Accounts where Number = ($1)`

const UpdAccountAmountOfGiver = `update Accounts set Amount = Amount - ($1) where Number = ($2)`

const UpdAccountAmountOfGainer = `update Accounts set Amount = Amount + ($1) where Number = ($2)`

const SelATM = `select * from atms`

const SelHistoryOfTransaction = `select * from HistoryOfTransaction where AccountNumber = ($1)`


const AuthorizationWindow = `1.Авторизация
2.Выход`

const AdminOperationWindow = `1.Показать баланс
2.Переводы
3.Адреса банкоматов
4.Добавить банкомат
5.История транзакций
6.Выйти`

const UserOperationWindow = `1.Показать баланс
2.Переводы
3.Адреса банкоматов
4.История транзакций
5.Выйти`