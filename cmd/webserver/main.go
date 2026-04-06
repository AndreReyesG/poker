package main

import (
	"log"
	"net/http"

	"github.com/AndreReyesG/poker"
)

func main() {
	server := poker.NewPlayerServer(poker.NewInMemoryPlayerStore())
	log.Print("starting server on :5000")
	log.Fatal(http.ListenAndServe(":5000", server))
}
