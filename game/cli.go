package game

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func PlayCli() {
	g := Game{
		ShowStats:        true,
		ShowOngoingStats: true,
		ShowGuesses:      false,
		ShowIntro:        true,
		ShowWinText:      true,
		ShowAnswerOnLoss: true,
	}
	g.PlayGame(getGuessFromCli)
}

func getGuessFromCli(lastGuess string, lastGuessStats string) string {
	scanner := bufio.NewScanner(os.Stdin)
	var guess string
	var isValid bool
	for {
		if scanner.Scan() {
			input := scanner.Text()
			isValid, guess = validateGuess(input)
			if !isValid {
				continue
			}
			// valid guess
			break
		}
	}
	return guess
}

func validateGuess(input string) (bool, string) {
	inpSplit := strings.Split(input, " ")
	if len(inpSplit) > 1 || len(inpSplit) == 0 || len(inpSplit[0]) != 5 {
		fmt.Println("Your guess must be exactly five letters")
		return false, ""
	}
	guess := inpSplit[0]

	// not sure I'm understanding slices vs arrays properly, but as far
	// as I can tell, this essentially just passes a reference to the
	// underlying array, and does not copy it, which is what I want.
	if !slices.Contains(VALID_GUESSES[:], guess) {
		fmt.Println("Invalid guess")
		return false, ""
	}

	return true, guess
}
