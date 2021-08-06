package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type ListResponse struct {
	ID         string `json:"id"`
	PlayerName string `json:"player_name"`
	PlayerMove string `json:"player_move"`
	HouseMove  string `json:"house_move"`
	Result     string `json:"result"`
	CreatedAt  string `json:"created_at"`
}

func (s Server) ListAll(w http.ResponseWriter, r *http.Request) {
	list, err := s.games.ListAll()
	if err != nil {
		log.Printf("failed to list games: %s\n", err.Error())
		response := Error{Reason: "internal server error"}
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Build response based on the usecase output
	response := make([]ListResponse, len(list))
	for i, game := range list {
		response[i].ID = game.ID
		response[i].PlayerName = game.PlayerName
		response[i].PlayerMove = string(game.PlayerMove)
		response[i].HouseMove = string(game.HouseMove)
		response[i].Result = string(game.Result)
		response[i].CreatedAt = game.CreatedAt.Format(DateLayout)
	}

	w.Header().Set(ContentType, JSONContentType)
	json.NewEncoder(w).Encode(response)
	log.Printf("sent list all games successful response with %d games", len(response))
}
