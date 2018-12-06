package main

import (
	"advent/utils"
	"fmt"
)

func main() {
	lines := utils.ReadAllLines("input.txt")

	for i, a := range lines[:len(lines)-1] {
		for _, b := range lines[i+1:] {
			if len(a) != len(b) {
				break
			}
			mismatch := 0
			for o:=0; o<len(a); o++ {
				if a[o] != b[o] {
					mismatch ++
				}
				if mismatch > 1 {
					break
				}
			}
			if mismatch == 1 {
				var chars []uint8
				for o:=0; o<len(a); o++ {
					if a[o] == b[o] {
						chars = append(chars, a[o])
					}
				}
				fmt.Println(string(chars))
				return
			}
		}
	}
}
