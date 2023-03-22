package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// User Struct
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Handler interface
type welcome string

func (wc welcome) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to our server!")
}

// Handler function
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Page!")
}

// Handler function
func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout Page!")
}

func getJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	switch r.Method {
	case "GET":
		w.Write([]byte(`"message":"GET is called!"`))
	case "POST":
		w.Write([]byte(`"message":"POST is called!"`))
	}
}

func checkUser(w http.ResponseWriter, r *http.Request) {
	var user User
	dbPassword := "neaj101"
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatal("error decoding into struct")
	}

	if user.Password == dbPassword {
		fmt.Println("Success Logged In")
	}

	fmt.Fprintf(w, "Response: %v", user)
}

func main() {

	//Router
	router := http.NewServeMux()

	//Handler
	var wc welcome

	router.Handle("/", wc)

	//Handler Funcs
	router.HandleFunc("/logout", logout)

	router.HandleFunc("/login", login)

	router.HandleFunc("/json", getJson)

	router.HandleFunc("/user", checkUser)

	//server
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	//run the server
	server.ListenAndServe()

}
