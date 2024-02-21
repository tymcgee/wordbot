package bot

import (
	"math/rand"

	"github.com/tymcgee/wordbot/game"
)

func RandomGame() {
	g := game.Game{
		ShowOngoingStats: true,
		ShowStats:        false,
		ShowGuesses:      true,
	}
	g.PlayGame(func(lastGuess string, lastGuessStats string) string {
		idx := rand.Intn(len(game.VALID_GUESSES))
		return game.VALID_GUESSES[idx]
	})
}

func BotGame(filterMethod func(validGuesses []string, gray []GameInformation, yellow []GameInformation, green []GameInformation) []string) {
	g := game.Game{
		ShowOngoingStats: true,
		ShowStats:        false,
		ShowGuesses:      true,
	}
	gray := make([]GameInformation, 0)
	yellow := make([]GameInformation, 0)
	green := make([]GameInformation, 0)
	validGuesses := game.VALID_GUESSES[:]

	g.PlayGame(func(lastGuess string, lastGuessStats string) string {
		if lastGuessStats == "" {
			// this is the first round, we don't have stats yet
			idx := rand.Intn(len(game.VALID_GUESSES))
			return game.VALID_GUESSES[idx]
		}

		getGameInfo(gray, yellow, green, lastGuess, lastGuessStats)

		validGuesses = filterMethod(validGuesses, gray, yellow, green)

		idx := rand.Intn(len(validGuesses))
		return validGuesses[idx]
	})
}

func getGameInfo(gray []GameInformation, yellow []GameInformation, green []GameInformation, lastGuess string, lastGuessStats string) {
	for i, letter := range []rune(lastGuessStats) {
		switch string(letter) {
		// TODO: try to stop filling in duplicates?
		case game.GREEN:
			green = append(green, GameInformation{
				Letter: []rune(lastGuess)[i],
				Index:  i,
			})
		case game.YELLOW:
			yellow = append(yellow, GameInformation{
				Letter: []rune(lastGuess)[i],
				Index:  i,
			})
		case game.GRAY:
			gray = append(gray, GameInformation{
				Letter: []rune(lastGuess)[i],
				Index:  i,
			})
		}
	}
}
