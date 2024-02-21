package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

const (
	GREEN  = "🟩"
	YELLOW = "🟨"
	GRAY   = "⬛"
)

func PlayGame(inputScanner *bufio.Scanner) {
	answer := generateAnswer()
	guessNum := 0
	stats := ""

	fmt.Println("Guess the word!")
	for {
		var guess string
		var isValid bool
		if inputScanner.Scan() {
			input := inputScanner.Text()
			isValid, guess = validateGuess(input)
			if !isValid {
				continue
			}
		}

		colors := getColors([]rune(guess), []rune(answer))
		stats += colors + "\n"
		fmt.Println(colors)

		if guess == answer {
			fmt.Println("You win!")
			fmt.Printf("%d/6\n%s", guessNum+1, stats)
			break
		}

		if guessNum == 5 {
			fmt.Printf("The answer was [%s]\n", answer)
			fmt.Printf("x/6\n%s", stats)
			break
		}
		guessNum++
	}
}

func generateAnswer() string {
	answerIndex := rand.Intn(len(VALID_ANSWERS))
	// fmt.Printf("answer is [%s]\n", answer)
	return VALID_ANSWERS[answerIndex]
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
