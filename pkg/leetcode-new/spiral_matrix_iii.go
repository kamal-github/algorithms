package main

/*
for each iteration until numOfVisitedElements == rows*cols
E,S,W,N
1,1,2,2
3,3,4,4
5,5,6,6
*/

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	visitedElements := make([][]int, 0)

	visitedElements = append(visitedElements, []int{rStart, cStart}) // start
	numOfVisitedElements := 1
	lastStepsCount := 1
	currRow, currCol := rStart, cStart
	for {
		if numOfVisitedElements >= rows*cols {
			break
		}

		// i=0 e -> s -> w -> w -> n -> n
		// i=1 e -> e -> e -> s -> s -> s ->s -> w -> w -> w -> w -> n -> n -> n -> n
		for eastSteps := 0; eastSteps < lastStepsCount; eastSteps++ {
			r, c := goEast(currRow, currCol)
			if r >= 0 && r < rows && c >= 0 && c < cols {
				visitedElements = append(visitedElements, []int{r, c})
				numOfVisitedElements++
			}
			currRow = r
			currCol = c
		}
		for southSteps := 0; southSteps < lastStepsCount; southSteps++ {
			r, c := goSouth(currRow, currCol)
			if r >= 0 && r < rows && c >= 0 && c < cols {
				visitedElements = append(visitedElements, []int{r, c})
				numOfVisitedElements++
			}
			currRow = r
			currCol = c
		}
		for westSteps := 0; westSteps < lastStepsCount+1; westSteps++ {
			r, c := goWest(currRow, currCol)
			if r >= 0 && r < rows && c >= 0 && c < cols {
				visitedElements = append(visitedElements, []int{r, c})
				numOfVisitedElements++
			}
			currRow = r
			currCol = c
		}
		for northSteps := 0; northSteps < lastStepsCount+1; northSteps++ {
			r, c := goNorth(currRow, currCol)
			if r >= 0 && r < rows && c >= 0 && c < cols {
				visitedElements = append(visitedElements, []int{r, c})
				numOfVisitedElements++
			}
			currRow = r
			currCol = c
		}

		lastStepsCount = (lastStepsCount + 1) + 1
	}

	return visitedElements
}

func goNorth(i, j int) (int, int) {
	return i - 1, j
}

func goSouth(i, j int) (int, int) {
	return i + 1, j
}

func goEast(i, j int) (int, int) {
	return i, j + 1
}

func goWest(i, j int) (int, int) {
	return i, j - 1
}
