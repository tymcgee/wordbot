package main

import "github.com/tymcgee/wordbot/bot"

func main() {
	// game.PlayCli()
	for range 5 {
		bot.BotGame(bot.FilterGuessesNoYellows)
	}
}
