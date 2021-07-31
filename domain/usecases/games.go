package usecases

import "github.com/womenwhogocwb/lets-go-wwg-app/store"

type Games struct {
	store store.GameStore
}

func NewGames(store store.GameStore) Games {
	return Games{
		store: store,
	}
}
