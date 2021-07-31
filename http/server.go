package http

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/womenwhogocwb/lets-go-wwg-app/domain/usecases"
)

const (
	ContentType     = "Content-Type"
	JSONContentType = "application/json"
	DateLayout      = "2006-01-02T15:04:05Z"
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
	router.HandleFunc("/games", server.ListAll).Methods(http.MethodGet)

	server.Handler = router
	return server
}
