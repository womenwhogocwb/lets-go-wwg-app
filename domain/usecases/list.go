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
