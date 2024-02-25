package bot

import (
	"math/rand"

	"github.com/tymcgee/wordbot/game"
)

type FilterMethod = func(validGuesses []string, gray []GameInformation, yellow []GameInformation, green []GameInformation) []string

// answer may be empty, in which case a real answer will be generated
// firstGuess may be empty, in which case the first guess will be random
func BotGame(answer string, firstGuess string, filter FilterMethod) game.Results {
	g := game.Game{
		ShowOngoingStats: false,
		ShowStats:        false,
		ShowGuesses:      false,
		ShowIntro:        false,
		ShowWinText:      false,
		ShowAnswerOnLoss: false,
	}
	gray := make([]GameInformation, 0)
	green := make([]GameInformation, 0)
	validGuesses := game.VALID_GUESSES[:]

	return g.PlayGameWithAnswer(answer, func(lastGuess string, lastGuessStats string) string {
		// remaking the yellow list on every loop will be nicer.
		// current logic forces the guess to contain all yellows from
		// previous answer, so either they're all still yellow
		// or some of them turned green.
		yellow := make([]GameInformation, 0)
		if lastGuessStats == "" {
			// this is the first round, we don't have stats yet
			if firstGuess == "" {
				idx := rand.Intn(len(game.VALID_GUESSES))
				return game.VALID_GUESSES[idx]
			}
			return firstGuess
		}

		gray, yellow, green = getGameInfo(gray, yellow, green, lastGuess, lastGuessStats)

		validGuesses = filter(validGuesses, gray, yellow, green)

		idx := rand.Intn(len(validGuesses))
		return validGuesses[idx]
	})
}

// append creates a new copy in the function which is why we have to return them back out
func getGameInfo(gray []GameInformation, yellow []GameInformation, green []GameInformation, lastGuess string, lastGuessStats string) ([]GameInformation, []GameInformation, []GameInformation) {
	for i, letter := range []rune(lastGuessStats) {
		switch string(letter) {
		// note: this will make duplicates, but it shouldn't matter in the filter function
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
	return gray, yellow, green
}
