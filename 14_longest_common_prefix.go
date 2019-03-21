package practice_leetcode

func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0 {
		return ""
	} else if l == 1 {
		return strs[0]
	}
	pos := 0
	for {
		if pos >= len(strs[0]) {
			goto end
		}
		for i := 1; i < len(strs); i++ {
			if pos >= len(strs[i]) {
				goto end
			}
			if strs[i][pos] != strs[i-1][pos] {
				goto end
			}
		}
		pos++
	}
end:
	return string(strs[0][:pos])
}
