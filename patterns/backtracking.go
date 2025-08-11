package patterns

import (
	"fmt"
)

// Backtracking Pattern
// Bài toán: Generate all permutations

func DemoBacktracking() {
	fmt.Println("\n🔄 7. BACKTRACKING PATTERN")
	fmt.Println("Bài toán: Generate All Permutations")

	nums := []int{1, 2, 3}
	fmt.Printf("Input: %v\n", nums)

	// Generate permutations
	permutations := permute(nums)
	fmt.Printf("✅ All permutations: %v\n", permutations)

	// Demo thêm: N-Queens
	n := 4
	solutions := solveNQueens(n)
	fmt.Printf("\n👑 %d-Queens solutions count: %d\n", n, len(solutions))
	if len(solutions) > 0 {
		fmt.Printf("First solution:\n")
		printBoard(solutions[0])
	}

	fmt.Println("\n🏢 Ứng dụng thực tế:")
	fmt.Println("- Sudoku solver")
	fmt.Println("- Route planning & pathfinding")
	fmt.Println("- Configuration space exploration")
	fmt.Println("- Game AI (chess, checkers)")
}

func permute(nums []int) [][]int {
	var result [][]int
	var current []int
	used := make([]bool, len(nums))

	backtrackPermute(nums, current, used, &result)
	return result
}

func backtrackPermute(nums []int, current []int, used []bool, result *[][]int) {
	if len(current) == len(nums) {
		// Tạo copy để tránh reference issues
		perm := make([]int, len(current))
		copy(perm, current)
		*result = append(*result, perm)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		// Choose
		current = append(current, nums[i])
		used[i] = true

		// Explore
		backtrackPermute(nums, current, used, result)

		// Unchoose (backtrack)
		current = current[:len(current)-1]
		used[i] = false
	}
}

// Ví dụ thực tế: N-Queens
func solveNQueens(n int) [][]string {
	var result [][]string
	board := make([][]rune, n)
	for i := range board {
		board[i] = make([]rune, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	backtrackQueens(board, 0, &result)
	return result
}

func backtrackQueens(board [][]rune, row int, result *[][]string) {
	if row == len(board) {
		solution := make([]string, len(board))
		for i, r := range board {
			solution[i] = string(r)
		}
		*result = append(*result, solution)
		return
	}

	for col := 0; col < len(board); col++ {
		if isValidQueen(board, row, col) {
			board[row][col] = 'Q'
			backtrackQueens(board, row+1, result)
			board[row][col] = '.'
		}
	}
}

func isValidQueen(board [][]rune, row, col int) bool {
	n := len(board)

	// Check column
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// Check diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// Check anti-diagonal
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}

func printBoard(board []string) {
	for _, row := range board {
		fmt.Println(row)
	}
}
