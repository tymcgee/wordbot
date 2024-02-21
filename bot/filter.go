package bot

import (
	"slices"
)

type GameInformation struct {
	Letter rune
	Index  int
}

// first try, don't bother with yellows
func FilterGuessesNoYellows(validGuesses []string, gray []GameInformation, yellow []GameInformation, green []GameInformation) []string {
	newValidGuesses := make([]string, 0)

	for _, guess := range validGuesses {
		isValid := true
		guessRunes := []rune(guess)
		for _, info := range gray {
			if guessRunes[info.Index] == info.Letter {
				isValid = false
				break
			}
		}

		for _, info := range green {
			if guessRunes[info.Index] != info.Letter {
				isValid = false
				break
			}
		}
		if isValid {
			newValidGuesses = append(newValidGuesses, guess)
		}
	}

	return newValidGuesses
}

func FilterGuesses(validGuesses []string, gray []GameInformation, yellow []GameInformation, green []GameInformation) []string {
	newValidGuesses := make([]string, 0)

	for _, guess := range validGuesses {
		isValid := true
		guessRunes := []rune(guess)
		for _, grayInfo := range gray {
			if guessRunes[grayInfo.Index] == grayInfo.Letter {
				isValid = false
				break
			}
		}

		for _, yellowInfo := range yellow {
			if !slices.Contains(guessRunes, yellowInfo.Letter) {
				isValid = false
			}
			// don't the same letter in the same spot twice
			if guessRunes[yellowInfo.Index] == yellowInfo.Letter {
				isValid = false
			}
		}

		for _, greenInfo := range green {
			if guessRunes[greenInfo.Index] != greenInfo.Letter {
				isValid = false
				break
			}
		}
		if isValid {
			newValidGuesses = append(newValidGuesses, guess)
		}
	}

	return newValidGuesses
}
