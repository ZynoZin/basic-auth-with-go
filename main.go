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
	if !ok || user != username || pass != password{
		fmt.Println("Authentication failed!")
		w.WriteHeader(401)
		return
	fmt.Println("Authentication succeeded!\n")
	fmt.Printf("Username: %s\n", user)
	fmt.Printf("Password: %s\n", pass)
	w.WriteHeader(200)
	return
}
