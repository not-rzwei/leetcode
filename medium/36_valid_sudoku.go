/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */
package medium

// @lc code=start
type Set []byte

func (s *Set) Add(x byte) {
	*s = append(*s, x)
}

func (s *Set) Has(x byte) bool {
	for _, b := range *s {
		if b == x {
			return true
		}
	}

	return false
}

func isValidSudoku(board [][]byte) bool {
	cols := [9]Set{}     // number container for each column
	rows := [9]Set{}     // number container for each row
	grids := [3][3]Set{} // number container for each grid

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			cell := board[row][col]

			// skip if cell is empty
			if cell == '.' {
				continue
			}

			// If cell is somewhere in cols, rows or grids
			// return false because it's a duplicate
			if rows[row].Has(cell) ||
				cols[col].Has(cell) ||
				grids[row/3][col/3].Has(cell) {
				return false
			}

			// Add cell to rows, cols and grids
			// Will happen when there is no dupe
			rows[row].Add(cell)
			cols[col].Add(cell)
			grids[row/3][col/3].Add(cell)
		}
	}

	return true
}

// @lc code=end
