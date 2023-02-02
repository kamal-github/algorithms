package main

/*
[[1,0,7],
[2,0,6],
[3,4,5],
[0,3,0],
[9,0,20]]
*/
func getMaximumGold(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	max := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			visited := make([][]bool, rows)
			for k := 0; k < len(visited); k++ {
				row := make([]bool, cols)
				visited[k] = append(visited[k], row...)
			}

			maxGoldCollectedAt := calPath(grid, visited, i, j, grid[i][j], grid[i][j], grid[i][j], grid[i][j])
			if maxGoldCollectedAt > max {
				max = maxGoldCollectedAt
			}
		}
	}
	return max
}

func calPath(grid [][]int, visited [][]bool, i, j, goldLeft, goldRight, goldTop, goldBottom int) int {
	visited[i][j] = true

	// left
	if j-1 >= 0 && (grid[i][j-1] > 0 && !visited[i][j-1]) {
		calPath(grid, visited, i, j-1, goldLeft, goldRight, goldTop, goldBottom)
		goldLeft = goldLeft + grid[i][j-1]
	}

	// right
	if j+1 < len(grid[0]) && (grid[i][j+1] > 0 && !visited[i][j+1]) {
		calPath(grid, visited, i, j+1, goldLeft, goldRight, goldTop, goldBottom)
		goldRight = goldRight + grid[i][j+1]
	}

	// top
	if i-1 >= 0 && (grid[i-1][j] > 0 && !visited[i-1][j]) {
		calPath(grid, visited, i-1, j, goldLeft, goldRight, goldTop, goldBottom)
		goldTop = goldTop + grid[i-1][j]
	}

	// bottom
	if i+1 < len(grid) && (grid[i+1][j] > 0 && !visited[i+1][j]) {
		calPath(grid, visited, i+1, j, goldLeft, goldRight, goldTop, goldBottom)
		goldBottom = goldBottom + grid[i+1][j]
	}

	return maxInArray([]int{goldTop, goldBottom, goldLeft, goldRight})
}

func maxInArray(arr []int) int {
	max := arr[0]

	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}

	return max
}
