package Hangman

import (
	"math/rand"
	"strings"
	"time"
)

func Aleatoire(file []byte) string {
	var randomWord string
	if file != nil {
		words := strings.Fields(string(file))

		// Initialiser le générateur de nombres aléatoires
		rand.Seed(time.Now().UnixNano()) 	 //do with chatgpt for learn new command and import

		// Choisir un mot aléatoire
		randomIndex := rand.Intn(len(words)) //do with chatgpt for learn new command and import
		randomWord = words[randomIndex]      //do with chatgpt for learn new command and import
		return randomWord
	} else {
		return ""
	}

}
