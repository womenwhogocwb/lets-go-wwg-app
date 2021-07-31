package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type PlayRequest struct {
	Name string
	Move string
}

type PlayResponse struct {
	ID         string
	PlayerName string
	PlayerMove string
	HouseMove  string
	Result     string
	CreatedAt  string
}

const (
	ContentType     = "Content-Type"
	JSONContentType = "application/json"
	DateLayout      = "2006-01-02T15:04:05Z"
)

func (s Server) Play(w http.ResponseWriter, r *http.Request) {
	var body PlayRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		response := Error{Reason: "invalid request body"}
		log.Printf("error decoding body: %s\n", err.Error())
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	game, err := s.games.Play(body.Name, body.Move)
	if err != nil {
		// TODO handle errors
		log.Printf("failed to play game: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := PlayResponse{
		ID:         game.ID,
		PlayerName: game.PlayerName,
		PlayerMove: string(game.PlayerMove),
		HouseMove:  string(game.HouseMove),
		Result:     string(game.Result),
		CreatedAt:  game.CreatedAt.Format(DateLayout),
	}

	w.Header().Set(ContentType, JSONContentType)
	json.NewEncoder(w).Encode(response)
	log.Printf("sent successful response for game %s\n", game.ID)
}
