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
	takenYellows := make([]int, 5)
	for i, guessRune := range guess {
		color := ""
		for j, answerRune := range answer {

			if i == j && guessRune == answerRune {
				color = GREEN
				break
			}

			// each letter in the answer can only map to one of either gray, yellow, or green.
			// if it's green, it can't also map to a yellow.
			// right now we're either gray or yellow.
			if guessRune == answerRune && !slices.Contains(takenYellows, j) {
				// look at our guess at this index to see if it would be green
				if guess[j] != answerRune {
					// our guess is not green at this letter, so we can use it for yellow
					color = YELLOW
					takenYellows = append(takenYellows, j)
					break
				}
			}
		}
		if color == "" {
			color = GRAY
		}
		colors += color
	}
	return colors
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
