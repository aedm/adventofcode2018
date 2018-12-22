package main

import (
	"advent/utils"
	"fmt"
	"regexp"
	"sort"
)

type Step struct {
	name string
	ins  map[*Step]bool
	outs map[*Step]bool
}

func main() {
	lines := utils.ReadAllLines("input.txt")

	steps := make(map[string]*Step)
	inCounts := make(map[*Step]int)

	reg := regexp.MustCompile("Step (\\w*) must be finished before step (\\w*) can begin.")
	for _, line := range lines {
		match := reg.FindStringSubmatch(line)
		name1 := match[1]
		name2 := match[2]
		node1, ok := steps[name1]
		if !ok {
			node1 = &Step{name1, make(map[*Step]bool), make(map[*Step]bool)}
			steps[name1] = node1
		}
		node2, ok := steps[name2]
		if !ok {
			node2 = &Step{name2, make(map[*Step]bool), make(map[*Step]bool)}
			steps[name2] = node2
		}
		node2.ins[node1] = true
		node1.outs[node2] = true
	}

	type Worker struct {
		end  int
		step *Step
	}

	var readyToStart []*Step

	for _, node := range steps {
		inCounts[node] = len(node.ins)
		if len(node.ins) == 0 {
			readyToStart = append(readyToStart, node)
		}
	}

	var workers []*Worker
	for i := 0; i < 5; i++ {
		workers = append(workers, &Worker{0, nil})
	}

	unstarted := len(steps)

	for time := 0; unstarted > 0; time++ {
		for _, worker := range workers {
			if (worker.step != nil) && (worker.end <= time) {
				for node := range worker.step.outs {
					inCounts[node]--
					if inCounts[node] == 0 {
						readyToStart = append(readyToStart, node)
					}
				}
				worker.step = nil
			}
		}

		sort.Slice(readyToStart, func(i, j int) bool { return readyToStart[i].name < readyToStart[j].name })

		for _, worker := range workers {
			if worker.step == nil && len(readyToStart) > 0 {
				worker.step = readyToStart[0]
				worker.end = time + int(readyToStart[0].name[0]-'A') + 61
				readyToStart = readyToStart[1:]
				unstarted--
			}
		}
	}

	lastTime := 0
	for _, worker := range workers {
		if worker.step != nil && worker.end > lastTime {
			lastTime = worker.end
		}
	}
	fmt.Println(lastTime)
}
