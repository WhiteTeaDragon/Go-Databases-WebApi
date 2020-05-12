package main

import (
	"encoding/json"
	"net/http"
)

func callsIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	callsTable, err := ReadAllCalls()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = json.NewEncoder(w).Encode(callsTable)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func showCallById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	var id Call
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	currCall, err := GetCallById(id.Id)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if currCall.Id == -1 {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(400)
		err = json.NewEncoder(w).Encode(nil)
		return
	} else {
		err = json.NewEncoder(w).Encode(currCall)
	}
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func showCallsBySeller(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	var seller Call
	err := json.NewDecoder(r.Body).Decode(&seller)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	calls, err := GetCallsBySeller(seller.Seller)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = json.NewEncoder(w).Encode(calls)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func showCallsByCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	var customer Call
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	calls, err := GetCallsByCustomer(customer.Customer)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = json.NewEncoder(w).Encode(calls)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func showCallsBySellerCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	var info Call
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	calls, err := GetCallsBySellerCustomer(info.Seller, info.Customer)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = json.NewEncoder(w).Encode(calls)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func addCall(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	var newCall Call
	err := json.NewDecoder(r.Body).Decode(&newCall)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err = InsertCall(&newCall)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func updateCall(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	var changedCall Call
	err := json.NewDecoder(r.Body).Decode(&changedCall)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err = UpdateCall(&changedCall)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func deleteCallById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	var id Call
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	rows, err := DeleteCallById(id.Id)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if rows == 0 {
		http.Error(w, http.StatusText(400), 400)
		return
	}
}
