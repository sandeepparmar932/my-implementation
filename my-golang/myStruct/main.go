package main

import "fmt"

func main() {

	sandeep := User{"Sandee", "abc@gmail", true, 24}
	fmt.Println(sandeep)
	aman := new(User)
	fmt.Println(aman)
	aman.Email = "abca"
	aman.Name = "aman"
	fmt.Println(aman)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
