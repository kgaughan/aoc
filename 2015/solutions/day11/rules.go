package day11

func hasAtLeastThreeAscending(runes []rune) bool {
	for i := 0; i < len(runes)-3; i++ {
		if runes[i] == runes[i+1]-1 && runes[i+1] == runes[i+2]-1 {
			return true
		}
	}
	return false
}

func hasNoConfusingCharacters(runes []rune) bool {
	for _, ch := range runes {
		if ch == 'i' || ch == 'l' || ch == 'o' {
			return false
		}
	}
	return true
}

func hasTwoPairs(runes []rune) bool {
	foundFirst := false
	var i int
	for i = 0; i < len(runes)-3; i++ {
		if runes[i] == runes[i+1] {
			foundFirst = true
			break
		}
	}
	if foundFirst {
		for i += 2; i < len(runes)-1; i++ {
			if runes[i] == runes[i+1] {
				return true
			}
		}

	}
	return false
}
