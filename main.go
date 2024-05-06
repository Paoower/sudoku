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
		fmt.Println("Erreur la grille n'est pas valide")
		return
	}

	// Essaye de résoudre le Sudoku.
	if !solveSudoku(board, 0, 0, false) {
		fmt.Println("Erreur la grille impossible à résoudre.")
		return
	}

	for _, arg := range board {
		for _, c := range arg {
			if c != '.' {
				// Créer une copie du sudoku pour vérifier s'il y a plusieurs solutions.
				boardCopy := make([]string, len(board))
				copy(boardCopy, board)
				// Vérifie s'il y a plusieurs solutions.
				if solveSudoku(boardCopy, 0, 0, true) && !areBoardsEqual(board, boardCopy) {
					fmt.Println("Erreur la grille a plusieurs solution")
				}
			}
		}
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
	//ligne
	for i := 0; i < 9; i++ {
		//colonne
		for j := 0; j < 9; j++ {
			//si n'est pas un . et emplacement du chiffre invalide renvoie false
			if board[i][j] != '.' && !isValid(board, i, j) {
				return false
			}
		}
	}
	return true
}

// Vérifie si un chiffre est valide là où il se trouve.
func isValid(board []string, row, col int) bool {
	// Vérifier la longueur de chaque ligne
	for _, row := range board {
		if len(row) != 9 {
			return false
		}
	}
	//stock le chiffre
	num := board[row][col]
	if num < '1' || num > '9' {
		return false
	}
	//remplace par un .
	board[row] = board[row][:col] + "." + board[row][col+1:]
	for i := 0; i < 9; i++ {
		//si l'emplacement contient deja le chiffre (stocké dans num) renvoie false
		if board[i][col] == num || board[row][i] == num || board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	//remet le chiffre
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
	//Si case n'est pas vide , appel récursif
	if board[row][col] != '.' {
		return solveSudoku(board, row, col+1, reverse)
	}
	//En fct du bool reverse définit le sens de résolution
	start := '1'
	end := '9'
	step := 1
	if reverse {
		start = '9'
		end = '1'
		step = -1
	}
	//Test chaque chiffre pour trouver celui qui convient
	for num := start; ; num = rune(int(num) + step) {
		// arret si on depasse end
		if (step > 0 && num > end) || (step < 0 && num < end) {
			break
		}
		// Place le chiffre num actuel dans la grille puis verifie si il convient
		board[row] = board[row][:col] + string(num) + board[row][col+1:]
		if isValid(board, row, col) && solveSudoku(board, row, col+1, reverse) {
			return true
		}
		//retire le chiffre si il ne convient pas
		board[row] = board[row][:col] + "." + board[row][col+1:]
	}
	//si ne trouve pas de solution renvoie false
	return false
}
