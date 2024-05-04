package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Récupérer le plateau de Sudoku à partir des arguments donné.
	board := os.Args[1:]

	// Vérifier que le plateau est valide. S'il ne l'est pas, afficher "Error".
	if len(board) != 9 || !isValidBoard(board) {
		fmt.Println("Error")
		return
	}

	// Créer une copie du plateau pour vérifier s'il y a plusieurs solutions.
	boardCopy := make([]string, len(board))
	copy(boardCopy, board)

	// Essaye de résoudre le Sudoku. Sinon affiche "Error".
	if !solveSudoku(board, 0, 0) {
		fmt.Println("Error")
		return
	}

	// Vérifie s'il y a plusieurs solutions. Si c'est le cas, afficher "Error".
	if solveSudoku(boardCopy, 0, 0) && !areBoardsEqual(board, boardCopy) {
		fmt.Println("Error")
		return
	}

	// Afficher la solution.
	for _, row := range board {
		for _, num := range row {
			z01.PrintRune(num)
		}
		z01.PrintRune('\n')
	}
}

// Vérifie si deux grilles de Sudoku sont identiques.
func areBoardsEqual(board1, board2 []string) bool {
	for i := range board1 {
		if board1[i] != board2[i] {
			return false
		}
	}
	return true
}

// Vérifie si la grille de Sudoku est valide.
func isValidBoard(board []string) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' && !isValid(board, i, j) {
				return false
			}
		}
	}
	return true
}

// Vérifie si un chiffre est valide la ou il est.
func isValid(board []string, row, col int) bool {
	num := board[row][col]
	board[row] = board[row][:col] + "." + board[row][col+1:]
	for i := 0; i < 9; i++ {
		if board[i][col] == num || board[row][i] == num || board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	board[row] = board[row][:col] + string(num) + board[row][col+1:]
	return true
}

// Essaie de résoudre le Sudoku.
func solveSudoku(board []string, row, col int) bool {
	if col == 9 {
		row++
		col = 0
		if row == 9 {
			return true
		}
	}
	if board[row][col] != '.' {
		return solveSudoku(board, row, col+1)
	}
	for num := '1'; num <= '9'; num++ {
		board[row] = board[row][:col] + string(num) + board[row][col+1:]
		if isValid(board, row, col) && solveSudoku(board, row, col+1) {
			return true
		}
		board[row] = board[row][:col] + "." + board[row][col+1:]
	}
	return false
}
