package dictionary

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

var words = make([]string, 0, 50)

func Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	//
	if err := scanner.Err(); err != nil { //erreur de scanner
		return err
	}
	return nil
}

func PickWord() string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return words[r.Intn(len(words))]
}
