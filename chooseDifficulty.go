package Hangman

import (
	"fmt"
	"os"
)

func ChooseDifficulty(difficulty string, repete int) ([]byte, int) {
	if difficulty == "easy" {
		file, err := os.ReadFile("text/words.txt")
		if err != nil {
			fmt.Printf("Erreur lors de la lecture du fichier")
		}
		return file, 1
	} else if difficulty == "medium" {
		file, err := os.ReadFile("text/words2.txt")
		if err != nil {
			fmt.Printf("Erreur lors de la lecture du fichier")
		}
		return file, 2
	} else if difficulty == "hard" {
		file, err := os.ReadFile("text/words3.txt")
		if err != nil {
			fmt.Printf("Erreur lors de la lecture du fichier")
		}
		return file, 3
	} else if difficulty == "spécial" {
		file, err := os.ReadFile("text/words4.txt")
		if err != nil {
			fmt.Printf("Erreur lors de la lecture du fichier")
		}
		return file, 4
	} else {
		repete++
		print("\n")
		if repete == 3 {
			fmt.Println("\033[1m\033[3mTu commence à me courir sur le haricot, alors choisi une p**** de difficulté !\033[0m")
			return ChooseDifficulty(difficulty, repete)
		} else if repete == 6 {
			fmt.Println("\033[1m\033[3mWesh ! Nan mais là gros tu dose pas, j'ai juré tu vas choisir une difficulté !\033[0m")
			return ChooseDifficulty(difficulty, repete)
		} else if repete == 10 {
			fmt.Println("\033[1m\033[3mAlors là gros fait gaffe à ta maison parce que là je vais venir te buter toi et toutes ta famille dans ton lit ce soir !!!!!!!!\033[0m")
			return nil, 0
		} else {
			fmt.Println("Vous n'avez pas entrée la bonne difficulté, veuillez choisir : \033[1measy\033[0m, \033[1mmedium\033[0m, \033[1mhard\033[0m ou \033[1mspécial\033[0m")
			return ChooseDifficulty(difficulty, repete)
		}
	}
}
