package main

import (
	"fmt"
	"strings"
)

var grille = []string{"-", "-", "-",
	"-", "-", "-",
	"-", "-", "-"}

var joueurActuel string
var gagnant string
var finJeu = false

func choixJoueur() {
	fmt.Print("Veuillez choisir soit une croix (X), soit un rond (O) : ")
	fmt.Scan(&joueurActuel)

	for {
		joueurActuel = strings.ToUpper(joueurActuel)
		if joueurActuel == "X" {
			fmt.Println("Vous avez choisi X. Le joueur 2 aura O")
			break
		} else if joueurActuel == "O" {
			fmt.Println("Vous avez choisi O. Le joueur 2 aura X")
			break
		} else {
			fmt.Print("Veuillez choisir soit (X) soit (O) : ")
			fmt.Scan(&joueurActuel)
		}
	}
}




