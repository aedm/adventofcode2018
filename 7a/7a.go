package main

import (
	"advent/utils"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type node struct {
	name string
	ins  map[*node]bool
	outs map[*node]bool
}

func main() {
	lines := utils.ReadAllLines("input.txt")

	nodes := make(map[string]*node)
	var sources []*node

	reg := regexp.MustCompile("Step (\\w*) must be finished before step (\\w*) can begin.")
	for _, line := range lines {
		match := reg.FindStringSubmatch(line)
		name1 := match[1]
		name2 := match[2]
		node1, ok := nodes[name1]
		if !ok {
			node1 = &node{name1, make(map[*node]bool), make(map[*node]bool)}
			nodes[name1] = node1
		}
		node2, ok := nodes[name2]
		if !ok {
			node2 = &node{name2, make(map[*node]bool), make(map[*node]bool)}
			nodes[name2] = node2
		}
		node2.ins[node1] = true
		node1.outs[node2] = true
	}

	for _, node := range nodes {
		if len(node.ins) == 0 {
			sources = append(sources, node)
		}
	}

	var result []string
	for len(sources) > 0 {
		sort.Slice(sources, func(i, j int) bool { return sources[i].name < sources[j].name })
		for node, _ := range sources[0].outs {
			delete(node.ins, sources[0])
			if len(node.ins) == 0 {
				sources = append(sources, node)
			}
		}
		result = append(result, sources[0].name)
		sources = sources[1:]
	}

	fmt.Println(strings.Join(result,""))
}
