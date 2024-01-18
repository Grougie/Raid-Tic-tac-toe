package main

import (
	"fmt"
	"strings"
)

var grid = []string{"-", "-", "-",
	"-", "-", "-",
	"-", "-", "-"}

var currentPlayer string
var winner string
var endGame = false

func choicePlayer() {
	fmt.Print("Please choose either a cross (X) or a circle (O) : ")
	fmt.Scan(&currentPlayer)

	for {
		currentPlayer = strings.ToUpper(currentPlayer)
		if currentPlayer == "X" {
			fmt.Println("You chose X. Player 2 will have O")
			break
		} else if currentPlayer == "O" {
			fmt.Println("You chose O. Player 2 will have X")
			break
		} else {
			fmt.Print("Please choose either (X) or (O) : ")
			fmt.Scan(&currentPlayer)
		}
	}
}

func printGrid() {
	fmt.Println("\n-------------")
	fmt.Printf("| %s | %s | %s |      | 1 | 2 | 3 |\n", grid[0], grid[1], grid[2])
	fmt.Println("-------------")
	fmt.Printf("| %s | %s | %s |      | 4 | 5 | 6 |\n", grid[3], grid[4], grid[5])
	fmt.Println("-------------")
	fmt.Printf("| %s | %s | %s |      | 7 | 8 | 9 |\n", grid[6], grid[7], grid[8])
	fmt.Printf("\n")
}

func round(player string) {
	fmt.Printf("It's the player's turn: %s\n", player)
	var pos string
	fmt.Print("Please select an empty space on the grid between 1 and 9 : ")
	fmt.Scan(&pos)

	valide := false
	for !valide {
		for _, p := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"} {
			if pos == p {
				break
			}
		}
		index := parsePosition(pos)
		if index >= 0 && index < 9 && grid[index] == "-" {
			valide = true
		} else {
			fmt.Print("Error, You cannot access this position. Please select an empty space on the grid between 1 and 9 : ")
			fmt.Scan(&pos)
		}
	}

	index := parsePosition(pos)
	grid[index] = player
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
	if grid[0] == grid[1] && grid[1] == grid[2] && grid[2] != "-" ||
		grid[3] == grid[4] && grid[4] == grid[5] && grid[5] != "-" ||
		grid[6] == grid[7] && grid[7] == grid[8] && grid[8] != "-" ||
		grid[0] == grid[3] && grid[3] == grid[6] && grid[6] != "-" ||
		grid[1] == grid[4] && grid[4] == grid[7] && grid[7] != "-" ||
		grid[2] == grid[5] && grid[5] == grid[8] && grid[8] != "-" ||
		grid[0] == grid[4] && grid[4] == grid[8] && grid[8] != "-" ||
		grid[2] == grid[4] && grid[4] == grid[6] && grid[6] != "-" {
		endGame = true
		winner = grid[0]
	}
}

func egalityGame() {
	if !strings.Contains(strings.Join(grid, ""), "-") {
		endGame = true
	}
}

func nextPlayer() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}

func result() {
	if winner == "X" || winner == "O" {
		fmt.Printf("The player : %s won!!\n", winner)
	} else {
		fmt.Println("Egality game")
	}
}

func game() {
	choicePlayer()
	printGrid()
	for !endGame {
		round(currentPlayer)
		endGameChecker()
		nextPlayer()
	}
	result()
}