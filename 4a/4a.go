package main

import (
	"advent/utils"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type EventType int

const (
	ShiftBegins EventType = iota
	FallsAsleep
	WakesUp
)

type entry struct {
	line      string
	time      int
	eventType EventType
	minute    int
	guardId   int
}

func parseInput(lines []string) []entry {
	var entries []entry

	timeRegexp := regexp.MustCompile("\\[(\\d{4})-(\\d{2})-(\\d{2}) (\\d{2}):(\\d{2})")
	beginsShiftRegexp := regexp.MustCompile("Guard #(\\d*) begins shift")
	for _, line := range lines {
		timeMatch := timeRegexp.FindStringSubmatch(line)
		year, _ := strconv.ParseInt(timeMatch[1], 10, 64)
		month, _ := strconv.ParseInt(timeMatch[2], 10, 64)
		day, _ := strconv.ParseInt(timeMatch[3], 10, 64)
		hour, _ := strconv.ParseInt(timeMatch[4], 10, 64)
		minute, _ := strconv.ParseInt(timeMatch[5], 10, 64)
		time := minute + 60*(hour+24*(day+31*(month+366*(year))))

		_ = time

		var eventType EventType
		guardId := int64(-1)
		if match := beginsShiftRegexp.FindStringSubmatch(line); len(match) > 0 {
			eventType = ShiftBegins
			guardId, _ = strconv.ParseInt(match[1], 10, 64)
		} else if strings.Contains(line, "falls asleep") {
			eventType = FallsAsleep
		} else if strings.Contains(line, "wakes up") {
			eventType = WakesUp
		}
		entries = append(entries, entry{line, int(time), eventType, int(minute), int(guardId)})
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].time < entries[j].time })
	return entries
}

func main() {
	lines := utils.ReadAllLines("input.txt")
	entries := parseInput(lines)

	asleep := make(map[int][]int)
	asleepSum := make(map[int]int)

	currentGuard := -1
	fellAsleep := 0
	for _, entry := range entries {
		switch entry.eventType {
		case ShiftBegins:
			currentGuard = entry.guardId
		case FallsAsleep:
			fellAsleep = entry.minute
			if fellAsleep < 0 || fellAsleep > 59 {
				panic(entry)
			}
		case WakesUp:
			if fellAsleep >= entry.minute {
				panic(entry)
			}
			minutes, hasKey := asleep[currentGuard]
			if !hasKey {
				minutes = make([]int, 60)
				asleep[currentGuard] = minutes
			}
			for i := fellAsleep; i < entry.minute; i++ {
				minutes[i]++
			}
			asleepSum[currentGuard] = asleepSum[currentGuard] + entry.minute - fellAsleep
		}
	}

	maxAsleep := 0
	maxGuard := -1

	for guard, minutes := range asleepSum {
		if minutes > maxAsleep {
			maxGuard = guard
			maxAsleep = minutes
		}
	}

	maxMinuteIndex := -1
	maxMinuteCount := 0
	x := asleep[maxGuard]
	for i := 0; i < 60; i++ {
		if maxMinuteCount < x[i] {
			maxMinuteCount = x[i]
			maxMinuteIndex = i
		}
	}

	fmt.Println(maxMinuteIndex * maxGuard)
}
