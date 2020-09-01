package _6_minWindow

func IncludeAll(s string, t string) bool {
	start := 0
	mapS := ToMap(s)
	mapT := ToMap(t)
	for ; start < len(s); start++ {
		if mapS[string(s[start])] > mapT[string(s[start])] {
			mapS[string(s[start])]--
			mapT[string(s[start])]--
		} else {
			return false
		}
	}
	return true
}

func MinWindow2(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	for gap := len(t) - 1; gap < len(s); gap++ {
		for i := 0; i < len(s)-gap; i++ {
			if IncludeAll(s[i:i+gap+1], t) {
				return s[i : i+gap+1]
			}
		}
	}
	return ""
}

func ToMap(s string) map[string]int {
	mapS := make(map[string]int)
	for _, i2 := range s {
		if _, ok := mapS[string(i2)]; ok {
			mapS[string(i2)]++
		} else {
			mapS[string(i2)] = 1
		}
	}
	return mapS
}

func MinWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	if len(s) < len(t) {
		return ""
	}
	start := 0
	end := len(s) - 1
	mapS := ToMap(s)
	mapT := ToMap(t)
	for ; start < len(s); start++ {
		if mapS[string(s[start])] > mapT[string(s[start])] {
			mapS[string(s[start])]--
		} else {
			break
		}
	}
	for ; end > 0; end-- {
		if mapS[string(s[end])] > mapT[string(s[end])] {
			mapS[string(s[end])]--
		} else {
			break
		}
	}

	return s[start : end+1]
}
