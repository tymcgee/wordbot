package bot

import (
	"slices"
)

type GameInformation struct {
	Letter rune
	Index  int
}

// Strategy where you don't filter any, so every guess becomes random
func FilterNone(validGuesses []string, gray []GameInformation, yellow []GameInformation, green []GameInformation) []string {
	return validGuesses
}

// If we know where a green is, all valid guesses must have that letter in that spot.
// Grays are a bit trickier because of double letters.
// If we know where a yellow is, valid guesses must have that letter in the guess, but it
// can't be in the same place as last time.
func FilterGuesses(validGuesses []string, gray []GameInformation, yellow []GameInformation, green []GameInformation) []string {
	newValidGuesses := make([]string, 0)

	for _, guess := range validGuesses {
		guessRunes := []rune(guess)

		// GREEN
		if !areGreensValid(guessRunes, gray, yellow, green) {
			continue
		}

		// GRAY
		if !areGraysValid(guessRunes, gray, yellow, green) {
			continue
		}

		// YELLOW
		if !areYellowsValid(guessRunes, gray, yellow, green) {
			continue
		}

		newValidGuesses = append(newValidGuesses, guess)
	}

	return newValidGuesses
}

func areGreensValid(guess []rune, gray []GameInformation, yellow []GameInformation, green []GameInformation) bool {
	for _, greenInfo := range green {
		// essentially hardmode -- if a letter was green,
		// it must be green again in the next guess
		if guess[greenInfo.Index] != greenInfo.Letter {
			return false
		}
	}
	return true
}

func areGraysValid(guess []rune, gray []GameInformation, yellow []GameInformation, green []GameInformation) bool {
	// double letters can have a gray but still have that letter in the word..
	// for example, "speed" with the first e green or yellow and second e gray
	for _, grayInfo := range gray {
		if !slices.Contains(guess, grayInfo.Letter) {
			continue
		}

		// now the only case when it IS valid is if it maps to a green or yellow
		// that we already know about
		if len(green) == 0 && len(yellow) == 0 {
			return false
		}

		// check yellow
		// (the case where it's a double letter and we already have a yellow clue for that letter)
		isYellow := false
		for _, yellowInfo := range yellow {
			if yellowInfo.Letter == grayInfo.Letter {
				isYellow = true
			}
		}
		if isYellow {
			// we are good, this is an okay case
			continue
		}

		// check green
		indices := getIndices(guess, grayInfo.Letter)
		isInGreen := false
		for _, i := range indices {
			for _, greenInfo := range green {
				if greenInfo.Index == i {
					isInGreen = true
				}
			}
		}

		if !isInGreen {
			return false
		}
	}
	return true
}

func areYellowsValid(guess []rune, gray []GameInformation, yellow []GameInformation, green []GameInformation) bool {
	for _, yellowInfo := range yellow {
		// also hardmode -- must contain a letter that was previously yellow
		if !slices.Contains(guess, yellowInfo.Letter) {
			return false
		}
		// don't put the same letter in the same spot twice
		if guess[yellowInfo.Index] == yellowInfo.Letter {
			return false
		}
	}
	return true
}

func getIndices(letters []rune, key rune) []int {
	indices := make([]int, 0)
	for i, letter := range letters {
		if letter == key {
			indices = append(indices, i)
		}
	}
	return indices
}
