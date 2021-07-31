package http

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/womenwhogocwb/lets-go-wwg-app/domain/usecases"
)

type Error struct {
	Reason string `json:"reason"`
}

type Server struct {
	games usecases.Games
	http.Handler
}

func NewServer(games usecases.Games) Server {
	server := Server{games: games}
	router := mux.NewRouter()

	// API Routes
	router.HandleFunc("/games", server.Play).Methods(http.MethodPost)

	server.Handler = router
	return server
}
