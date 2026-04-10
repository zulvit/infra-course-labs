package main

import "fmt"

type User struct {
	Name string
}

func main() {
	var u1 User
	var u2 *User

	fmt.Println("u1 Name:", u1.Name)
	fmt.Println("u2 Name:", u2.Name)
}
