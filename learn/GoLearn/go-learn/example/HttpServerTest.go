package example

import (
	"fmt"
	"log"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}

func PlayerServer2(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	fmt.Fprint(w, GetPlayerScore(player))
}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}


func Start() {
	handler := http.HandlerFunc(PlayerServer2)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
