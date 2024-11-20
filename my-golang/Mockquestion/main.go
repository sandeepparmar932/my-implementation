package main

import "fmt"

// Marker interface
type Identifiable interface{}

type User struct{}
type Admin struct{}

func main() {
	var u User
	var a Admin

	fmt.Printf("User implements Identifiable: %v\n", isIdentifiable(u))
	fmt.Printf("Admin implements Identifiable: %v\n", isIdentifiable(a))
}

func isIdentifiable(i Identifiable) bool {
	_, ok := i.(Identifiable)
	return ok
}
