package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

func main() {
	lines := utils.ReadAllLines("input.txt")
	var sum int64 = 0
	set := make(map[int64]bool)
loop:
	for {
		for _, line := range lines {
			set[sum] = true
			num, _ := strconv.ParseInt(line[1:], 0, 64)
			if line[0] == '-' {
				num = -num
			}
			sum = sum + num
			if set[sum] {
				break loop
			}
		}
	}

	fmt.Println(sum)
}
