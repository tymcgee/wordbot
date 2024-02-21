package bot

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
