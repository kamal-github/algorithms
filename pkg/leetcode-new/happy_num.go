package main

func isHappy(n int) bool {
	localSum := n
	// todo what if is 0
	for localSum != 1 {
		localSum = calNewLocalSum(localSum)
		if localSum == 0 || localSum < 10 && localSum > 1 {
			return false
		}
	}

	return true
}

func calNewLocalSum(n int) int {
	localSum := 0

	for n != 0 {
		lastDigit := n % 10
		n = int(n / 10)

		localSum = (lastDigit * lastDigit) + localSum
	}

	return localSum
}
