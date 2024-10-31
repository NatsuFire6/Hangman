package Hangman

import (
	"fmt"
)

func Jeu(word string, wordhelp string, life int, tab map[rune]bool, point int, difficultypoint int, aide string) int {
	pointAide := 1
	wordfind := make([]string, len(word))
	pass := false
	verif := false
	var letter string

	if word == "" {
		return point/2
	}
	if aide == "oui" {	
		pointAide = 2
		
	}
	if wordhelp != ""{						//permet de démarrer avec le mot d'aide
		for i := 0; i < len(word); i++ {
			wordfind[i] = string(wordhelp[i])
			if wordfind[i] == "_"{
				fmt.Print("_ ")
			}else{
				print(wordfind[i])
			}
		}
	}else{
		for i := 0; i < len(word); i++ {
			fmt.Print("_ ")
			wordfind[i] = "_"
		}
	}
	fmt.Println()
	for life >= 1 {
		fmt.Scan(&letter)
		if len(letter) >= 2 {
			fmt.Println("Vous ne devez entrer qu'une seul valeur.")
		} else if rune(letter) < 97 || rune(letter) > 122{ 
			fmt.Println("Veuillez entrer une lettre de l'alphabet en minuscule.")
		} else {
			for i := 0; i < len(word); i++ {
				if letter == string(word[i]) {
					wordfind[i] = letter
					for j := 0; j < len(wordfind); j++ {
						if wordfind[j] == string(word[j]) {
							verif = true
						} else {
							verif = false
							break
						}
					}
					if verif {
						DisplayMessage(verif, life)			
						fmt.Println(wordfind)				
						DisplayAlphabet(letter, tab)
						return (point + (difficultypoint * life)) / pointAide
					}
					pass = true
				}
			}

			if !pass && !verif{
				fmt.Print("\033[H\033[2J")
				life--
				DisplayHangman(life)
				DisplayMessage(verif, life)
				fmt.Println(wordfind )
				DisplayAlphabet(letter, tab)
			}else{
				fmt.Print("\033[H\033[2J")	//clear le terminal
				DisplayHangman(life)
				DisplayMessage(verif, life)
				fmt.Println(wordfind)
				DisplayAlphabet(letter, tab)
				pass = false
			}
		}
		if life == 0 {
			fmt.Printf("La réponse était : %s", string(word))
			break
		}
	}
	return point/2
}
