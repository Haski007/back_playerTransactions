package main

import "gopkg.in/macaron.v1"
import (
	"net/http"
// 	"log"
	"fmt"
)

/*
**	Structures
*/

type user struct {
	id			int
	balance		float64
	token		string
}

func (u user) createUser(id int, balance float64, token string) {
	u.id = id
	u.balance = balance
	u.token = token
}

const GusersApiResp = `
<html>
<head>
	<meta charset="UTF-8">
	<meta name="mainPage">
	<title>mainPage</title>
</head>
<body>
	<form action="main.go" method="POST">	
		<p><input type="text" name="fname">Name</p>
		<p><input type="text" name="lname">Last name</p>
		<input type="submit" value="send">
	</form>
</body>
</html>
`
var GusersCounter int

var GreportCounter int

func main() {
	m := macaron.Classic()
	// m.Get("/", func() string {
	// 	return "Hello world!"
	// })
	m.Get("/", mainPage)
	// var u user
	// u.createUser(1, 7.8, "DaAaaaAaaa")
	// fmt.Println(u)
	m.Get("/", mainPage)
	m.Run()
}


func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("We got a request on /users")
	GreportCounter++
	str := fmt.Sprintf("/reports API call count: %v", GreportCounter)
	fmt.Fprint(w, str)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	GusersCounter++
	str := fmt.Sprintf(GusersApiResp, r.Method, GusersCounter)
	fmt.Fprint(w, str)
}

// func usersHandleFunc(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("We got a request on /users")
// 	GusersCounter++
// 	str := fmt.Sprintf(GusersApiResp, r.Method, GusersCounter)
// 	fmt.Fprint(w, str)
// }