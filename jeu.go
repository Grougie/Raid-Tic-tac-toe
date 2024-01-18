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

func choicePlayer() {
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

func printGrid() {
	fmt.Println("\n-------------")
	fmt.Printf("| %s | %s | %s |      | 1 | 2 | 3 |\n", grille[0], grille[1], grille[2])
	fmt.Println("-------------")
	fmt.Printf("| %s | %s | %s |      | 4 | 5 | 6 |\n", grille[3], grille[4], grille[5])
	fmt.Println("-------------")
	fmt.Printf("| %s | %s | %s |      | 7 | 8 | 9 |\n", grille[6], grille[7], grille[8])
	fmt.Printf("\n")
}

func round(joueur string) {
	fmt.Printf("C'est le tour du joueur: %s\n", joueur)
	var pos string
	fmt.Print("Veuillez sélectionner un espace vide sur la grille entre 1 et 9 : ")
	fmt.Scan(&pos)

	valide := false
	for !valide {
		for _, p := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"} {
			if pos == p {
				break
			}
		}
		index := parsePosition(pos)
		if index >= 0 && index < 9 && grille[index] == "-" {
			valide = true
		} else {
			fmt.Print("Vous ne pouvez pas accéder à cette position. Veuillez sélectionner un espace vide sur la grille entre 1 et 9 : ")
			fmt.Scan(&pos)
		}
	}

	index := parsePosition(pos)
	grille[index] = joueur
	printGrid()
}

func parsePosition(pos string) int {
	index := -1
	fmt.Sscanf(pos, "%d", &index)
	return index - 1
}

func endGameChecker() {
	victory()
	egalityGame()
}

func victory() {
	if grille[0] == grille[1] && grille[1] == grille[2] && grille[2] != "-" ||
		grille[3] == grille[4] && grille[4] == grille[5] && grille[5] != "-" ||
		grille[6] == grille[7] && grille[7] == grille[8] && grille[8] != "-" ||
		grille[0] == grille[3] && grille[3] == grille[6] && grille[6] != "-" ||
		grille[1] == grille[4] && grille[4] == grille[7] && grille[7] != "-" ||
		grille[2] == grille[5] && grille[5] == grille[8] && grille[8] != "-" ||
		grille[0] == grille[4] && grille[4] == grille[8] && grille[8] != "-" ||
		grille[2] == grille[4] && grille[4] == grille[6] && grille[6] != "-" {
		finJeu = true
		gagnant = grille[0]
	}
}

func egalityGame() {
	if !strings.Contains(strings.Join(grille, ""), "-") {
		finJeu = true
	}
}

func nextPlayer() {
	if joueurActuel == "X" {
		joueurActuel = "O"
	} else {
		joueurActuel = "X"
	}
}

func result() {
	if gagnant == "X" || gagnant == "O" {
		fmt.Printf("Le joueur : %s a gagné\n", gagnant)
	} else {
		fmt.Println("Match nul")
	}
}

func game() {
	choicePlayer()
	printGrid()
	for !finJeu {
		round(joueurActuel)
		endGameChecker()
		nextPlayer()
	}
	result()
}

func main() {
	game()
}


