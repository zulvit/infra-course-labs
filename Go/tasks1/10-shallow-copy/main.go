package main

import "fmt"

type Worker struct {
	ID   int
	Tags []string
}

func main() {
	w1 := Worker{ID: 1, Tags: []string{"go", "backend"}}
	w2 := w1

	w2.ID = 2
	w2.Tags[0] = "rust"

	fmt.Println("w1:", w1.ID, w1.Tags)
	fmt.Println("w2:", w2.ID, w2.Tags)
}
