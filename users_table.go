package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	Lastname string `json:"lastname,omitempty"`
}

func CreateUsersTable() error {
	createForm, err := database.Prepare("create table " + databaseName + ".Users(" +
											"id int auto_increment primary key, " +
											"name varchar(10000) not null, " +
											"lastname varchar(10000) not null" +
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

func ReadAllUsers() ([]User, error) {
	rows, err := database.Query("select * from " + databaseName + ".users")
	if err != nil {
		return []User{}, err
	}
	defer func() {
		err = rows.Close()
	}()
	var users []User

	for rows.Next() {
		p := User{}
		err := rows.Scan(&p.Id, &p.Name, &p.Lastname)
		if err != nil {
			return users, err
		}
		users = append(users, p)
	}
	return users, err
}

func GetUserById(id int) (User, error) {
	selDB, err := database.Query("SELECT * FROM " + databaseName + ".Users WHERE id=?", id)
	if err != nil {
		return User{}, err
	}
	currUser := User{}
	currUser.Id = -1
	for selDB.Next() {
		err = selDB.Scan(&currUser.Id, &currUser.Name, &currUser.Lastname)
		if err != nil {
			return currUser, err
		}
	}
	return currUser, nil
}

func InsertUser(newUser *User) error {
	insForm, err := database.Prepare("INSERT INTO " + databaseName +".Users(name, lastname) VALUES(?,?)")
	if err != nil {
		return err
	}
	_, err = insForm.Exec(newUser.Name, newUser.Lastname)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(changedUser *User) error {
	_, err := database.Exec("update "+databaseName+".Users set name = ?, lastname = ? where id = ?",
		changedUser.Name, changedUser.Lastname, changedUser.Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserById(id int) (int64, error) {
	res, err := database.Exec("delete from "+databaseName+".Users where id = ?", id)
	if err != nil {
		return 0, err
	}
	rowsAff, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}
	return rowsAff, nil
}

func ClearUsersTable() error {
	_, err := database.Exec("truncate table " + databaseName + ".Users")
	if err != nil {
		return err
	}
	return nil
}
