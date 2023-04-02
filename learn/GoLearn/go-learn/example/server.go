package example

import (
	"fmt"
	"log"
	"net/http"
)

type PlayerServer3 struct {
	store PlayerStore
}

type PlayerStore interface {
	GetPlayerScore(name string) int
}

func (p *PlayerServer3) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	player := r.URL.Path[len("/players/"):]

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func StartServer() {

	server := &PlayerServer3{}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

func (p *PlayerServer3) ServeHTTP2(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.processWin(w)
	case http.MethodGet:
		p.showScore(w, r)
	}
}

func (p *PlayerServer3) showScore(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players"):]

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer3) processWin(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}

/////////////////////////////////////////////////////////////
type PlayerServer4 struct {
	store StubPlayerStore4
}

type StubPlayerStore4 struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore4) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore4) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (p *PlayerServer4) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	player := r.URL.Path[len("/players/"):]

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}