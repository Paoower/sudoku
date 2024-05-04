package main

import "fmt"
		"os"

func main() {
	board := os.Args[1:]
	if !isValidBoard(board) {
		fmt.Println("Error")
		return
	}
	if resSudoku() {
		fmt.Println("Solution du Sudoku")
	} else {
		fmt.Println("Pas de solution trouvée.")
	}
}

// vérifie si le plateau de Sudoku de base est valide.
func isValidBoard(board []string) bool {
	// ligne
	for i := 0; i < 9; i++ {
		// colonne 
		for j := 0; j < 9; j++ {
			// La case a un chiffre mais pas valide
			if board[i][j] != '.' && !isValid(board, i, j) {
				return false
			}
		}
	}
	return true
}

// vérifie si un chiffre est valide la ou il se trouve
func isValid(board []string, row, col int) bool {
	// Récupére le chiffre 
	num := board[row][col]
	// Remplace temporairement par un point
	board[row] = board[row][:col] + "." + board[row][col+1:]
	// Parcours chaque case (ligne, colonne et 3x3)
	for i := 0; i < 9; i++ {
		// Si existe deja
		if board[i][col] == num || board[row][i] == num || board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	// Si n'existe pas (et replace le chiffre)
	board[row] = board[row][:col] + string(num) + board[row][col+1:]
	return true
}
