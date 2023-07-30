package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/robryanx/adventofcode2019/modules/readinput"
)

type Node struct {
	name   string
	parent *Node
}

func main() {
	node_list, _ := build_tree()

	// get the nodes of interest
	you_node, _ := node_exists(node_list, "YOU")
	you_parents := get_parents_list(you_node)

	santa_node, _ := node_exists(node_list, "SAN")
	santa_parents := get_parents_list(santa_node)

	distance := 0
outer_loop:
	for i, you_parent := range you_parents {
		for j, santa_parent := range santa_parents {
			if you_parent == santa_parent {
				distance = i + j - 2

				break outer_loop
			}
		}
	}

	fmt.Println(distance)
}

func get_parents_list(current_node *Node) []*Node {
	var node_list []*Node

	for current_node.parent != nil {
		node_list = append(node_list, current_node)
		current_node = current_node.parent
	}

	return node_list
}

func build_tree() ([]*Node, *Node) {
	var node_list []*Node
	var root_node *Node

	for _, association := range readinput.ReadStrings("inputs/6/input.txt", "\n") {
		association_part := strings.Split(association, ")")

		current_node, err := node_exists(node_list, association_part[0])
		if err != nil {
			current_node = &Node{name: association_part[0]}
			if len(node_list) == 0 {
				root_node = current_node
			}

			node_list = append(node_list, current_node)
		}

		association_node, err := node_exists(node_list, association_part[1])
		if err != nil {
			association_node = &Node{name: association_part[1]}
			node_list = append(node_list, association_node)
		}

		association_node.parent = current_node
	}

	return node_list, root_node
}

func node_exists(node_list []*Node, name string) (*Node, error) {
	for i := range node_list {
		if node_list[i].name == name {
			return node_list[i], nil
		}
	}

	return &Node{}, errors.New("Not found")
}
