package main

import (
	"fmt"
	"github.com/oren/user/db"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/user", newUser)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func newUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, r.Method+" not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := "/tmp/user-db"
	store := db.New(path)
	output, err := store.AddUser(r.Body)

	if err != nil {
		http.Error(w, "error!", http.StatusBadRequest)
		return
	}

	fmt.Println("output", output)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	w.WriteHeader(http.StatusCreated)
}
