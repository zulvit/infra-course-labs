package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	password string `json:"pass"`
}

func main() {
	u := User{Name: "Alice", password: "123"}
	data, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
