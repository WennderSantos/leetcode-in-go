package availablecapturesforrook

func rPosition(board [][]byte) []int {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'R' {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func captures(board [][]byte, rposition []int, x int, y int) int {
	i, j := rposition[0], rposition[1]

	for i < len(board) && i >= 0 && j < len(board[0]) && j >= 0 {
		if board[i][j] == 'p' {
			return 1
		}

		if board[i][j] == 'B' {
			return 0
		}

		i, j = i+x, j+y
	}

	return 0
}

//NumRookCaptures executes in
//Time: O(n*m) worst case will look into the entire matrix to find 'R'
//Memory: O(1)
func NumRookCaptures(board [][]byte) int {
	rposition := rPosition(board)

	left := captures(board, rposition, 0, -1)
	right := captures(board, rposition, 0, +1)
	top := captures(board, rposition, -1, 0)
	bottom := captures(board, rposition, +1, 0)

	return left + right + top + bottom
}

//999. Available Captures for Rook
//https://leetcode.com/problems/available-captures-for-rook/
