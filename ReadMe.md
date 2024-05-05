# Explication du code Sudoku

## Fonction `main`

C'est le début du programme. Elle prend le Sudoku à partir des arguments donnés. Les arguments sont les lignes du Sudoku. Elle vérifie ensuite si le Sudoku est bon avec la fonction `isValidBoard`. Si le Sudoku n'est pas bon, elle affiche "Error" et arrête le programme. Sinon, elle essaie de résoudre le Sudoku avec la fonction `solveSudoku`. Si le Sudoku ne peut pas être résolu, elle affiche une erreur et arrête le programme. Sinon, elle affiche la solution.

## Fonction `areBoardsEqual`

Cette fonction prend deux Sudokus et vérifie s'ils sont pareils. Elle utilise une boucle `for` pour parcourir chaque ligne des deux Sudokus. Si une ligne n'est pas pareille dans les deux Sudokus, elle renvoie `false`. Sinon, elle renvoie `true`.

## Fonction `isValidBoard`

Cette fonction vérifie si un Sudoku est bon. Elle utilise deux boucles `for` pour parcourir chaque case du Sudoku. Pour chaque case, elle vérifie si le chiffre dans la case est bon avec la fonction `isValid`. Si le chiffre n'est pas bon, elle renvoie `false`. Sinon, elle continue avec la case suivante. Si toutes les cases sont bonnes, elle renvoie `true`.

## Fonction `isValid`

Cette fonction vérifie si un chiffre est bon à une position donnée dans le Sudoku. Elle vérifie si le chiffre est unique dans sa ligne, sa colonne et son bloc 3x3. Pour cela, elle utilise une boucle `for` pour parcourir chaque case de la ligne, de la colonne et du bloc 3x3. Si le chiffre est déjà présent, elle renvoie `false`. Sinon, elle renvoie `true`.

## Fonction `solveSudoku`

Cette fonction essaie de résoudre le Sudoku. Elle utilise une boucle `for` pour essayer chaque chiffre dans chaque case vide du Sudoku. Pour chaque chiffre, elle vérifie si le chiffre est bon avec la fonction `isValid`. Si le chiffre est bon, elle met le chiffre dans la case et continue avec la case suivante. Si elle ne peut pas mettre un chiffre bon dans une case, elle remet la case à vide et essaie le chiffre suivant. Si elle peut mettre un chiffre bon dans toutes les cases, elle renvoie `true`. Sinon, elle renvoie `false`.
