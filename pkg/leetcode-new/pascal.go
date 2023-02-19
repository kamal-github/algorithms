package main

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}

	pascal := [][]int{{1}, {1, 1}}

	for i := 2; i < numRows; i++ {
		prevRow := pascal[i-1]
		curRow := make([]int, i+1)
		curRow[0], curRow[len(curRow)-1] = 1, 1
		for j := 1; j < len(curRow)-1; j++ {
			curRow[j] = prevRow[j-1] + prevRow[j]
		}
		pascal = append(pascal, curRow)
	}

	return pascal
}
