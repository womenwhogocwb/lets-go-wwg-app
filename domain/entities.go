package domain

import "time"

type Move string

const (
	Pedra   Move = "pedra"
	Papel   Move = "papel"
	Tesoura Move = "tesoura"
)

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
