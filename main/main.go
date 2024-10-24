package main

import (
	"Hangman"
	"fmt"
	"os"
)

func main() {
	var point int
	fmt.Println("bienvenue sur le jeu du pendu (Hangman).")
	start(point)
	
}

func start (point int){
	var life int = 10
	var faitchier int = 0
	tab := make(map[rune]bool)
	aide := ""
	wordhelp := ""

	fmt.Println("Veuillez choisir une difficulté : \033[1measy\033[0m, \033[1mmedium\033[0m, \033[1mhard\033[0m, \033[1mspécial\033[0m")
	difficulty, difficultypoint := Hangman.ChooseDifficulty(faitchier)
	word := Hangman.Aleatoire(difficulty)

	fmt.Println("Voulez-vous de l'aide ? oui/non")
	fmt.Scan(&aide)
	if aide == "oui" {
		wordhelp = Hangman.DisplayHalfWord(word)
	}else{
		fmt.Println("D'accord, vous naurez pas d'aide")
	}
	point = Hangman.Jeu(word, wordhelp, life, tab, point, difficultypoint, aide) //cacher et vérifier la lettre
	fmt.Print("\n")
	fmt.Printf("\nvous avez %v point\n", point)
	PlayAgain(point)
}

func PlayAgain(point int) {
	answer := ""
	
	fmt.Println("Voulez-vous rejouez ?")
	fmt.Scan(&answer)
	if answer == "oui" {
		start(point)
	} else if answer == "non" {
		fmt.Println("\033[31m\033[1mOK😒😤😭😡!!!\033[0m")
		fmt.Println("Votre score était de", point)
		os.Exit(0)
	} else {
		fmt.Println("\nVous n'avez pas répondu correctement, veuillez entrer \033[1moui\033[0m ou \033[1mnon\033[m .")
		PlayAgain(point)
	}
}
