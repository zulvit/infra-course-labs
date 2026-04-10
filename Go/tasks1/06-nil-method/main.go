package main

import "fmt"

type Node struct{ Value int }

func (n *Node) Print() {
	if n == nil {
		fmt.Println("I am nil!")
		return
	}
	fmt.Println(n.Value)
}

func main() {
	var n *Node
	n.Print()
}
