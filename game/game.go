package game

import (
	"fmt"
	"math/rand"
	"slices"
)

const (
	GREEN  = "🟩"
	YELLOW = "🟨"
	GRAY   = "⬛"
)

type Game struct {
	ShowStats        bool
	ShowOngoingStats bool
	ShowGuesses      bool
	ShowIntro        bool
}

func (g *Game) PlayGame(getGuess func(lastGuess string, lastGuessStats string) string) {
	answer := generateAnswer()
	guessNum := 0
	stats := ""
	colors := ""
	guess := ""

	if g.ShowIntro {
		fmt.Println("Guess the word!")
	}
	for {
		guess = getGuess(guess, colors)

		colors = getColors([]rune(guess), []rune(answer))
		stats += colors + "\n"

		if g.ShowGuesses {
			fmt.Printf("%s  ", guess)
		}

		if g.ShowOngoingStats {
			fmt.Println(colors)
		}

		if guess == answer {
			fmt.Println("You win!")
			if g.ShowStats {
				fmt.Printf("%d/6\n%s", guessNum+1, stats)
			}
			break
		}

		if guessNum == 5 {
			fmt.Printf("The answer was [%s]\n", answer)
			if g.ShowStats {
				fmt.Printf("x/6\n%s", stats)
			}
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
