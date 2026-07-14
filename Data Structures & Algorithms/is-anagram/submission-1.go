func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freqS := make(map[rune]int)
	
	for i := 0; i < len(s); i++ {
		freqS[rune(s[i])]++
	}
	
	freqT := make(map[rune]int)

	for i := 0; i < len(t); i++ {
		freqT[rune(t[i])]++
	}

	for key, val1 := range freqS {
        val2, exists := freqT[key]
        if !exists || val1 != val2 {
            return false
        }
    }

	return true
}
