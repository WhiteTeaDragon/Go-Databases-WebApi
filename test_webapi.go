package main

import (
	"database/sql"
	"net/http"
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
	if err = db.Ping(); err != nil {
		panic(err.Error())
	}
	return db
}

func init() {
	database = dbConn()
}

func main() {
	http.HandleFunc("/users", usersIndex)
	http.HandleFunc("/calls", callsIndex)
	http.HandleFunc("/users/show", showUserById)
	http.HandleFunc("/calls/show", showCallById)
	http.HandleFunc("/calls/showBySeller", showCallsBySeller)
	http.HandleFunc("/calls/showByCustomer", showCallsByCustomer)
	http.HandleFunc("/calls/showBySelCus", showCallsBySellerCustomer)
	http.HandleFunc("/users/create", createUser)
	http.HandleFunc("/calls/create", addCall)
	http.HandleFunc("/users/update", updateUser)
	http.HandleFunc("/calls/update", updateCall)
	http.HandleFunc("/users/delete", deleteUserById)
	http.HandleFunc("/calls/delete", deleteCallById)
	http.ListenAndServe(":3000", nil)
}
