package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidMove = errors.New("invalid move")
	ErrInvalidName = errors.New("invalid name")
)

type Move string

const (
	Pedra   Move = "pedra"
	Papel   Move = "papel"
	Tesoura Move = "tesoura"
)

var ValidMoves = [3]Move{Pedra, Papel, Tesoura}

type Result string

const (
	Victory Result = "victory"
	Defeat  Result = "defeat"
	Draw    Result = "draw"
)

type Game struct {
	ID         string
	PlayerName string
	PlayerMove Move
	HouseMove  Move
	Result     Result
	CreatedAt  time.Time
}

func NewGame(name, move string) (Game, error) {
	if Move(move) != Pedra && Move(move) != Papel && Move(move) != Tesoura {
		return Game{}, ErrInvalidMove
	}

	if name == "" {
		return Game{}, ErrInvalidName
	}

	return Game{
		ID:         uuid.New().String(),
		PlayerName: name,
		PlayerMove: Move(move),
		CreatedAt:  time.Now(),
	}, nil
}
