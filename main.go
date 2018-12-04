package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readAllLines(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func main1a() {
	lines := readAllLines("input.txt")
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

func main() {
	lines := readAllLines("input.txt")
	var sum int64 = 0
	set := make(map[int64]bool)
	loop: for {
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
