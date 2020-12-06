package db

const  CreateUsersAccount = `create table if not exists Users
(
    Id integer primary key autoincrement,
    Admin boolean,
    Name text not null,
    Surname text not null,
    Age integer not null,
    Gender text not null,
    Login text not null,
    Password text not null,
    Remove boolean not null default false
);
`

const CreateATMsTable = `create table if not exists atms
(
    Id integer primary key autoincrement,
    name text not null,
	status boolean
);
`
const CreateTransactionTable = `create table if not exists HistoryOfTransaction (
Id integer primary key autoincrement,
Date text not null,
OperationAmount integer not null,
AccountNumber integer not null,
ReceiverAccountNumber integer not null,
TransactionLimit integer not null
)`

const CreateAccountTable = `create table if not exists Accounts (
Id integer primary key autoincrement unique,
UsersId integer references Users(id),
Name text not null,
Number integer not null unique,
Amount integer not null,
Currency text not null
)`
