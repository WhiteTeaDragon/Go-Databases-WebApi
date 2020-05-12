package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type Call struct {
	Id            int    `json:"id"`
	Seller        int    `json:"seller,omitempty"`
	Customer      int    `json:"customer,omitempty"`
	CallTimestamp string `json:"callTimestamp,omitempty"`
	Duration      string `json:"duration,omitempty"`
	Video         string `json:"video,omitempty"`
	Transcription string `json:"transcription,omitempty"`
}

func CreateCallsTable() error {
	createForm, err := database.Prepare("create table " + databaseName + ".Calls(" +
											"id int auto_increment primary key, " +
											"seller int not null, " +
											"customer int not null, " +
											"callTimestamp datetime not null, " +
											"duration time not null, " +
											"video varchar(10000) not null, " +
											"transcription varchar(10000) not null" +
											") charset=utf8;")
	if err != nil {
		return err
	}
	_, err = createForm.Exec()
	if err != nil {
		return err
	}
	return nil
}

func ReadAllCalls() ([]Call, error) {
	rows, err := database.Query("select * from " + databaseName + ".calls")
	if err != nil {
		return []Call{}, err
	}
	defer func() {
		err = rows.Close()
	}()
	var calls []Call

	for rows.Next() {
		p := Call{}
		err := rows.Scan(&p.Id, &p.Seller, &p.Customer, &p.CallTimestamp, &p.Duration, &p.Video, &p.Transcription)
		if err != nil {
			return calls, err
		}
		calls = append(calls, p)
	}
	return calls, err
}

func GetCallById(id int) (Call, error) {
	selDB, err := database.Query("SELECT * FROM " + databaseName + ".Calls WHERE id=?", id)
	if err != nil {
		return Call{}, err
	}
	defer func() {
		err = selDB.Close()
	}()

	currCall := Call{}
	currCall.Id = -1
	for selDB.Next() {
		err = selDB.Scan(&currCall.Id, &currCall.Seller, &currCall.Customer, &currCall.CallTimestamp,
			&currCall.Duration, &currCall.Video, &currCall.Transcription)
		if err != nil {
			return currCall, err
		}
	}
	return currCall, err
}

func GetCallsBySeller(seller int) ([]Call, error) {
	selDB, err := database.Query("SELECT * FROM " + databaseName + ".Calls WHERE seller=?", seller)
	if err != nil {
		return []Call{}, err
	}
	defer func() {
		err = selDB.Close()
	}()

	var calls []Call
	for selDB.Next() {
		p := Call{}
		err := selDB.Scan(&p.Id, &p.Seller, &p.Customer, &p.CallTimestamp, &p.Duration, &p.Video, &p.Transcription)
		if err != nil {
			return calls, err
		}
		calls = append(calls, p)
	}
	return calls, err
}

func GetCallsByCustomer(customer int) ([]Call, error) {
	selDB, err := database.Query("SELECT * FROM " + databaseName + ".Calls WHERE customer=?", customer)
	if err != nil {
		return []Call{}, err
	}
	defer func() {
		err = selDB.Close()
	}()

	var calls []Call
	for selDB.Next() {
		p := Call{}
		err := selDB.Scan(&p.Id, &p.Seller, &p.Customer, &p.CallTimestamp, &p.Duration, &p.Video, &p.Transcription)
		if err != nil {
			return calls, err
		}
		calls = append(calls, p)
	}
	return calls, err
}

func GetCallsBySellerCustomer(seller, customer int) ([]Call, error) {
	selDB, err := database.Query("SELECT * FROM " + databaseName + ".Calls WHERE seller=? and customer=?", seller,
		customer)
	if err != nil {
		return []Call{}, err
	}
	defer func() {
		err = selDB.Close()
	}()

	var calls []Call
	for selDB.Next() {
		p := Call{}
		err := selDB.Scan(&p.Id, &p.Seller, &p.Customer, &p.CallTimestamp, &p.Duration, &p.Video, &p.Transcription)
		if err != nil {
			return calls, err
		}
		calls = append(calls, p)
	}
	return calls, err
}

func InsertCall(newCall *Call) error {
	insForm, err := database.Prepare("INSERT INTO " + databaseName + ".Calls(seller, customer, callTimestamp, " +
		"duration, video, transcription) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = insForm.Exec(newCall.Seller, newCall.Customer, newCall.CallTimestamp, newCall.Duration, newCall.Video,
		newCall.Transcription)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCall(changedCall *Call) error {
	_, err := database.Exec("update "+databaseName+".Calls set seller = ?, customer = ?, callTimestamp = ?, " +
		"duration = ?, video = ?, transcription = ? where id = ?", changedCall.Seller, changedCall.Customer,
		changedCall.CallTimestamp, changedCall.Duration, changedCall.Video, changedCall.Transcription, changedCall.Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCallById(id int) (int64, error) {
	res, err := database.Exec("delete from "+databaseName+".Calls where id = ?", id)
	if err != nil {
		return 0, err
	}
	rowsAff, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}
	return rowsAff, nil
}

func ClearCallsTable() error {
	_, err := database.Exec("truncate table " + databaseName + ".Calls")
	if err != nil {
		return err
	}
	return nil
}
