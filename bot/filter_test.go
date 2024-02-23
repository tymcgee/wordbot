package bot

import (
	"strings"
	"testing"

	"github.com/tymcgee/wordbot/game"
)

// We know we have no n's based on the gray in index 3.
// So none of the remaining possible answers should have an n anywhere.
func TestFilterSingleGrays(t *testing.T) {
	// info: ⬛⬛🟩⬛⬛
	//       s p a n e
	gray := []GameInformation{
		{Letter: []rune("n")[0], Index: 3},
	}
	green := []GameInformation{
		{Letter: []rune("e")[0], Index: 2},
	}
	yellow := make([]GameInformation, 0)
	validGuesses := FilterGuesses(game.VALID_GUESSES[:], gray, yellow, green)

	for _, guess := range validGuesses {
		if strings.Contains(guess, "n") {
			t.Fatalf("Found guess [%v] which has an 'n' in the wrong spot (info is ⬛⬛🟩⬛⬛)", guess)
		}
	}
}

// Based on the gray e in index 3, we know we can't have any
// more e's. The green e should guarantee that all future guesses
// have an e in that position.
func TestFilterDoubleLetterOneGrayOneGreen(t *testing.T) {
	// info: ⬛⬛🟩⬛⬛
	//       s p e e d
	gray := []GameInformation{
		{Letter: []rune("e")[0], Index: 3},
	}
	green := []GameInformation{
		{Letter: []rune("e")[0], Index: 2},
	}
	yellow := make([]GameInformation, 0)
	validGuesses := FilterGuesses(game.VALID_GUESSES[:], gray, yellow, green)

	for _, guess := range validGuesses {
		if strings.Count(guess, "e") > 1 || string(guess[2]) != "e" {
			t.Fatalf("Found guess [%v] which has an 'e' which isn't at index 2 (info is ⬛⬛🟩⬛⬛)", guess)
		}
	}
}

// for a test case i saw fail:
// tafia  ⬛⬛⬛⬛⬛
// prise  ⬛🟩⬛🟩⬛
// brash  🟩🟩⬛🟩🟩
func TestFilterTwo(t *testing.T) {
	gray := []GameInformation{
		{Letter: []rune("t")[0], Index: 0},
		{Letter: []rune("a")[0], Index: 1},
		{Letter: []rune("f")[0], Index: 2},
		{Letter: []rune("i")[0], Index: 3},
		{Letter: []rune("a")[0], Index: 4},
	}
	green := make([]GameInformation, 0)
	yellow := make([]GameInformation, 0)
	validGuesses := FilterGuesses(game.VALID_GUESSES[:], gray, yellow, green)

	for _, guess := range validGuesses {
		if strings.ContainsAny(guess, "tafia") {
			t.Fatalf("Got guess [%v] when all of 'tafia' are gray", guess)
		}
	}
}
