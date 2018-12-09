package main

import (
	"advent/utils"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
)

type elf struct {
	left, top, width, height int
}

type entry struct {
	start int
	delta int
}

func main() {
	lines := utils.ReadAllLines("input.txt")

	re := regexp.MustCompile(".* @ (.*),(.*): (.*)x(.*)")
	var elves []elf
	rowCount := 0
	for _, line := range lines {
		m := re.FindStringSubmatch(line)
		left, _ := strconv.ParseInt(m[1], 10, 64)
		top, _ := strconv.ParseInt(m[2], 10, 64)
		width, _ := strconv.ParseInt(m[3], 10, 64)
		height, _ := strconv.ParseInt(m[4], 10, 64)
		elves = append(elves, elf{int(left), int(top), int(width), int(height)})
		if int(top+height) > rowCount {
			rowCount = int(top + height)
		}
	}

	rows := make([][]entry, rowCount)
	for _, elf := range elves {
		for y := 0; y < elf.height; y++ {
			rows[elf.top+y] = append(rows[elf.top+y],
				entry{start: elf.left, delta: 1},
				entry{start: elf.left + elf.width, delta: -1})
		}
	}

	overlap := 0
	for y := 0; y < rowCount; y++ {
		row := rows[y]
		sort.Slice(row, func(i, j int) bool { return row[i].start < row[j].start })

		layers := 0
		lastIndex := 0
		for _, e := range row {
			if layers > 1 {
				overlap += e.start - lastIndex
			}
			layers += e.delta
			lastIndex = e.start
		}

		if layers < 0 {
			log.Fatal(layers)
		}
	}

	fmt.Println(overlap)
}
