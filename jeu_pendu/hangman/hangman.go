package hangman

import (
	"errors"
	"strings"
)

type Game struct {
	State        string   //Game state
	Letters      []string // Letters in the word to find
	FoundLetters []string //Good guesses
	UsedLetters  []string //USed Letters
	TurnsLeft    int      // Remaining attempts
}

// Fonction qui initialise un jeu
func New(turns int, word string) (*Game, error) {
	if len(word) < 3 {
		return nil, errors.New("word must have at least 3 characters")
	}

	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}

	game := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}

	return game, nil
}

// Fonction qui vérifie le choix du joueur et décide s'il a gagné, perdu ou peu continuer à jouer
func (game *Game) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)
	if LetterInWord(guess, game.UsedLetters) {
		game.State = "alreadyGuess"

	} else if LetterInWord(guess, game.Letters) {
		game.TurnsLeft--
		game.State = "goodGuess"
		game.RevealLetter(guess)

		if hasWon(game.Letters, game.FoundLetters) {
			game.State = "won"
		}
	} else {

		if game.isEndOfGame(game.TurnsLeft) {
			game.State = "lost"
		} else {
			game.State = "badGuess"
		}

	}
}

// afficher les lettres trouvées dans leurs positions
func (game *Game) RevealLetter(guess string) {
	game.UsedLetters = append(game.UsedLetters, guess)
	for i, l := range game.Letters {
		if l == guess {
			game.FoundLetters[i] = strings.ToUpper(guess)
		}
	}
}

// vérifie si un joueur a gagné
func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}

// vérifier si la lettre choisie par le joueur est dans le mot a deviné
func LetterInWord(guess string, letters []string) bool {
	for _, letter := range letters {
		if guess == letter {
			return true
		}
	}
	return false
}

// Vérifie si un jeu est terminé ou pas
func (game *Game) isEndOfGame(turns int) bool {
	game.TurnsLeft--
	if game.TurnsLeft > 0 { //on vérifie si le nombre de tourne n'est pas terminé
		return false
	}
	game.TurnsLeft = 0 //réinitialise le nombre de tentative
	return true
}
