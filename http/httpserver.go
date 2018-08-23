package main

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	fmt.Fprint(w, GetPlayerScore(player))

}

func GetPlayerScore(player string) string {

	if player == "Pepper" {
		return "20"
	}

	if player == "Floyd" {
		return "10"
	}

	return ""
}
