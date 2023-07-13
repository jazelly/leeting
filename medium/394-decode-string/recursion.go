package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func decodeString(s string) string {
	return decodeBlock(s, 1)
}

func decodeBlock(s string, times int) string {
	res := ""
	l := 0
	for ; l < len(s); l++ {
		// base case, just pass
		if unicode.IsLower(rune(s[l])) {
			res += string(s[l])
			continue
		}

		_, err := strconv.Atoi(string(s[l]))
		if err == nil {
			// find digits
			dr := l
			for s[dr] != byte('[') {
				dr++
			}
			num, _ := strconv.Atoi(string(s[l:dr]))

			// find block start and end
			innerBlock := 0
			j := dr + 1
			for ; j < len(s); j++ {
				if s[j] == byte('[') {
					innerBlock++
				} else if s[j] == byte(']') {
					if innerBlock == 0 {
						break
					}
					innerBlock--
				}
			}

			decoded := decodeBlock(s[dr+1:j], num)
			res += decoded
			l = j
		}
	}
	finalRes := ""
	for i := 0; i < times; i++ {
		finalRes += res
	}
	return finalRes
}

func main() {
	a := decodeString("100[leetcode]")
	fmt.Println(a)
}
