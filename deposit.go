package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
**	Structures
*/

type deposit struct {
	DepositID uint64  `json:"depositid"`
	UserID    uint64  `json:"userid"`
	Amount    float64 `json:"amount"`
	Token     string  `json:"token"`
}

var usersDepositMap = map[uint64][]deposit{}


/*
**	Code
*/

func getOldBalance(userId uint64) float64 {
	var oldBalance float64

	deposits, ok := usersDepositMap[userId]
	if !ok {
		return oldBalance
	}

	for _, val := range deposits {
		oldBalance += val.Amount
	}
	return oldBalance
}

func addDeposit(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(`{"error" :` + err.Error() + `}`))
		return
	}
	var data deposit

	json.Unmarshal(bytes, &data)
	usr, ok := usersMap[data.UserID]
	if !ok {
		fmt.Println("there aren't such user")
		w.Write([]byte(`{"error" :` + "there aren't such user" + `}`))
		return
	}
	usr.DepositSum += data.Amount
	usr.DepositCount++
	usr.Balance += getOldBalance(data.UserID) + data.Amount
	deposits := usersDepositMap[usr.ID]
	usersDepositMap[usr.ID] = append(deposits, data)
	fmt.Println(string(bytes))
	w.Write([]byte(fmt.Sprintf("{\"error\" : \"\", \"balance\" : %f}", usr.Balance)))
}