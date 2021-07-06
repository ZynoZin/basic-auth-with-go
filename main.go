package main

import (
	"fmt"
	"net/http"
)

var (
	username = "abc"
	password = "123"
)

func main() {
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	user, pass, ok := r.BasicAuth()
	if !ok {
		fmt.Println("Error parsing basic auth")
		w.WriteHeader(401)
		return
	} else if user != username {
		fmt.Printf("Username provided is incorrect")
		w.WriteHeader(401)
		return
	} else if pass != password {
		fmt.Printf("Password provided is incorrect")
		w.WriteHeader(401)
		return
	}
	fmt.Println("Auth succeeded!\n")
	fmt.Printf("Username: %s\n", user)
	fmt.Printf("Password: %s\n", pass)
	w.WriteHeader(200)
	return
}
