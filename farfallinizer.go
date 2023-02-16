package farfallinizer

import "strings"

func Farfallinize(s string) (f string) {
	for charidx := 0; charidx < len(s); charidx++ {
		if strings.Contains("aeiou", s[charidx]) {
			println(1)
		} else {
			println(2)
		}
	}
}
