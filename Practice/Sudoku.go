package main

import "fmt"

func main() {
	Puzzle := [][]int{
		[]int{0, 7, 0, 0, 2, 9, 0, 4, 5},
		[]int{0, 0, 0, 1, 5, 8, 0, 0, 6},
		[]int{2, 0, 0, 3, 0, 0, 0, 9, 0},
		[]int{0, 0, 0, 0, 0, 1, 0, 0, 0},
		[]int{9, 0, 0, 2, 4, 0, 6, 0, 7},
		[]int{4, 0, 0, 0, 6, 0, 9, 0, 3},
		[]int{0, 6, 0, 0, 0, 0, 0, 0, 8},
		[]int{0, 0, 0, 4, 3, 0, 0, 0, 0},
		[]int{0, 3, 0, 5, 0, 6, 0, 7, 0},
	}
	Display(Puzzle, 0, 0)
	if Solve(Puzzle) {
		fmt.Println("After Solving:")
		Display(Puzzle, 0, 0)
	} else {
		fmt.Println("Cannot Solve")
	}
}
func rc(Puzzle [][]int, r, c int) (int, int, bool) {
	if r > len(Puzzle)-1 {
		return 0, 0, true
	}
	if c < 9 {
		if Puzzle[r][c] == 0 {
			return r, c, false // if you found some empty element in row, then return
		}
		return rc(Puzzle, r, c+1)
	}
	return rc(Puzzle, r+1, 0)
}
func PlacingNum(row, col, number int, Puzzle [][]int) bool {
	if number > 9 {
		return false
	}
	if IsSafe(Puzzle, row, col, number) {
		Puzzle[row][col] = number
		if Solve(Puzzle) {
			// found the answer
			return true
		} else {
			// backtrack
			Puzzle[row][col] = 0
		}
	}
	return PlacingNum(row, col, number+1, Puzzle)
}
func Solve(Puzzle [][]int) bool {
	row := 0
	col := 0
	emptyleft := true
	// this is how we are replacing the r,c from arguments
	row, col, emptyleft = rc(Puzzle, 0, 0)

	if emptyleft == true {
		return true
		// soduko is solved
	}
	return PlacingNum(row, col, 1, Puzzle)
}
func RCheck(row, col, num int, Puzzle [][]int) bool {
	if col > len(Puzzle)-1 {
		return true
	}
	// check if the number is in the row
	if Puzzle[row][col] == num {
		return false
	}
	return RCheck(row, col+1, num, Puzzle)
}
func CCheck(row, col, num int, Puzzle [][]int) bool {
	if row > len(Puzzle)-1 {
		return true
	}
	// check if the number is in the col
	if Puzzle[row][col] == num {
		return false
	}
	return CCheck(row+1, col, num, Puzzle)
}
func BoxCheck(rowstart, colstart, r, c, num int, Puzzle [][]int) bool {
	if r > rowstart+2 {
		return true
	}
	if c < colstart+3 {
		if Puzzle[r][c] == num {
			return false
		}
		return BoxCheck(rowstart, colstart, r, c+1, num, Puzzle)
	}
	return BoxCheck(rowstart, colstart, r+1, colstart, num, Puzzle)
}
func IsSafe(Puzzle [][]int, row, col, num int) bool {
	// check the row
	if !RCheck(row, 0, num, Puzzle) {
		return false
	}
	// check the col
	if !CCheck(0, col, num, Puzzle) {
		return false
	}
	rowStart := row - row%3
	colStart := col - col%3
	return BoxCheck(rowStart, colStart, row, col, num, Puzzle)
}
func Display(Puzzle [][]int, r, c int) {
	if r > len(Puzzle)-1 {
		return
	}
	if c < len(Puzzle) {
		if r%3 == 0 && c == 0 {
			Lines("_____", 0)
			fmt.Println()
		}
		if r%3 != 0 && c == 0 {
			Lines("-----", 0)
			fmt.Println()
		}
		if c == 0 || c%3 == 0 {
			fmt.Print(" | ")
		}
		fmt.Print(Puzzle[r][c], " | ")
		if r == 8 && c == 8 {
			fmt.Println()
			Lines("_____", 0)
		}
		Display(Puzzle, r, c+1)
	} else {
		fmt.Println()
		Display(Puzzle, r+1, 0)
	}
}
func Lines(str string, r int) {
	if r > 8 {
		return
	}
	fmt.Print(str)
	Lines(str, r+1)
}
