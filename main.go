package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Récupérer la grille  de Sudoku à partir des arguments donné.
	board := os.Args[1:]
	// Vérifier que le sudoku est valide.
	if len(board) != 9 || !isValidBoard(board) {
		fmt.Println("Erreur la grille n'est pas de bonne taille.")
		return
	}
	// Créer une copie du sudoku pour vérifier s'il y a plusieurs solutions.
	boardCopy := make([]string, len(board))
	copy(boardCopy, board)
	// Essaye de résoudre le Sudoku.
	if !solveSudoku(board, 0, 0, false) {
		fmt.Println("Erreur la grille impossible à résoudre.")
		return
	}
	// Vérifie s'il y a plusieurs solutions.
	if solveSudoku(boardCopy, 0, 0, true) && !areBoardsEqual(board, boardCopy) {
		fmt.Println("Erreur la grille a plusieurs solution")
		return
	}
	// Affiche la solution.
	for _, row := range board {
		for _, num := range row {
			z01.PrintRune(num)
		}
		z01.PrintRune('\n')
	}
}

// Vérifie si les deux grilles de Sudoku sont identiques.
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

// Vérifie si un chiffre est valide là où il se trouve.
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
func solveSudoku(board []string, row, col int, reverse bool) bool {
	if col == 9 {
		row++
		col = 0
		if row == 9 {
			return true
		}
	}
	if board[row][col] != '.' {
		return solveSudoku(board, row, col+1, reverse)
	}
	start := '1'
	end := '9'
	step := 1
	if reverse {
		start = '9'
		end = '1'
		step = -1
	}
	for num := start; ; num = rune(int(num) + step) {
		if (step > 0 && num > end) || (step < 0 && num < end) {
			break
		}
		board[row] = board[row][:col] + string(num) + board[row][col+1:]
		if isValid(board, row, col) && solveSudoku(board, row, col+1, reverse) {
			return true
		}
		board[row] = board[row][:col] + "." + board[row][col+1:]
	}
	return false
}
