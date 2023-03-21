package main

import (
	"fmt"
	"net/http"
)

type welcome string

func (wc welcome) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to our server!")
}

func main() {
	var wc welcome
	server := http.Server{
		Addr:    ":8080",
		Handler: wc,
	}

	//run the server
	server.ListenAndServe()

}
