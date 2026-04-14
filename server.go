package poker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()

	router.Handle("GET /league", http.HandlerFunc(p.leagueHandler))
	router.Handle("GET /players/{name}", http.HandlerFunc(p.showScoreHandler))
	router.Handle("POST /players/{name}", http.HandlerFunc(p.processWinHandler))

	p.Handler = router

	return p
}

const JSONContentType = "application/json"

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", JSONContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) showScoreHandler(w http.ResponseWriter, r *http.Request) {
	player := r.PathValue("name")
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWinHandler(w http.ResponseWriter, r *http.Request) {
	player := r.PathValue("name")
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
