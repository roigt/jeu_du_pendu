package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadGuess() (guess string, err error) {
	valid := false

	for !valid {
		fmt.Printf("\nPlease enter your guess: ")
		guess, err = reader.ReadString('\n')
		if err != nil {
			return guess, err
		}
		guess = strings.TrimSpace(guess)
		if len(guess) != 1 { //on vérifie que la proposition est un seul caractère
			fmt.Printf("Invalid letter input. letter=%v, len=%v", guess, len(guess))
			continue
		}
		valid = true

	}
	return guess, nil
}
