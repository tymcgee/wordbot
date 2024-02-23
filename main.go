package main

import (
	"fmt"

	"github.com/tymcgee/wordbot/bot"
)

func main() {
	// game.PlayCli()
	wins := 0
	totalGames := 1000
	for range totalGames {
		if bot.BotGame(bot.FilterGuesses) {
			wins++
		}
	}
	fmt.Printf("Won [%v] games out of [%v]", wins, totalGames)
}
