package store

import (
	"testing"

	"github.com/womenwhogocwb/lets-go-wwg-app/domain"
)

func TestStoreGame(t *testing.T) {
	store := NewGameStore()
	t.Run("should store a game successfully", func(t *testing.T) {
		name := "Leslie Knope"
		move := "tesoura"
		game, _ := domain.NewGame(name, move)

		err := store.StoreGame(game)
		if err != nil {
			t.Errorf("expected nil; got '%v'", err)
		}

		if store.gameStore[game.ID].PlayerName != name {
			t.Errorf("expected '%s'; got '%s'", name, store.gameStore[game.ID].PlayerName)
		}
	})
	t.Run("should return error if game ID is empty", func(t *testing.T) {
		game := domain.Game{
			ID:         "",
			PlayerName: "Ann Perkins",
			PlayerMove: "pedra",
		}

		err := store.StoreGame(game)
		if err != ErrEmptyID {
			t.Errorf("expected '%s'; got '%v'", ErrEmptyID, err)
		}
	})
}
