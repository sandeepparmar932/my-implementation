package main

import "fmt"

func mySwitchTypeInference(i interface{}) {

	switch v := i.(type) {
	case string:
		fmt.Println("String type")
	case int:
		fmt.Println("int type")
	case bool:
		fmt.Println("bool type")
	default:
		fmt.Println("Not defined - ", v)
	}
}

func main() {
	// mySwitchTypeInference(1)
	// mySwitchTypeInference("sandeep")
	// mySwitchTypeInference(true)

	var a interface{} = true

	b := a.(int)

	fmt.Println("b is ", b)

}
