package main

import (
	"log"
	"net/http"

	"github.com/AndreReyesG/poker"
)

func main() {
	handler := http.HandlerFunc(poker.PlayerServer)
	log.Print("starting server on :5000")
	log.Fatal(http.ListenAndServe(":5000", handler))
}
