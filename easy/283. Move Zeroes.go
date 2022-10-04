// 1 Solution
func moveZeroes(nums []int) {
	zeroes := []int{}
	numbers := []int{}

	for _, k := range nums {
		if k == 0 {
			zeroes = append(zeroes, 0)
		} else {
			numbers = append(numbers, k)
		}
	}

	for i, k := range append(numbers, zeroes...) {
		nums[i] = k
	}
}

// 2 Solution
func moveZeroes(nums []int) {
	lastZero := 0

	for _, k := range nums {
		if k != 0 {
			nums[lastZero] = k
			lastZero++
		}
	}

	for ; lastZero < len(nums); lastZero++ {
		nums[lastZero] = 0
	}
}