pakage main

func main() {
	if resSudoku() {
		fmt.Println("Solution du Sudoku")
	} else {
		fmt.Println("Pas de solution trouvée.")
	}
}