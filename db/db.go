package db

import (
	"encoding/json"
	"fmt"
	"github.com/google/cayley"
	_ "github.com/google/cayley/graph/bolt"
	"io"
	"log"
	"net/http"
)

import "github.com/google/cayley/graph"

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AddUser(w http.ResponseWriter, r io.Reader) {
	fmt.Println("in db func")
	decoder := json.NewDecoder(r)

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

	insert(u)
}

func validEmail(email string) bool {
	fmt.Println("validEmail", email)
	return true
}

func insert(u User) error {
	fmt.Println("insert", u)
	path := "/tmp/user-db"
	graph.InitQuadStore("bolt", path, nil)

	store, err := cayley.NewGraph("bolt", path, nil)

	if err != nil {
		log.Fatalln(err)
	}

	store.AddQuad(cayley.Quad("person:sophie", "type", "Person", ""))
	store.AddQuad(cayley.Quad("person:sophie", "email", "3432432e", ""))
	return nil
}
