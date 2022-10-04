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