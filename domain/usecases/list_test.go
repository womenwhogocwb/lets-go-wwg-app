package usecases

import (
	"testing"

	"github.com/womenwhogocwb/lets-go-wwg-app/store"
)

func TestGames_ListAll(t *testing.T) {
	t.Run("should return list of game successfully", func(t *testing.T) {
		gameStore := store.NewGameStore()
		games := NewGames(gameStore)

		_, err := games.Play("April Ludgate", "tesoura")
		if err != nil {
			t.Fatal("game should have been created successfully")
		}
		_, err = games.Play("April Ludgate", "pedra")
		if err != nil {
			t.Fatal("game should have been created successfully")
		}
		_, err = games.Play("Diane Lewis", "papel")
		if err != nil {
			t.Fatal("game should have been created successfully")
		}

		list, err := games.ListAll()
		if err != nil {
			t.Errorf("expected nil; got '%s'", err.Error())
		}

		want := 3
		if len(list) != want {
			t.Errorf("expected %d; got %d", want, len(list))
		}
	})
	t.Run("should return empty slice when there are no games", func(t *testing.T) {
		gameStore := store.NewGameStore()
		games := NewGames(gameStore)

		list, err := games.ListAll()
		if err != nil {
			t.Errorf("expected nil; got '%s'", err.Error())
		}

		want := 0
		if len(list) != want {
			t.Errorf("expected %d; got %d", want, len(list))
		}
	})
}
