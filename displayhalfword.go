package Hangman

import (
    "math/rand"
    "time"
)

func DisplayHalfWord(word string) string {
    runes := []rune(word)
    length := len(runes)
    maskedWord := make([]rune, length)

    // Masquer toutes les lettres au début
    for i := 0; i < length; i++ {
        maskedWord[i] = '_'
    }

    // Calculer le nombre de lettres à révéler
    revealCount := (length / 2) - 1
    revealedIndices := make(map[int]bool)

    rand.Seed(time.Now().UnixNano())

    // Révéler aléatoirement des lettres jusqu'à ce qu'on atteigne le revealCount
    for i := 0; i < revealCount; {
        index := rand.Intn(length) // Choisir un index aléatoire
        if !revealedIndices[index] {
            maskedWord[index] = runes[index]
            revealedIndices[index] = true
            i++
        }
    }

    return string(maskedWord)
}