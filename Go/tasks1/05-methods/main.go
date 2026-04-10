package main

import "fmt"

type User struct{ Name string }

func (u User) SetNameVal(n string)  { u.Name = n }
func (u *User) SetNamePtr(n string) { u.Name = n }

func main() {
	usr := User{Name: "Bob"}

	usr.SetNameVal("Alice")
	fmt.Println("После SetNameVal:", usr.Name)

	usr.SetNamePtr("Eve")
	fmt.Println("После SetNamePtr:", usr.Name)
}
