package store

import "github.com/womenwhogocwb/lets-go-wwg-app/domain"

func (g GameStore) StoreGame(game domain.Game) error {
	if game.ID == "" {
		return ErrEmptyID
	}

	g.gameStore[game.ID] = game
	return nil
}
