package main

import (
	"fmt"
	"net/http"
)

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

func getJson

func main() {

	//Router
	router := http.NewServeMux()

	//Handler
	var wc welcome

	router.Handle("/", wc)

	//Handler Funcs
	router.HandleFunc("/logout", logout)

	router.HandleFunc("/login", login)

	//server
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	//run the server
	server.ListenAndServe()

}
