package Hangman

import (
	"fmt"
	"os"
	"strings"
)

func DisplayHangman(viesRestantes int) {

	// Lire le fichier "hangman.txt" dans le dossier "text"
	file, err := os.ReadFile("text/hangman.txt")
	if err != nil {
		// Afficher un message d'erreur en cas de problème lors de la lecture du fichier
		fmt.Printf("Erreur lors de la lecture du fichier : %v", err)
	}
	// Séparer le contenu du fichier en différents éléments, en utilisant une virgule comme séparateur
	elements := strings.Split(string(file), ",")

	// Afficher un message spécial lorsque le joueur a toutes ses vies
	if viesRestantes == 10 {
		fmt.Println("Full Vies")
	}
	
	// Afficher l'état du pendu et le nombre de vies restantes pour chaque cas de viesRestantes
	if viesRestantes == 9 {
		fmt.Println(elements[0]) // Afficher l'état correspondant pour 9 vies restantes
		fmt.Printf("%v Vies Restantes\n", viesRestantes)
	}
	if viesRestantes == 8{
		fmt.Println(elements[1])
		fmt.Printf("%v Vies Restantes\n", viesRestantes)
	}
	if viesRestantes == 7{
		fmt.Println(elements[2])
		fmt.Printf("%v Vies Restantes\n", viesRestantes)
	}
	if viesRestantes == 6{
		fmt.Println(elements[3])
		fmt.Printf("%v Vies Restantes\n", viesRestantes)
	}
	if viesRestantes == 5{
		fmt.Println(elements[4])
		fmt.Printf("%v Vies Restantes\n", viesRestantes)
	}
	if viesRestantes == 4{
		fmt.Println(elements[5])
		fmt.Printf("%v Vies Restantes\n", viesRestantes)
	}
	if viesRestantes == 3{
		fmt.Println(elements[6])
		fmt.Printf("%v Vies Restantes\n", viesRestantes)
	}
	if viesRestantes == 2{
		fmt.Println(elements[7])
		fmt.Printf("%v Vies Restantes\n", viesRestantes)
	}
	if viesRestantes == 1{
		fmt.Println(elements[8])
		fmt.Printf("%v Vies Restantes\n", viesRestantes)
	}
	if viesRestantes == 0{
		fmt.Println(elements[9])
		fmt.Printf("%v Vies Restantes\n", viesRestantes)
	}
}
