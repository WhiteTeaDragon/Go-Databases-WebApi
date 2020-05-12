package main

import (
	"database/sql"
	"fmt"
)

var databaseName = "information_flows"
var database *sql.DB

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "1234567890"
	dbName := databaseName
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	database = dbConn()
	CreateCallsTable()
	CreateUsersTable()
	GerWay := User{name: "Джерард", lastname: "Вей"}
	ReyTor := User{name: "Ray", lastname: "Toro"}
	FraIer := User{name: "Frank", lastname: "Iero"}
	MikWay := User{name: "Mikey", lastname: "Way"}
	MeMyself := User{name: "Sasha", lastname: "Sasha"}
	InsertUser(&GerWay)
	InsertUser(&ReyTor)
	InsertUser(&FraIer)
	InsertUser(&MikWay)
	InsertUser(&MeMyself)
	call1 := Call{seller: 15, customer: 57, callTimestamp: "1000-01-01 00:00:00", duration: "18.00", video: "C:\\",
		transcription: "C:\\"}
	call2 := Call{seller: 18, customer: 179, callTimestamp: "1099-03-23 15:23:23", duration: "00.01", video: "C:\\",
		transcription: "C:\\"}
	call3 := Call{seller: 1543, customer: 57, callTimestamp: "2018-02-01 12:00:01", duration: "12.23", video: "C:\\",
		transcription: "C:\\"}
	InsertCall(&call1)
	InsertCall(&call2)
	InsertCall(&call3)
	usersTable, err := ReadAllUsers()
	if err != nil {
		panic(err.Error())
	}
	id := -1
	for i := 0; i < len(usersTable); i++ {
		fmt.Println(usersTable[i].id, usersTable[i].name, usersTable[i].lastname)
		if usersTable[i].name == "Mikey" {
			id = usersTable[i].id
		}
	}

	MikeyUser, err := GetUserById(id)
	fmt.Println("Our favourite ", MikeyUser.id, MikeyUser.name, MikeyUser.lastname)
	MikeyUser.name = "Milky"
	UpdateUser(&MikeyUser)

	fmt.Println("After update:")
	usersTable, err = ReadAllUsers()
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < len(usersTable); i++ {
		fmt.Println(usersTable[i].id, usersTable[i].name, usersTable[i].lastname)
	}

	MikeyUser, err = GetUserById(id)
	fmt.Println("Our second favourite ", MikeyUser.id, MikeyUser.name, MikeyUser.lastname)
	DeleteUserById(id)

	callsTable, err := ReadAllCalls()
	if err != nil {
		panic(err.Error())
	}
	id = -1
	fmt.Println("Reading Calls Table")
	fmt.Println(len(callsTable))
	for i := 0; i < len(callsTable); i++ {
		fmt.Println(callsTable[i].id, callsTable[i].seller, callsTable[i].customer, callsTable[i].callTimestamp,
			callsTable[i].duration, callsTable[i].video, callsTable[i].transcription)
		if callsTable[i].seller == 1543 {
			id = callsTable[i].id
		}
	}
	BestCall, err := GetCallById(id)
	fmt.Println("Our favourite ", BestCall.id, BestCall.customer, BestCall.seller)
	BestCall.customer = 666
	UpdateCall(&BestCall)

	callsTable, err = ReadAllCalls()
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < len(callsTable); i++ {
		fmt.Println(callsTable[i].id, callsTable[i].customer, callsTable[i].seller)
	}

	BestCall, err = GetCallById(id)
	fmt.Println("Our second favourite ", BestCall.id, BestCall.customer, BestCall.seller)
	DeleteCallById(id)

	usersTable, err = ReadAllUsers()
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < len(usersTable); i++ {
		fmt.Println(usersTable[i].id, usersTable[i].name, usersTable[i].lastname)
	}

	callsTable, err = ReadAllCalls()
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < len(callsTable); i++ {
		fmt.Println(callsTable[i].id, callsTable[i].customer, callsTable[i].seller)
	}

	ClearUsersTable()
	ClearCallsTable()
	err = database.Close()
	if err != nil {
		panic(err.Error())
	}
}
