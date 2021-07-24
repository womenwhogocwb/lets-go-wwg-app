package domain

import "testing"

func TestNewGame(t *testing.T) {
	t.Run("should create game successfully", func(t *testing.T) {
		name := "Thalía"
		move := "pedra"

		game, err := NewGame(name, move)
		if err != nil {
			t.Errorf("expected nil; got '%s'", err.Error())
		}

		if game.PlayerName != name {
			t.Errorf("expected '%s'; got '%s'", name, game.PlayerName)
		}

		if game.PlayerMove != Move(move) {
			t.Errorf("expected '%s'; got '%s'", move, game.PlayerMove)
		}

		if game.CreatedAt.IsZero() {
			t.Error("expected create at not to be zero")
		}

		if game.ID == "" {
			t.Error("expected ID not to be empty")
		}
	})
	t.Run("should return error when move is invalid", func(t *testing.T) {
		name := "Shakira"
		move := "par"

		_, err := NewGame(name, move)
		if err != ErrInvalidMove {
			t.Errorf("expected '%s'; got '%v'", ErrInvalidMove, err)
		}
	})
	t.Run("should return error when name is invalid", func(t *testing.T) {
		name := ""
		move := "tesoura"

		_, err := NewGame(name, move)
		if err != ErrInvalidName {
			t.Errorf("expected '%s'; got '%v'", ErrInvalidName, err)
		}
	})
}
