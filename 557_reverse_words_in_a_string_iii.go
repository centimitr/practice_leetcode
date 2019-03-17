package practice_leetcode

import "strings"

func reverseWords(s string) string {
	var words []string
	for _, word := range strings.Split(s, " ") {
		b := []rune(word)
		l := len(b)
		for i := 0; i < l/2; i++ {
			b[i], b[l-i-1] = b[l-i-1], b[i]
		}
		words = append(words, string(b))
	}
	return strings.Join(words, " ")
}

//  or split('').reverse().join('') by char
//  then split(' ').reverse().join(' ') by word
