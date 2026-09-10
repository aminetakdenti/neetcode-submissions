func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := map[byte]int{}
	tMap := map[byte]int{}

	for i, _ := range s {
		if _, sok := sMap[s[i]]; sok {
			sMap[s[i]] += 1
		} else {
			sMap[s[i]] = 1
		}

		if _, tok := tMap[t[i]]; tok {
			tMap[t[i]] += 1
		} else {
			tMap[t[i]] = 1
		}
	}

	for key := range sMap {
		if sMap[key] != tMap[key] {
			return false
		}
	}

	return true
}
