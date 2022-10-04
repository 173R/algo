func twoSum(numbers []int, target int) []int {
	first := 0
	sum := 0
	last := len(numbers) - 1

	for first < last {
		sum = numbers[first] + numbers[last]

		if sum > target {
			last--
		} else if sum < target {
			first++
		} else {
			break
		}
	}

	return []int{first + 1, last + 1}
}