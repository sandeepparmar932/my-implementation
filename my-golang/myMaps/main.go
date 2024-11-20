package main

import (
	"fmt"
)

func main() {
	dummy1 := map[string]int{"a": 1, "b": 2, "c": 3}

	//fmt.Println("map is ", dummy1)
	//looping
	// for key, value := range dummy1 {
	// 	fmt.Printf("key is %v and value is %v\n", key, value)
	// }
	dummy2 := map[string]int{"a": 1, "b": 2, "c": 4}

	for key, value := range dummy1 {
		value2, available := dummy2[key]
		if !available || value != value2 {
			fmt.Println("value not avalable is ", value2, available)
		}

	}

	//delete value
	//delete(dummy, "lmn")

}
