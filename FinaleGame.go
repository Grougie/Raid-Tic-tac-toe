package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var grid = []string{"-", "-", "-", "-", "-", "-", "-", "-", "-"}
var currentPlayer string
var endGame = false

func printGrid(conn net.Conn) {
	fmt.Println("\n▬▬▬▬▬▬▬▬▬▬▬▬▬")
	fmt.Printf("| %s | %s | %s |      | 1 | 2 | 3 |\n", grid[0], grid[1], grid[2])
	fmt.Println("▬▬▬▬▬▬▬▬▬▬▬▬▬")
	fmt.Printf("| %s | %s | %s |      | 4 | 5 | 6 |\n", grid[3], grid[4], grid[5])
	fmt.Println("▬▬▬▬▬▬▬▬▬▬▬▬▬")
	fmt.Printf("| %s | %s | %s |      | 7 | 8 | 9 |\n", grid[6], grid[7], grid[8])
	fmt.Println("▬▬▬▬▬▬▬▬▬▬▬▬▬")
	fmt.Printf("\n")

	if conn != nil {
		gridRepresentation := fmt.Sprintf("\n▬▬▬▬▬▬▬▬▬▬▬▬▬\n| %s | %s | %s |      | 1 | 2 | 3 |\n▬▬▬▬▬▬▬▬▬▬▬▬▬\n| %s | %s | %s |      | 4 | 5 | 6 |\n▬▬▬▬▬▬▬▬▬▬▬▬▬\n| %s | %s | %s |      | 7 | 8 | 9 |\n▬▬▬▬▬▬▬▬▬▬▬▬▬\n\n",
			grid[0], grid[1], grid[2], grid[3], grid[4], grid[5], grid[6], grid[7], grid[8])

		conn.Write([]byte(gridRepresentation))
	}
}

func main() {
	var role string
	fmt.Print("Choose role (server/client): ")
	fmt.Scanln(&role)

	if role == "server" {
		Server()
	} else if role == "client" {
		var serverIP string
		fmt.Print("Enter server IP address: ")
		fmt.Scanln(&serverIP)
		Client(serverIP)
	} else {
		fmt.Println("Invalid role.")
	}
}

func Server() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port 8000")
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	choicePlayer()

	for !endGame {
		round(currentPlayer)
		endGameChecker()
		nextPlayer()

		sendGridToClient(conn)
		if endGame {
			break
		}

		receiveMoveFromClient(conn)
		endGameChecker()
		nextPlayer()
	}
	result()
}

func Client(IP string) {
	conn, err := net.Dial("tcp", IP+":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	choicePlayer()

	for !endGame {
		receiveMoveFromServer(conn)
		endGameChecker()
		nextPlayer()

		round(currentPlayer)
		endGameChecker()
		nextPlayer()

		sendGridToServer(conn)
	}
	result()
}

func sendGridToClient(conn net.Conn) {
	printGrid(conn)
}

func receiveMoveFromClient(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	index := parsePosition(strings.TrimSpace(message))
	if index >= 0 && index < 9 && grid[index] == "-" {
		grid[index] = currentPlayer
	}
}

func ClientPlayer(IP string) {
    conn, err := net.Dial("tcp", IP+":8000")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    choicePlayer()

    for !endGame {
        receiveMoveFromServer(conn)
        endGameChecker()
        nextPlayer()

        round(currentPlayer)
        endGameChecker()
        nextPlayer()

        sendGridToServer(conn)
    }
    result()
}

func receiveMoveFromServer(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	index := parsePosition(strings.TrimSpace(message))
	if index >= 0 && index < 9 && grid[index] == "-" {
		grid[index] = currentPlayer
	}
}

func sendGridToServer(conn net.Conn) {
	printGrid(conn)
}

func choicePlayer() {
	fmt.Println("☆━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━☆")
	fmt.Println("Welcome to the Tic Tac Toe game !")
	fmt.Println("☆━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━☆")
	fmt.Printf("\n")
	fmt.Println("The grid is composed of 9 boxes numbered from 1 to 9.")
	fmt.Println("To play, you must choose a box by entering the number corresponding to the box.")
	fmt.Println("The first player to align 3 boxes wins the game.")
	fmt.Printf("\n")
	fmt.Print("Please choose either a cross (X) or a circle (O) : ")
	fmt.Scan(&currentPlayer)

	for {
		currentPlayer = strings.ToUpper(currentPlayer)
		if currentPlayer == "X" {
			fmt.Println("You chose X.")
			break
		} else if currentPlayer == "O" {
			fmt.Println("You chose O.")
			break
		} else {
			fmt.Print("Please choose either (X) or (O) : ")
			fmt.Scan(&currentPlayer)
		}
	}
}

func round(player string) {
	fmt.Printf("It's the player's turn: %s\n", player)
	var pos string
	fmt.Print("Please select an empty space on the grid between 1 and 9 : ")
	fmt.Scan(&pos)

	valid := false
	for !valid {
		index := parsePosition(pos)
		if index >= 0 && index < 9 && grid[index] == "-" {
			valid = true
		} else {
			fmt.Printf("Player %s wins!\n", grid[0])
			fmt.Print("Error. Please select an empty space on the grid between 1 and 9 : ")
			fmt.Scan(&pos)
		}
	}

	index := parsePosition(pos)
	grid[index] = player
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
	if !endGame {
		return
	}
	printGrid(nil)
	if grid[0] == grid[1] && grid[1] == grid[2] && grid[2] != "-" ||
		grid[3] == grid[4] && grid[4] == grid[5] && grid[5] != "-" ||
		grid[6] == grid[7] && grid[7] == grid[8] && grid[8] != "-" ||
		grid[0] == grid[3] && grid[3] == grid[6] && grid[6] != "-" ||
		grid[1] == grid[4] && grid[4] == grid[7] && grid[7] != "-" ||
		grid[2] == grid[5] && grid[5] == grid[8] && grid[8] != "-" ||
		grid[0] == grid[4] && grid[4] == grid[8] && grid[8] != "-" ||
		grid[2] == grid[4] && grid[4] == grid[6] && grid[6] != "-" {
		fmt.Printf("Player %s wins!\n", grid[0])
	} else {
		fmt.Println("It's a tie!")
	}
}
