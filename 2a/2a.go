package main

import (
	"advent/utils"
	"fmt"
)

func main() {
	lines := utils.ReadAllLines("input.txt")

	const letterCount = int('z') - int('a') + 1

	twos, threes := 0, 0
	for _, line := range lines {
		var l [letterCount]int
		for _, c := range line {
			l[int(c)-int('a')] ++
		}
		for _, v := range l {
			if v == 2 {
				twos ++
				break
			}
		}
		for _, v := range l {
			if v == 3 {
				threes ++
				break
			}
		}
	}

	fmt.Println(twos * threes)
}
