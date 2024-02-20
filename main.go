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
		if scanner.Scan() {
			inp := scanner.Text()
			inpSplit := strings.Split(inp, " ")
			if len(inpSplit) > 1 || len(inpSplit) == 0 || len(inpSplit[0]) != 5 {
				fmt.Println("Guess must be exactly five letters")
				continue
			}
			guess = inpSplit[0]

			// not sure I'm understanding slices vs arrays properly, but as far
			// as I can tell, this essentially just passes a reference to the
			// underlying array, and does not copy it, which is what I want.
			if !slices.Contains(VALID_GUESSES[:], guess) {
				fmt.Println("Invalid guess")
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
