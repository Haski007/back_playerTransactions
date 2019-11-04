package main

// import "gopkg.in/macaron.v1"
import (
	"html/template"
	"net/http"
	"log"
	// "fmt"
)

/*
**	Structures
*/
type User struct {
	Id			int
	Balance		float64
	Token		string

}

// func (u user) createUser(id int, balance float64, token string) {
// 	u.id = id
// 	u.balance = balance
// 	u.token = token
// }
var GusersCounter int
var GreportCounter int

func	main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/users", userPage)

	port := ":8080"
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func 	userPage(w http.ResponseWriter, r *http.Request) {
	// u := []user{
	// 	id: 1,
	// 	balance: 10.2,
	// 	token: "Sobaka",
	// }
	users := []User{User{1, 10.2, "Hren"}, User{2, 22.3, "Ssanina"}}
	GusersCounter++
	tmpl, err := template.ParseFiles("static/users.html")
	if (err != nil) {
		http.Error(w, err.Error(), 400)
		return
	}
	// fmt.Println(users)
	if err := tmpl.Execute(w, users); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func	mainPage(w http.ResponseWriter, r *http.Request) {
	GusersCounter++
	tmpl, err := template.ParseFiles("static/index.html")
	if (err != nil) {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}


// func ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("We got a request on /users")
// 	GreportCounter++
// 	str := fmt.Sprintf("/reports API call count: %v", GreportCounter)
// 	fmt.Fprint(w, str)
// }

// func usersHandleFunc(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("We got a request on /users")
	// 	GusersCounter++
	// 	str := fmt.Sprintf(GusersApiResp, r.Method, GusersCounter)
	// 	fmt.Fprint(w, str)
	// }