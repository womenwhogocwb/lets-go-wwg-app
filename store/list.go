package store

import "github.com/womenwhogocwb/lets-go-wwg-app/domain"

func (g GameStore) ListAll() (map[string]domain.Game, error) {
	return g.gameStore, nil
}
