package main

import (
	"encoding/json"
	"fmt"
	"github.com/oren/user/db"
	"log"
	"net/http"
)

type User struct {
	email    string
	password string
}

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/user", register)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, r.Method+" not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)

	var u User
	err := decoder.Decode(&u)

	if err != nil {
		panic(err)
	}

	fmt.Println(u)

	ok := validEmail(u.Email)

	if !ok {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	db.Register(u)
	w.WriteHeader(http.StatusCreated)
	return
}

func validEmail(email string) bool {
	fmt.Println("email in func", email)
	return true
}
