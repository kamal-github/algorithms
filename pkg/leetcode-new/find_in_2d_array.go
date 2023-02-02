package main

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}

	cols := len(matrix[0])
	return binarySearch(matrix, target, 0, 0, rows, cols)
}

func binarySearch(matrix [][]int, target, startRow, startCol, endRow, endCol int) bool {
	midRow := (startRow + endRow - 1) / 2
	midCol := (startCol + endCol - 1) / 2

	if midRow < startRow || midCol < startCol || midRow > endRow || midCol > endCol {
		return false
	}
	midElement := matrix[midRow][midCol]
	if midElement == target {
		return true
	} else if midElement < target {
		// right half
		return binarySearch(matrix, target, midRow, midCol+1, endRow, endCol)
	}

	// left half
	return binarySearch(matrix, target, startRow, startCol, midRow, midCol)

}
