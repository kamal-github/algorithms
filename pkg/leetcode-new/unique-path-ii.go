package main

import "fmt"

var cacheKey = "%d-%d"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var paths int
	cache := make(map[string]int, 0)
	traverse(obstacleGrid, 0, 0, &paths, cache)
	return paths
}

func traverse(o [][]int, i, j int, paths *int, cache map[string]int) {
	if o[i][j] == 1 {
		return
	}
	if i == len(o)-1 && j == len(o[0])-1 {
		*paths = (*paths) + 1
		return
	}

	k := fmt.Sprintf(cacheKey, i, j)
	if p, ok := cache[k]; ok {
		*paths = (*paths) + p
		return
	}

	if i+1 < len(o) {
		traverse(o, i+1, j, paths, cache)
	}

	if j+1 < len(o[0]) {
		traverse(o, i, j+1, paths, cache)
	}

	cache[k] = *paths
}
