package main

import (
	"ExpertGo/jeu_pendu/dictionary"
	"ExpertGo/jeu_pendu/hangman"
	"fmt"
	"os"
)

func main() {

	err := dictionary.Load("words.txt")
	if err != nil {
		fmt.Println("Could not load dictionary", err)
	}

	g, err := hangman.New(8, dictionary.PickWord())
	if err != nil {
		fmt.Println(err)
	}

	hangman.DrawWelcome()

	guess := ""
	for {
		hangman.Draw(g, guess)
		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read guess:%v ", err)
			os.Exit(1)
		}
		guess = l //la lettre devinée
		g.MakeAGuess(guess)
		fmt.Printf("Rest of attempts:%v\n", g.TurnsLeft)

	}

}
