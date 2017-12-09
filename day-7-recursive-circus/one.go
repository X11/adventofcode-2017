package main

import (
	"fmt"
	"github.com/X11/adventofcode-2017/util"
	"strings"
)

type Node struct {
	Parent   *Node
	Children []*Node
	Name     string
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

func getData(s string) string {
	parts := strings.Split(s, " ")
	return parts[0]
}

func main() {
	input := util.ReadInputFromFile("./input.txt")

	nodes := make(map[string]*Node)

	createOrFindNode := func(name string) *Node {
		node, ok := nodes[name]
		if !ok {
			node = &Node{
				Name: name,
			}
			nodes[name] = node
		}
		return node
	}

	var lastNode *Node
	instructions := strings.Split(input, "\n")
	for _, instruction := range instructions {
		if instruction == "" {
			continue
		}

		parts := strings.Split(instruction, "->")
		name := getData(parts[0])
		children := []*Node{}
		if len(parts) > 1 {
			childs := strings.Split(parts[1], ",")
			for _, child := range childs {
				children = append(children, createOrFindNode(strings.TrimSpace(child)))
			}
		}

		node := createOrFindNode(name)
		if len(children) > 0 {
			node.SetChildren(children)
		}
		lastNode = node
	}

	for {
		if lastNode.Parent != nil {
			fmt.Printf("lastNode.Parent = %+v\n", lastNode.Parent)
			lastNode = lastNode.Parent
		} else {
			fmt.Printf("lastNode.Name = %+v\n", lastNode.Name)
			break
		}
	}
}
