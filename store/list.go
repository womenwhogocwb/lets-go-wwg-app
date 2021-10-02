package store

import "github.com/womenwhogocwb/lets-go-wwg-app/domain"

func (g GameStore) ListAll() ([]domain.Game, error) {
	var list []domain.Game
	for _, game := range g.gameStore {
		list = append(list, game)
	}

	return list, nil
}

func (g GameStore) ListOne(id domain.GameId) (domain.Game, error) {
	listAll, _ := g.ListAll()

	var listOne domain.Game

	for _, game := range listAll {
		if id == domain.GameId(game.ID) {
			listOne = game
		}
	}

	if listOne.ID == "" {
		return domain.Game{}, ErrEmptyID
	} else {
		return listOne, nil
	}
}
