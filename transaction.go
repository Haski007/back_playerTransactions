package main

import (
	"errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
**	Structures
*/

type transaction struct {
	TransactionID uint64          `json:"transactionId"`
	UserID        uint64          `json:"userid"`
	Type          TransactionType `json:"type"`
	Amount        float64         `json:"amount"`
	Token         string          `json:"token"`
}

var usersTransactionsMap = map[uint64][]transaction{}

type TransactionType string

const (
	Bet TransactionType = "Bet"
	Win TransactionType = "Win"
)

/*
**	Code
*/


func makeTransaction(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	var data transaction
	var balance float64

	json.Unmarshal(bytes, &data)
	switch data.Type {
	case Bet:
		balance, err = MakeBet(data)
	case Win:
		balance, err = MakeWin(data)
	default:
		err = errors.New("undefined transaction type")
	}
	if err != nil {
		log.Println(err.Error)
		w.Write([]byte(`{"error" :` + err.Error() + `}`))
		return
	}

	w.Write([]byte(fmt.Sprintf("{\"error\" : \"\", \"balance\" :%f}", balance)))
}

func MakeBet(data transaction) (float64, error) {
	if data.Amount <= 0 {
		return 0, errors.New("invalid amount")
	}

	usr, ok := usersMap[data.UserID]
	if !ok {
		return 0, errors.New("there aren't such user")
	}
	if usr.Balance < data.Amount {
		return 0, errors.New("no money")
	}

	usr.Balance -= data.Amount
	usersTransactionsMap[usr.ID] = append(usersTransactionsMap[usr.ID], data)
	usr.BetSum += data.Amount
	usr.BetCount++
	return usr.Balance, nil
}

func MakeWin(data transaction) (float64, error) {
	usr, ok := usersMap[data.UserID]
	if !ok {
		return 0, errors.New("there aren't such user")
	}

	usr.Balance += data.Amount
	usersTransactionsMap[usr.ID] = append(usersTransactionsMap[usr.ID], data)
	usr.WinSum += data.Amount	
	usr.WinCount++
	return usr.Balance, nil
}
