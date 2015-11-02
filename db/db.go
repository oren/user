package db

import (
	"fmt"
	"github.com/google/cayley"
	_ "github.com/google/cayley/graph/bolt"
	"log"
	"net/http"
)

import "github.com/google/cayley/graph"

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

func Register(w http.ResponseWriter) {
	fmt.Fprintf(w, "register already (:")
}
