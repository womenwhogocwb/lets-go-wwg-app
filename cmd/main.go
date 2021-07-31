package main

import (
	"log"
	"net/http"

	"github.com/womenwhogocwb/lets-go-wwg-app/domain/usecases"
	api "github.com/womenwhogocwb/lets-go-wwg-app/http"
	"github.com/womenwhogocwb/lets-go-wwg-app/store"
)

const addr = ":3000"

func main() {
	gameStore := store.NewGameStore()
	gamesUsecase := usecases.NewGames(gameStore)
	server := api.NewServer(gamesUsecase)

	log.Printf("starting server on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, server))
}
