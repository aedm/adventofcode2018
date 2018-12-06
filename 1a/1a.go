package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

func main() {
	lines := utils.ReadAllLines("input.txt")
	var sum int64 = 0
	for _, line := range lines {
		num, _ := strconv.ParseInt(line[1:], 0, 64)
		if line[0] == '-' {
			num = -num
		}
		sum = sum + num
	}
	fmt.Println(sum)
}
