package game

import "testing"

func TestColorsOne(t *testing.T) {
	answer := "blink"
	guess := "links"
	expected := YELLOW + YELLOW + YELLOW + YELLOW + GRAY
	testColors(answer, guess, expected, t)
}

func TestColorsTwo(t *testing.T) {
	answer := "bleep"
	guess := "scour"
	expected := GRAY + GRAY + GRAY + GRAY + GRAY
	testColors(answer, guess, expected, t)
}

func TestColorsThree(t *testing.T) {
	answer := "bleep"
	guess := "blear"
	expected := GREEN + GREEN + GREEN + GRAY + GRAY
	testColors(answer, guess, expected, t)
}

func TestColorsFour(t *testing.T) {
	answer := "slide"
	guess := "sleep"
	expected := GREEN + GREEN + YELLOW + GRAY + GRAY
	testColors(answer, guess, expected, t)
}

func TestColorsFive(t *testing.T) {
	answer := "bleep"
	guess := "seene"
	expected := GRAY + YELLOW + GREEN + GRAY + GRAY
	testColors(answer, guess, expected, t)
}

func TestColorsSix(t *testing.T) {
	answer := "bleep"
	guess := "senne"
	expected := GRAY + YELLOW + GRAY + GRAY + YELLOW
	testColors(answer, guess, expected, t)
}

func TestColorsSeven(t *testing.T) {
	answer := "hotly"
	guess := "holly"
	expected := GREEN + GREEN + GRAY + GREEN + GREEN
	testColors(answer, guess, expected, t)
}

func TestColorsEight(t *testing.T) {
	answer := "scene"
	guess := "peele"
	expected := GRAY + GRAY + GREEN + GRAY + GREEN
	testColors(answer, guess, expected, t)
}

func TestColorsNine(t *testing.T) {
	answer := "might"
	guess := "tight"
	expected := GRAY + GREEN + GREEN + GREEN + GREEN
	testColors(answer, guess, expected, t)
}

func TestColorsTen(t *testing.T) {
	answer := "might"
	guess := "trite"
	expected := YELLOW + GRAY + YELLOW + GRAY + GRAY
	testColors(answer, guess, expected, t)
}

func TestColorsEleven(t *testing.T) {
	answer := "tteee"
	guess := "eexee"
	expected := YELLOW + GRAY + GRAY + GREEN + GREEN
	testColors(answer, guess, expected, t)
}

func testColors(answer string, guess string, expected string, t *testing.T) {
	actual := getColors([]rune(guess), []rune(answer))
	if actual != expected {
		t.Fatalf("Guessed [%v] [%v] with answer [%v] [%v]", guess, actual, answer, expected)
	}
}
