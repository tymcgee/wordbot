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
	for i := range totalGames {
		fmt.Printf("\rPlaying game %d", i+1)
		results := bot.BotGame(answer, bot.FilterGuesses)
		numGuesses += results.Guesses
		if results.Won {
			numGuessesOnWinningGames += results.Guesses
			wins++
		}
	}
	fmt.Printf("\nPlayed %d games\n", totalGames)
	fmt.Printf("Win percentage is ~%.2f%%\n", (float64(wins)/float64(totalGames))*100)
	fmt.Printf("Average number of guesses when winning is %.3f\n", float64(numGuessesOnWinningGames)/float64(totalGames))
	fmt.Printf("Average number of guesses overall is %.3f\n", float64(numGuesses)/float64(totalGames))
}
