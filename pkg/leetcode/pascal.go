package leetcode

func GetRow(k int) []int {
	a := make([]int, k+1)
	mem := initArray(k+1)
	for j:=k; j>=0; j-- {
		a[j] = pascalEff(mem, k, j)
	}
	return a
}


// Generate generates pascal triangle of size n
func Generate(n int) [][]int {
	a := initArray(n)
	for i:=n-1; i>=0; i-- {
		for j:=i; j>=0; j-- {
			a[i][j] = pascal( i, j)
		}
	}
	return a
}

// Simple solution
func pascal(i, j int ) int {
	if j == 0 || j == i {
		return 1
	}

	return pascal(i-1, j-1) + pascal(i-1, j)
}

// Generate generates pascal triangle of size n
// using Dynamic Programming
func GenerateEff(n int) [][]int {
	a := initArray(n)
	for i:=n-1; i>=0; i-- {
		for j:=i; j>=0; j-- {
			a[i][j] = pascalEff(a, i, j)
		}
	}
	return a
}

// Efficient solution using memoization (DP)
func pascalEff(mem [][]int, i, j int) int {
	if j == 0 || j == i {
		return 1
	}
	if mem[i][j] != 0 {
		return mem[i][j]
	}

	mem[i][j] = pascalEff(mem, i-1, j-1) + pascalEff(mem, i-1, j)
	return mem[i][j]
}

func initArray(n int) [][]int {
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, i+1)
	}
	return a
}

