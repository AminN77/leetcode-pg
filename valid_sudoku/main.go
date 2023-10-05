package main

import "log"

func main() {
	input := [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	}
	log.Println(isValidSudoku(input))
}

func isValidSudoku(board [][]byte) bool {
	// validate rows
	for i := 0; i < 9; i++ {
		tMap := make(map[byte]int)
		for j := 0; j < 9; j++ {
			if tMap[board[i][j]] == 1 && board[i][j] != '.' {
				return false
			}
			tMap[board[i][j]]++
		}
	}

	// validate cols
	for i := 0; i < 9; i++ {
		tMap := make(map[byte]int)
		for j := 0; j < 9; j++ {
			if tMap[board[j][i]] == 1 && board[j][i] != '.' {
				return false
			}
			tMap[board[j][i]]++
		}
	}

	defRow := 0
	defCol := 0
	for defRow < 9 {
		for defCol < 9 {
			tMap := make(map[byte]int)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if tMap[board[i+defRow][j+defCol]] == 1 && board[i+defRow][j+defCol] != '.' {
						return false
					}
					tMap[board[i+defRow][j+defCol]]++
				}
			}
			defCol += 3
		}

		defRow += 3
		defCol = 0
	}

	return true
}
