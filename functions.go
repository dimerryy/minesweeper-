package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func intro() {
	fmt.Println("Hello, this is the game Minesweeper!")
	fmt.Println("Minesweeper rules are very simple. The board is divided into cells, with mines randomly distributed. To win, you need to open all the cells. The number on a cell shows the number of mines adjacent to it. Using this information, you can determine cells that are safe, and cells that contain mines. Cells suspected of being mines can be marked with a flag using the right mouse button.")
}

func chooseMode() int {
	var mode int
	fmt.Println("Choose a mode:")
	fmt.Println("1. Generate a random map")
	fmt.Println("2. Enter a custom map")
	fmt.Println("Enter your choice: ")
	for {
		fmt.Scanf("%d", &mode)
		if mode == 1 || mode == 2 {
			break
		}
		fmt.Println("Enter only numbers 1 or 2: ")
	}
	return mode
}

func readTwoNums() (int, int) {
	var n1, n2 string
	var input string
	var r rune
	fmt.Scanf("%c", &r)
	for r != '\n' {
		input += string(r)
		fmt.Scanf("%c", &r)
	}
	if len(input) == 0 {
		return -1, -1
	}
	space := false
	for i, char := range input {
		if char == ' ' {
			if i == 0 || space {
				return -1, -1
			}
			space = true
		}
		if char >= '0' && char <= '9' {
			if !space {
				n1 += string(char)
			} else {
				n2 += string(char)
			}
		} else {
			return -1, -1
		}
	}

	if len(n2) == 0 {
		return -1, -1
	}

	k1, _ := strconv.Atoi(n1)
	k2, _ := strconv.Atoi(n2)
	return k1, k2
}
func randomMap(matrix *[][]int) [][]int {
	for i := 1; i < height; i++ {
		for j := 1; j < width; j++ {
			r := rand.Intn(6)
			if r == 0 {
				(*matrix)[i][j] == -1
			}
		}
	}
}
