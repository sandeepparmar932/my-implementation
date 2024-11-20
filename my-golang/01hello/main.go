package main

import (
	"fmt"
	"net"
)

func main() {
	//var input string
	// fmt.Println("Enter and Ip")
	// fmt.Scanln(&input)
	// if validateIP(input) {
	// 	fmt.Println("Valid Ip")
	// } else {
	// 	fmt.Println("Not a valid Ip")
	// }
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	var b = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1968"}

	if checkMap(a, b) {
		fmt.Println("\nSame map")
	} else {
		fmt.Println("Not a same map")
	}

}

func checkMap(a map[string]string, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for key, val := range a {
		val1, ok := b[key]
		if ok && val1 == val {
			fmt.Printf("Present %v", val1)

		} else {
			fmt.Println("Not present")
			return false
		}
	}
	return true
}

func validateIP(input string) bool {
	var check net.IP
	check = net.ParseIP(input)
	fmt.Println("check ", check)
	return check != nil && check.To4() != nil
}
