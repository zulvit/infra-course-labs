package main

import "fmt"

type User struct {
	Name   string
	Active bool
}

func main() {
	users := []User{{Name: "Bob", Active: false}}

	for _, u := range users {
		u.Active = true
	}

	fmt.Println(users[0].Active)
}
