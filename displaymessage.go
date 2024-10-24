package Hangman

import (
	"fmt"
)

func DisplayMessage(motComplete bool, life int) {
	// Vérifier si le mot est complet (le joueur a gagné)
	if motComplete {
		// Afficher un message de félicitations
		fmt.Println("Félicitations !")
	} else if !motComplete && life == 0 {
		// Si le mot est incomplet et que le joueur n'a plus de vies, il a perdu
		fmt.Println("Vous avez perdu !")
	}
}
