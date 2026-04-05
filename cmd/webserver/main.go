package main

import (
	"log"
	"net/http"

	"github.com/AndreReyesG/poker"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := poker.NewPlayerServer(&InMemoryPlayerStore{})
	log.Print("starting server on :5000")
	log.Fatal(http.ListenAndServe(":5000", server))
}
