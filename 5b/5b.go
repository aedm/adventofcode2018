package main

import (
	"advent/utils"
	"fmt"
	"strings"
	"unicode"
)

func collapsedLength(line string) int {
	var s []rune
	for _, c := range line {
		if len(s) > 0 {
			c2 := s[len(s)-1]
			if (unicode.IsUpper(c) != unicode.IsUpper(c2)) && (unicode.ToUpper(c) == unicode.ToUpper(c2)) {
				s = s[:len(s)-1]
				continue
			}
		}
		s = append(s, c)
	}
	return len(s)
}

func main() {
	line := utils.ReadAllLines("input.txt")[0]

	min := len(line)
	for c := 'a'; c <= 'z'; c++ {
		f1 := strings.Replace(line, string(c), "", -1)
		f2 := strings.Replace(f1, string(unicode.ToUpper(c)), "", -1)
		l := collapsedLength(f2)
		if l < min {
			min = l
		}
	}
	fmt.Println(min)
}
