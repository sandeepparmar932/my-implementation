package newFolder

import "fmt"

func Devisible(num int) bool {
	return num%7 == 0 && num%5 != 0
}

func SquareOfOdd(arr []int) []int {
	var tempArr []int
	if len(arr) == 0 {
		return arr
	}
	for index, val := range arr {
		fmt.Printf("Value is %v and index is %v\n", val, index)
		if val%2 != 0 {
			tempArr = append(tempArr, val*val)
		}
	}

	return tempArr
}
