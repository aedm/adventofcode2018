package main

import (
	"advent/utils"
	"fmt"
	"regexp"
	"strconv"
)

type elf struct {
	left, top, width, height int
}

func main() {
	lines := utils.ReadAllLines("input.txt")

	re := regexp.MustCompile(".* @ (.*),(.*): (.*)x(.*)")
	var elves []elf
	for _, line := range lines {
		m := re.FindStringSubmatch(line)
		left, _ := strconv.ParseInt(m[1], 10, 64)
		top, _ := strconv.ParseInt(m[2], 10, 64)
		width, _ := strconv.ParseInt(m[3], 10, 64)
		height, _ := strconv.ParseInt(m[4], 10, 64)
		elves = append(elves, elf{int(left), int(top), int(width), int(height)})
	}

loop:
	for i, e1 := range elves {
		for o, e2 := range elves {
			if !(i == o || e1.left >= e2.left+e2.width || e2.left >= e1.left+e1.width ||
				e1.top >= e2.top+e2.height || e2.top >= e1.top+e1.height) {
				continue loop
			}
		}
		fmt.Println(i + 1)
	}
}
