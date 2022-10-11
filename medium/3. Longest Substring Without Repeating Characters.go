func lengthOfLongestSubstring(s string) int {
	sub := map[rune]int{}
	str := []rune(s)
	//sub := []rune{}
	left := 0
	right := 0
	length := 0

	for i := 0; i < len(str); i++ {
		if v, ok := sub[str[i]]; ok != true {
			sub[str[i]] = i
		} else {
			left = v + 1
			sub = sliceToMap(str[left:(right+1)], left)
		}

		if length < len(sub) {
			length = len(sub)
		}
		right++
	}

	return length
}

func sliceToMap(str []rune, offset int) map[rune]int {
	result := map[rune]int{}
	for index, value := range str {
		result[value] = index + offset
	}

	return result
}