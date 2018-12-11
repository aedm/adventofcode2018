package main

import (
	"advent/utils"
	"fmt"
	"unicode"
)

func main() {
	line := utils.ReadAllLines("input.txt")[0]

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

	fmt.Println(string(s))
	fmt.Println(len(s))
}
