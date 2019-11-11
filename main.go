package main

import (
	"log"
	"net/http"

	"github.com/globalsign/mgo"
)

var TokenFromConfig = "testtask"


var usersCollection *mgo.Collection


func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	usersCollection = session.DB("paymets").C("users")
	http.HandleFunc("/user/create", addUser)
	http.HandleFunc("/user/get", getUser)
	http.HandleFunc("/user/deposit", addDeposit)
	http.HandleFunc("/transaction", makeTransaction)

	port := ":8080"
	err = http.ListenAndServe(port, nil)	
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
