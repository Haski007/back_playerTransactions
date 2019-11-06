package main

import (
	"log"
	"net/http"
)

var TokenFromConfig = "testtask"


func main() {
	http.HandleFunc("/user/create", addUser)
	http.HandleFunc("/user/get", getUser)
	http.HandleFunc("/user/deposit", addDeposit)
	http.HandleFunc("/transaction", makeTransaction)

	port := ":8080"
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
