package store

import (
	"testing"

	"github.com/womenwhogocwb/lets-go-wwg-app/domain"
)

func TestGameStore_ListAll(t *testing.T) {
	store := NewGameStore()

	t.Run("should return all games successfully", func(t *testing.T) {
		game1, err := domain.NewGame("Donna Meagle", "pedra")
		if err != nil {
			t.Fatal("game should have been created successfully")
		}
		err = store.StoreGame(game1)
		if err != nil {
			t.Fatal("game should have been stored successfully")
		}

		game2, err := domain.NewGame("Ann Perkins", "tesoura")
		if err != nil {
			t.Fatal("game should have been created successfully")
		}
		err = store.StoreGame(game2)
		if err != nil {
			t.Fatal("game should have been stored successfully")
		}

		games, err := store.ListAll()
		if err != nil {
			t.Errorf("expected nil; got '%s'", err.Error())
		}

		want := 2
		if len(games) != want {
			t.Errorf("expected %d; got %d", want, len(games))
		}

		if games[game1.ID] != game1 {
			t.Errorf("expected %v; got %v", game1, games[game1.ID])
		}

		if games[game2.ID] != game2 {
			t.Errorf("expected %v; got %v", game2, games[game2.ID])
		}
	})
}
