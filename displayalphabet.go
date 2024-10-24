package Hangman

import (
	"fmt"
	"strings"
)

func DisplayAlphabet(input string, usedLetters map[rune]bool) {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	input = strings.ToLower(input)
	

	// Marquer les lettres utilisées
	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			usedLetters[char] = true
		}
	}

	var result strings.Builder

	// Parcourir l'alphabet et ajouter les couleurs
	for _, letter := range alphabet {
		if usedLetters[letter] {
			result.WriteString("\033[32m") // Vert
		} else {
			result.WriteString("\033[31m") // Rouge
		}
		result.WriteString(string(letter))
		result.WriteString("\033[0m") // Réinitialiser la couleur
		result.WriteString(" ")
	}

	fmt.Println(result.String())
}
