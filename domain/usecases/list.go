package usecases

import (
	"fmt"

	"github.com/womenwhogocwb/lets-go-wwg-app/domain"
)

func (g Games) ListAll() ([]domain.Game, error) {
	list, err := g.store.ListAll()
	if err != nil {
		return nil, fmt.Errorf("could not list all games: %w\n", err)
	}

	return list, nil
}

func (g Games) ListGameById(id domain.GameId) (domain.Game, error) {
	game, err := g.store.ListOne(id)

	if err != nil {
		return domain.Game{}, fmt.Errorf("Could not list game: %v\n", err)
	}

	return game, nil
}
