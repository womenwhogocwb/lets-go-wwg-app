package store

import (
	"errors"

	"github.com/womenwhogocwb/lets-go-wwg-app/domain"
)

var (
	ErrEmptyID = errors.New("ID cannot be empty")
)

type GameStore struct {
	gameStore map[string]domain.Game
}

func NewGameStore() GameStore {
	ns := make(map[string]domain.Game)
	return GameStore{
		gameStore: ns,
	}
}
