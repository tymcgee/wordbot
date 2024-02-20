package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strings"
)

const (
	GREEN  = "🟩"
	YELLOW = "🟨"
	GRAY   = "⬛"
)

func main() {
	answerIndex := rand.Intn(len(VALID_ANSWERS))
	answer := VALID_ANSWERS[answerIndex]
	fmt.Printf("answer is [%s]\n", answer)

	scanner := bufio.NewScanner(os.Stdin)

	guessNum := 0
	for {
		var guess string
		var isValid bool
		if scanner.Scan() {
			input := scanner.Text()
			isValid, guess = validateGuess(input)
			if !isValid {
				continue
			}
		}

		fmt.Printf("guess %v is %v\n", guessNum, guess)

		if guessNum == 5 {
			fmt.Print("YOU LOSE\n")
			break
		}
		guessNum++
	}
}

func getColors(guess []rune, answer []rune) string {
	colors := ""
	runningAnswer := string(answer)
	for i := 0; i < len(guess); i++ {
		guessRune := guess[i]
		answerRune := answer[i]
		if guessRune == answerRune {
			colors += GREEN
			runningAnswer = strings.Replace(runningAnswer, string(guessRune), "", 1)
			continue
		}
		if strings.ContainsRune(runningAnswer, guessRune) {
			colors += YELLOW
			runningAnswer = strings.Replace(runningAnswer, string(guessRune), "", 1)
			continue
		}
		colors += GRAY
	}
	return colors
}

func validateGuess(input string) (bool, string) {
	inpSplit := strings.Split(input, " ")
	if len(inpSplit) > 1 || len(inpSplit) == 0 || len(inpSplit[0]) != 5 {
		fmt.Println("Guess must be exactly five letters")
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
