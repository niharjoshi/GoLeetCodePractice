func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	mapS := make(map[rune]int)

	for _, char := range s {
		mapS[char]++
	}

	for _, char := range t {

		if mapS[char] == 0 {
			return false
		}

		mapS[char]--

	}

	return true
}