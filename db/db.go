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

func open() {
	path := "/tmp/user-db"
	// Initialize the database
	graph.InitQuadStore("bolt", path, nil)

	// Open and use the database
	store, err := cayley.NewGraph("bolt", path, nil)
	if err != nil {
		log.Fatalln(err)
	}

	store.AddQuad(cayley.Quad("person:sophie", "type", "Person", ""))
	store.AddQuad(cayley.Quad("person:sophie", "name", "Sophie Grégoire", ""))
	store.AddQuad(cayley.Quad("person:sophie", "born", "1974", ""))
	store.AddQuad(cayley.Quad("person:sophie", "lives in", "country:canada", ""))

	store.AddQuad(cayley.Quad("person:justin", "type", "Person", ""))
	store.AddQuad(cayley.Quad("person:justin", "name", "Justin Trudeau", ""))
	store.AddQuad(cayley.Quad("person:justin", "born", "1972", ""))
	store.AddQuad(cayley.Quad("person:justin", "lives in", "country:canada", ""))
	store.AddQuad(cayley.Quad("person:justin", "in love with", "person:sophie", ""))
}

func main() {
	open()
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
}

func validEmail(email string) bool {
	fmt.Println("email in func", email)
	return true
}
