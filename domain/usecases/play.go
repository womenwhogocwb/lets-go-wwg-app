package usecases

import (
	"math/rand"
	"time"

	"github.com/womenwhogocwb/lets-go-wwg-app/domain"
)

func (g Games) Play(name, move string) (domain.Game, error) {
	game, err := domain.NewGame(name, move)
	if err != nil {
		return domain.Game{}, err
	}

	playedGame := play(game)
	err = g.store.StoreGame(playedGame)
	if err != nil {
		return domain.Game{}, err
	}

	return playedGame, nil
}

func play(game domain.Game) domain.Game {
	houseMove := playAsHouse()
	game.HouseMove = houseMove

	if game.PlayerMove == houseMove {
		game.Result = domain.Draw
		return game
	}

	switch game.PlayerMove {
	case domain.Pedra:
		if houseMove == domain.Tesoura {
			game.Result = domain.Victory
		} else if houseMove == domain.Papel {
			game.Result = domain.Defeat
		}
	case domain.Papel:
		if houseMove == domain.Pedra {
			game.Result = domain.Victory
		} else if houseMove == domain.Tesoura {
			game.Result = domain.Defeat
		}
	case domain.Tesoura:
		if houseMove == domain.Papel {
			game.Result = domain.Victory
		} else if houseMove == domain.Pedra {
			game.Result = domain.Defeat
		}
	}

	return game
}

func playAsHouse() domain.Move {
	rand.Seed(time.Now().UnixNano())
	houseMove := domain.ValidMoves[rand.Intn(3)]
	return houseMove
}
