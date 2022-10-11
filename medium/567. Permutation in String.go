func checkInclusion(s1 string, s2 string) bool {
	subLen := len(s1)
	mainLen := len(s2)

	if subLen > mainLen {
		return false
	}

	targetArray := [256]uint16{}
	currentArray := [256]uint16{}

	for i := 0; i < subLen; i++ {
		targetArray[int(s1[i])]++
		currentArray[int(s2[i])]++
	}

	if targetArray == currentArray {
		return true
	}

	for right := subLen; right < mainLen; right++ {
		currentArray[s2[right-subLen]]--
		currentArray[s2[right]]++
		if currentArray == targetArray {
			return true
		}
	}

	return false
}
