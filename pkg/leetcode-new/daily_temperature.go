package main

// O(nˆ2)
func dailyTemperatures(temperatures []int) []int {
	answers := make([]int, len(temperatures))

	i := 0
	answers[len(temperatures)-1] = 0
	for i < len(temperatures)-1 {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				answers[i] = j - i
				break
			}
		}
		i++
	}

	return answers
}
