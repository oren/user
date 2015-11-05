package db

import (
	"encoding/json"
	"fmt"
	"github.com/Xe/uuid"
	"github.com/google/cayley"
	_ "github.com/google/cayley/graph/bolt"
	"io"
	"log"
)

import "github.com/google/cayley/graph"

func New(path string) Db {
	db := Db{location: path}
	graph.InitQuadStore("bolt", path, nil)
	store, err := cayley.NewGraph("bolt", path, nil)
	db.Store = *store

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Db struct {
	location string
	Store    cayley.Handle
}

func (db Db) AddUser(r io.Reader) (string, error) {
	uuid := uuid.NewUUID()

	if uuid == nil {
		log.Fatal("uuid is nil")
	}

	fmt.Println(uuid)

	decoder := json.NewDecoder(r)

	var u User
	err := decoder.Decode(&u)

	if err != nil {
		return "", fmt.Errorf("invalid json")
	}

	fmt.Println(u)

	ok := validEmail(u.Email)

	if !ok {
		return "", fmt.Errorf("invalid email")
	}

	db.Store.AddQuad(cayley.Quad("person:"+uuid.String(), "type", "Person", ""))
	// store.AddQuad(cayley.Quad("person:UUID", "email", u.email, ""))
	// store.AddQuad(cayley.Quad("person:UUID", "password", u.password, ""))

	return db.location, nil
}

func validEmail(email string) bool {
	return true
}
