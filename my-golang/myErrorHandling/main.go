package main

import (
	"errors"
	"fmt"
)

func main() {
	// val, err := returnErr(4, 0)
	// if err != nil {
	// 	fmt.Println("Error is ", err)
	// 	log.Fatal(err.Error())
	// } else {
	// 	fmt.Println("Val is ", val)
	// }

	fmt.Println("Start")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	returnPanic()
	fmt.Println("End")
}

func returnErr(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func returnPanic() {
	panic("something wrong")
}
