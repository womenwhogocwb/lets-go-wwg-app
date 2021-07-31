package usecases

import (
	"fmt"

	"github.com/womenwhogocwb/lets-go-wwg-app/domain"
)

func (g Games) ListAll() ([]domain.Game, error) {
	storedGames, err := g.store.ListAll()
	if err != nil {
		return nil, fmt.Errorf("could not list all games: %w\n", err)
	}

	var list []domain.Game
	for _, game := range storedGames {
		list = append(list, game)
	}

	return list, nil
}
