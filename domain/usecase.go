package domain

type Usecase interface {
	Play(string, string) (Game, error)
	ListAll() ([]Game, error)
	ListGameById(GameId) (Game, error)
}
