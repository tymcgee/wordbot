package main

import (
	"fmt"

	"github.com/tymcgee/wordbot/bot"
)

func main() {
	// game.PlayCli()
	numGuessesOnWinningGames := 0
	numGuesses := 0
	wins := 0
	totalGames := 1000
	for range totalGames {
		results := bot.BotGame(bot.FilterGuesses)
		numGuesses += results.Guesses
		if results.Won {
			numGuessesOnWinningGames += results.Guesses
			wins++
		}
	}
	fmt.Printf("Win percentage is ~%.2f%%\n", (float64(wins)/float64(totalGames))*100)
	fmt.Printf("Average number of guesses when winning is %.3f\n", float64(numGuessesOnWinningGames)/float64(totalGames))
	fmt.Printf("Average number of guesses overall is %.3f\n", float64(numGuesses)/float64(totalGames))
}
