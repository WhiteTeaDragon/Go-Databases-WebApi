package main

import (
	"encoding/json"
	"net/http"
)

func usersIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	usersTable, err := ReadAllUsers()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = json.NewEncoder(w).Encode(usersTable)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func showUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	var id User
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	currUser, err := GetUserById(id.Id)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if currUser.Id == -1 {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(400)
		err = json.NewEncoder(w).Encode(nil)
		return
	}
	err = json.NewEncoder(w).Encode(currUser)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err = InsertUser(&newUser)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	var changedUser User
	err := json.NewDecoder(r.Body).Decode(&changedUser)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err = UpdateUser(&changedUser)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func deleteUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	var id User
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	rows, err := DeleteUserById(id.Id)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if rows == 0 {
		http.Error(w, http.StatusText(400), 400)
		return
	}
}
