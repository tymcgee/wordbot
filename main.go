package main

import (
	"bufio"
	"os"

	"github.com/tymcgee/wordbot/game"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	game.PlayGame(scanner)
}
