type user struct {
	ID           uint64  `json:"id"`
	Balance      float64 `json:"balance"`
	DepositCount int
	Token        string `json:"token"`
}

/*
**	Code
*/

var usersMap = map[uint64]*user{}

func getUser(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	var dat user

	json.Unmarshal(bytes, &dat)

	res, err := json.Marshal(usersMap[dat.ID])
	if err != nil {
		w.Write([]byte(`{"error" :` + err.Error() + `}`))
		return
	}
	w.Write([]byte(res))
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var u user

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = json.Unmarshal(bytes, &u)
	if err != nil {
		w.Write([]byte(`{"error" :` + err.Error() + `}`))
		return
	}

	if u.Token != TokenFromConfig {
		w.Write([]byte(`{"error" : "invalid token"}`))
		return
	}

	if _, ok := usersMap[u.ID]; ok {
		w.Write([]byte(`{"error": "user already exist"}`))
		return
	} else {
		usersMap[u.ID] = &u
		w.Write([]byte(`{"error" : ""}`))

	}
}