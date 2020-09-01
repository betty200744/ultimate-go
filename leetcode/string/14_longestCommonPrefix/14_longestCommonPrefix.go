package _4_longestCommonPrefix

func LongestCommonPrefix(strs []string) string {
	end := 0
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			if len(strs[j]) <= i {
				return strs[0][0:end]
			}
			if strs[j][i] == strs[0][i] {
				// the last strs item
				if j == len(strs)-1 {
					end = i + 1
				}
			} else {
				return strs[0][0:end]
			}
		}
	}
	return strs[0][0:end]
}
