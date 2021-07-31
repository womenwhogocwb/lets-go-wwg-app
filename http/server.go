package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/womenwhogocwb/lets-go-wwg-app/domain"
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
	games domain.Usecase
	http.Handler
}

func NewServer(usecase domain.Usecase) Server {
	server := Server{games: usecase}
	router := mux.NewRouter()

	// API Routes
	router.HandleFunc("/games", server.Play).Methods(http.MethodPost)
	router.HandleFunc("/games", server.ListAll).Methods(http.MethodGet)

	server.Handler = router
	return server
}
