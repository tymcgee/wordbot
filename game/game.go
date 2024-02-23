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

type Results struct {
	Won     bool
	Guesses int
}

func (g *Game) PlayGameWithAnswer(answer string, getGuess func(lastGuess string, lastGuessStats string) string) Results {
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
			return Results{
				Won:     true,
				Guesses: guessNum,
			}
		}

		if guessNum == 5 {
			fmt.Printf("The answer was [%s]\n", answer)
			if g.ShowStats {
				fmt.Printf("x/6\n%s", stats)
			}
			return Results{
				Won:     false,
				Guesses: guessNum,
			}
		}
		guessNum++
	}
}

func (g *Game) PlayGame(getGuess func(lastGuess string, lastGuessStats string) string) Results {
	answer := generateAnswer()
	return g.PlayGameWithAnswer(answer, getGuess)
}

func generateAnswer() string {
	answerIndex := rand.Intn(len(VALID_ANSWERS))
	// fmt.Printf("answer is [%v]\n", VALID_ANSWERS[answerIndex])
	return VALID_ANSWERS[answerIndex]
}

func getColors(guess []rune, answer []rune) string {
	colors := ""
	takenYellows := make([]int, 0)
	for i, guessRune := range guess {
		color := ""
		// early exit if it's green
		if answer[i] == guessRune {
			colors += GREEN
			continue
		}
		for j, answerRune := range answer {
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
