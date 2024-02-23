package bot

import (
	"slices"
)

type GameInformation struct {
	Letter rune
	Index  int
}

func FilterNone(validGuesses []string, gray []GameInformation, yellow []GameInformation, green []GameInformation) []string {
	return validGuesses
}

func FilterGuesses(validGuesses []string, gray []GameInformation, yellow []GameInformation, green []GameInformation) []string {
	newValidGuesses := make([]string, 0)

	// fmt.Printf("validGuesses: %v\n", validGuesses)

	for _, guess := range validGuesses {
		isValid := true
		guessRunes := []rune(guess)
		for _, greenInfo := range green {
			// essentially hardmode -- if a letter was green,
			// it must be green again in the next guess
			if guessRunes[greenInfo.Index] != greenInfo.Letter {
				// fmt.Printf("invalid guess [%v] because of green [%v]\n", guess, string(greenInfo.Letter))
				isValid = false
				break
			}
		}

		// double letters can have a gray but still have that letter in the word..
		// for example, "speed" with the first e green and second e gray
		for _, grayInfo := range gray {
			if !slices.Contains(guessRunes, grayInfo.Letter) {
				continue
			}

			// now the only case when it IS valid is if it maps to a green or yellow
			// that we already know about
			if len(green) == 0 && len(yellow) == 0 {
				// fmt.Printf("invalid guess [%v] because gray + no green [%v]\n", guess, string(grayInfo.Letter))
				isValid = false
				break
			}

			// check yellow
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
			indices := getIndices(guessRunes, grayInfo.Letter)
			isInGreen := false
			for _, i := range indices {
				for _, greenInfo := range green {
					if greenInfo.Index == i {
						isInGreen = true
					}
				}
			}

			if !isInGreen {
				// fmt.Printf("invalid guess [%v] because gray [%v]\n", guess, string(grayInfo.Letter))
				isValid = false
				break
			}
		}

		for _, yellowInfo := range yellow {
			// also hardmode -- must contain a letter that was previously yellow
			if !slices.Contains(guessRunes, yellowInfo.Letter) {
				// fmt.Printf("invalid guess [%v] because doesn't contain yellow letter [%v]\n", guess, string(yellowInfo.Letter))
				isValid = false
				break
			}
			// don't put the same letter in the same spot twice
			if guessRunes[yellowInfo.Index] == yellowInfo.Letter {
				// fmt.Printf("invalid guess [%v] because yellow letter in same spot twice [%v]\n", guess, string(yellowInfo.Letter))
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

func getIndices(letters []rune, key rune) []int {
	indices := make([]int, 0)
	for i, letter := range letters {
		if letter == key {
			indices = append(indices, i)
		}
	}
	return indices
}
