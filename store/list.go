package store

import "github.com/womenwhogocwb/lets-go-wwg-app/domain"

func (g GameStore) ListAll() ([]domain.Game, error) {
	var list []domain.Game
	for _, game := range g.gameStore {
		list = append(list, game)
	}

	return list, nil
}
