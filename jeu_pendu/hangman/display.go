package hangman

import (
	"fmt"
)

func DrawWelcome() {

	fmt.Println(`
			   __    __                                                                 
		|  \  |  \                                                                
		| $$  | $$  ______   _______    ______   ______ ____    ______   _______  
		| $$__| $$ |      \ |       \  /      \ |      \    \  |      \ |       \ 
		| $$    $$  \$$$$$$\| $$$$$$$\|  $$$$$$\| $$$$$$\$$$$\  \$$$$$$\| $$$$$$$\
		| $$$$$$$$ /      $$| $$  | $$| $$  | $$| $$ | $$ | $$ /      $$| $$  | $$
		| $$  | $$|  $$$$$$$| $$  | $$| $$__| $$| $$ | $$ | $$|  $$$$$$$| $$  | $$
		| $$  | $$ \$$    $$| $$  | $$ \$$    $$| $$ | $$ | $$ \$$    $$| $$  | $$
		 \$$   \$$  \$$$$$$$ \$$   \$$ _\$$$$$$$ \$$  \$$  \$$  \$$$$$$$ \$$   \$$
									  |  \__| $$                                  
									   \$$    $$                                  
										\$$$$$$                                   
    `)

}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)

}
func drawTurns(l int) {
	var draw string

	switch l {

	case 8:
		draw = `
		 
		 
		 
		`

	case 7:
		draw = `
		|
		|
		|
		`

	case 6:
		draw = `
		|
		|
		|____
		`

	case 5:
		draw = `
		|----
		|
		|____
		`

	case 4:
		draw = `
		|----
		|   |
		|____
		`

	case 3:
		draw = `
		|----
		|   |
		|   O
		|____
		`

	case 2:
		draw = `
		|----
		|   |
		|   O
		|  /|\
		|____
		`

	case 1:
		draw = `
		|----
		|   |
		|   O
		|  /|\
		|  /
		|____
		`

	case 0:
		draw = `
		|----
		|   |
		|   O
		|  /|\
		|  / \
		|____
		`
	default:
		draw = ""
	}

	fmt.Println(draw)
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed:")
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Print("Good guess!")
	case "alreadyGuess":
		fmt.Printf("Already guess! Letter:%v ", guess)
	case "badGuess":
		fmt.Printf("Bad guess! Letter:%v", guess)
	case "lost":
		fmt.Printf("You lost :(! The word was:")
		drawLetters(g.Letters)
	case "won":
		fmt.Printf("You won! The word was:")
		drawLetters(g.Letters)

	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}
