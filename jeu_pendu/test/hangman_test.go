package test

import (
	"ExpertGo/jeu_pendu/hangman"
	"testing"
)

// test pour vérifier si la lettre devinée est dans le mot
func TestLetterInWord(t *testing.T) {
	word := []string{"b", "o", "b"}
	guess := "b"
	hasLetter := hangman.LetterInWord(guess, word)
	if !hasLetter {
		t.Errorf("LetterInWord(%s, %s) should be true", word, guess)
	}
}

func TestLetterNotInWord(t *testing.T) {
	word := []string{"b", "o", "b"}
	guess := "a"
	hasLetter := hangman.LetterInWord(guess, word)
	if hasLetter {
		t.Errorf("LetterInWord(%s, %s) should be false", word, guess)
	}
}

func TestInvalidWord(t *testing.T) {
	_, err := hangman.New(3, "")
	if err == nil {
		t.Errorf("Error should be returned when using an invalid word:%v", err)
	}
}

func TestGameGoodGuess(t *testing.T) {
	g, _ := hangman.New(3, "bob")
	g.MakeAGuess("b")
	valideState(t, "goodGuess", g.State)
}

func TestGameBadGuess(t *testing.T) {
	g, _ := hangman.New(3, "bob")
	g.MakeAGuess("m")
	valideState(t, "badGuess", g.State)
}

func TestGameAlreadyGuessed(t *testing.T) {
	g, _ := hangman.New(3, "bob")
	g.MakeAGuess("b")
	g.MakeAGuess("b")
	valideState(t, "alreadyGuess", g.State)
}

func TestGameWon(t *testing.T) {
	g, _ := hangman.New(3, "bob")
	g.MakeAGuess("b")
	g.MakeAGuess("o")
	valideState(t, "won", g.State)
}

func TestGameLost(t *testing.T) {
	g, _ := hangman.New(3, "bob")
	g.MakeAGuess("b")
	g.MakeAGuess("a")
	g.MakeAGuess("d")
	valideState(t, "lost", g.State)
}
func valideState(t *testing.T, expectedState string, actualState string) bool {
	if expectedState != actualState {
		t.Errorf("Expected: %s, Actual: %s", expectedState, actualState)
		return false
	}
	return true
}
