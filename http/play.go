package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/womenwhogocwb/lets-go-wwg-app/domain"
)

type PlayRequest struct {
	Name string `json:"name"`
	Move string `json:"move"`
}

type PlayResponse struct {
	ID         string `json:"id"`
	PlayerName string `json:"player_name"`
	PlayerMove string `json:"player_move"`
	HouseMove  string `json:"house_move"`
	Result     string `json:"result"`
	CreatedAt  string `json:"created_at"`
}

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
		log.Printf("failed to play game: %s\n", err.Error())
		switch {
		case errors.Is(err, domain.ErrInvalidMove):
			response := Error{Reason: "invalid move"}
			w.Header().Set(ContentType, JSONContentType)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
		case errors.Is(err, domain.ErrInvalidName):
			response := Error{Reason: "invalid name"}
			w.Header().Set(ContentType, JSONContentType)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
		default:
			response := Error{Reason: "internal server error"}
			w.Header().Set(ContentType, JSONContentType)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
		}
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
