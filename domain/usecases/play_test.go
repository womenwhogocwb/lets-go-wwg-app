package usecases

import (
	"testing"

	"github.com/womenwhogocwb/lets-go-wwg-app/store"
)

func TestGames_Play(t *testing.T) {
	gameStore := store.NewGameStore()
	games := NewGames(gameStore)

	t.Run("should play a game successfully", func(t *testing.T) {
		name := "Donna Meagle"
		move := "tesoura"
		game, err := games.Play(name, move)
		if err != nil {
			t.Errorf("expected nil; got error '%s'", err.Error())
		}

		if game.HouseMove == "" {
			t.Error("expected house move not to be empty")
		}

		if game.Result == "" {
			t.Error("expected result not to be empty")
		}
	})
}
