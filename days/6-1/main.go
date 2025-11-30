package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/robryanx/adventofcode2019/util/readinput"
)

type Node struct {
	name   string
	parent *Node
}

func main() {
	node_list := build_tree()

	total_count := 0

	for _, base_node := range node_list {
		path_count := 0
		node := base_node
		for node.parent != nil {
			path_count++
			node = node.parent
		}

		total_count += path_count
	}

	fmt.Println(total_count)
}

func build_tree() []*Node {
	var node_list []*Node

	for _, association := range readinput.ReadStrings("inputs/6/input.txt", "\n") {
		association_part := strings.Split(association, ")")

		current_node, err := node_exists(node_list, association_part[0])
		if err != nil {
			current_node = &Node{name: association_part[0]}
			node_list = append(node_list, current_node)
		}

		association_node, err := node_exists(node_list, association_part[1])
		if err != nil {
			association_node = &Node{name: association_part[1]}
			node_list = append(node_list, association_node)
		}

		association_node.parent = current_node
	}

	return node_list
}

func node_exists(node_list []*Node, name string) (*Node, error) {
	for i := range node_list {
		if node_list[i].name == name {
			return node_list[i], nil
		}
	}

	return &Node{}, errors.New("Not found")
}
