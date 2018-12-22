package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadAllLines("input.txt")

	minx, miny, maxx, maxy := 10000, 10000, 0, 0

	var xs []int
	var ys []int

	for _, line := range lines {
		nums := strings.Split(line, ", ")
		xc, _ := strconv.ParseInt(nums[0], 10, 64)
		yc, _ := strconv.ParseInt(nums[1], 10, 64)
		x := int(xc)
		y := int(yc)
		if minx > x {
			minx = x
		}
		if miny > y {
			miny = y
		}
		if maxx < x {
			maxx = x
		}
		if maxy < y {
			maxy = y
		}
		xs = append(xs, x)
		ys = append(ys, y)
	}

	w := maxx - minx + 1
	h := maxy - miny + 1

	onedge := make([]bool, len(xs))
	count := make([]int, len(xs))

	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			n := -1
			mind := w + h
			for i := 0; i < len(xs); i++ {
				dx := xs[i] - x
				dy := ys[i] - y
				if dx < 0 {
					dx = -dx
				}
				if dy < 0 {
					dy = -dy
				}
				d := dx + dy
				if d == mind {
					n = -1
				} else if d < mind {
					n = i
					mind = d
				}
			}
			if n >= 0 {
				if x == minx || x == maxx || y == miny || y == maxy {
					onedge[n] = true
				}
				count[n]++
			}
		}
	}

	maxc := 0
	for i, _ := range xs {
		if count[i] > maxc && !onedge[i] {
			maxc = count[i]
		}
	}

	fmt.Println(maxc)
}
