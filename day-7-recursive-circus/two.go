package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/X11/adventofcode-2017/util"
)

type Node struct {
	Parent   *Node
	Children []*Node
	Name     string
	Weight   int
}

func (n *Node) SetParent(p *Node) {
	n.Parent = p
}

func (n *Node) SetChildren(children []*Node) {
	for _, child := range children {
		child.SetParent(n)
	}
	n.Children = children
}

func (n *Node) GetTotalWeight() int {
	weight := n.Weight
	for _, child := range n.Children {
		weight += child.GetTotalWeight()
	}

	return weight
}

func getData(s string) (string, int) {
	parts := strings.Split(s, " ")
	weightStr := strings.Trim(parts[1], "()")
	weight, err := strconv.Atoi(weightStr)
	if err != nil {
		panic(err)
	}
	return parts[0], weight
}

func main() {
	input := util.ReadInputFromFile("./input.txt")

	nodes := make(map[string]*Node)

	createOrFindNode := func(name string, weight int) *Node {
		node, ok := nodes[name]
		if !ok {
			node = &Node{
				Name: name,
			}
			nodes[name] = node
		}

		if weight != 0 {
			node.Weight = weight
		}
		return node
	}

	// Setup tree
	var lastNode *Node
	instructions := strings.Split(input, "\n")
	for _, instruction := range instructions {
		if instruction == "" {
			continue
		}

		parts := strings.Split(instruction, "->")
		name, weight := getData(parts[0])
		children := []*Node{}
		if len(parts) > 1 {
			childs := strings.Split(parts[1], ",")
			for _, child := range childs {
				children = append(children, createOrFindNode(strings.TrimSpace(child), 0))
			}
		}

		node := createOrFindNode(name, weight)
		if len(children) > 0 {
			node.SetChildren(children)
		}
		lastNode = node
	}

	// Find root node
	for {
		if lastNode.Parent != nil {
			lastNode = lastNode.Parent
		} else {
			break
		}
	}

	index := 0
	var findCorruption func(n *Node)
	findCorruption = func(n *Node) {
		if index < 1 {
			fmt.Printf("+%s %+v - %+v\n", strings.Repeat("++", index), n.Weight, n.GetTotalWeight())
		}
		index++
		for _, child := range n.Children {
			if index < 2 {
				fmt.Printf("%s %+v - %+v\n", strings.Repeat("++", index), child.Weight, child.GetTotalWeight())
			}
			findCorruption(child)
		}
		index--
	}

	getNode := func(node *Node, ints []int) *Node {
		for _, k := range ints {
			node = node.Children[k]
		}
		return node
	}

	findCorruption(getNode(lastNode, []int{2, 2, 2}))
}
