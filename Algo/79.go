package algo

const start int = 0
const fromUp int = 1
const fromDown int = 2
const fromLeft int = 3
const fromRight int = 4

func nextChar(board [][]byte, i, j int, word string, from int) bool {
	if word == "" {
		return true
	}
	board[i][j] += 100
	if i > 0 && board[i-1][j] == word[0] && from != fromUp && nextChar(board, i-1, j, word[1:], fromDown) {
		return true
	}
	if i < len(board)-1 && board[i+1][j] == word[0] && from != fromDown && nextChar(board, i+1, j, word[1:], fromUp) {
		return true
	}
	if j > 0 && board[i][j-1] == word[0] && from != fromLeft && nextChar(board, i, j-1, word[1:], fromRight) {
		return true
	}
	if j < len(board[i])-1 && board[i][j+1] == word[0] && from != fromRight && nextChar(board, i, j+1, word[1:], fromLeft) {
		return true
	}
	board[i][j] -= 100
	return false
}

func Exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				if nextChar(board, i, j, word[1:], start) {
					return true
				}
			}
		}
	}
	return false
}
