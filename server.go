package main

import (
	"fmt"
	"github.com/oren/user/db"
	"log"
	"net/http"
)

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

	db.AddUser(w, r.Body)
	w.WriteHeader(http.StatusCreated)
	return
}
