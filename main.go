package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/tymcgee/wordbot/bot"
	"github.com/tymcgee/wordbot/game"
)

var IsAlpha = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func main() {
	var isBot bool
	var answer string
	var numOfGames int
	var firstGuess string
	flag.BoolVar(&isBot, "bot", false, "whether to use the bot")
	flag.StringVar(&answer, "a", "", "predefined answer")
	flag.IntVar(&numOfGames, "n", 1, "number of games to play (only works in conjunction with --bot)")
	flag.StringVar(&firstGuess, "first", "", "first guess to use (only works in conjunction with --bot) (default first guess is random)")
	flag.Parse()

	validateWord(answer)

	if !isBot {
		game.PlayCli(answer)
		return
	}

	validateWord(firstGuess)

	numGuessesOnWinningGames := 0
	numGuesses := 0
	wins := 0
	for i := range numOfGames {
		fmt.Printf("\rPlaying game %d", i+1)
		results := bot.BotGame(answer, firstGuess, bot.FilterGuesses)
		numGuesses += results.Guesses
		if results.Won {
			numGuessesOnWinningGames += results.Guesses
			wins++
		}
	}
	fmt.Printf("\nPlayed %d games\n", numOfGames)
	fmt.Printf("Win percentage is ~%.2f%%\n", (float64(wins)/float64(numOfGames))*100)
	fmt.Printf("Average number of guesses when winning is %.3f\n", float64(numGuessesOnWinningGames)/float64(numOfGames))
	fmt.Printf("Average number of guesses overall is %.3f\n", float64(numGuesses)/float64(numOfGames))
}

// Exits if the word is not five alpha characters or empty.
func validateWord(word string) {
	if (len(word) != 0 && len(word) != 5) || !IsAlpha(word) {
		fmt.Printf("\"%v\" is an invalid word\n", word)
		os.Exit(1)
	}
}
