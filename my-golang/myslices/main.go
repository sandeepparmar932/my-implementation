package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	sliceList := []string{"abc", "xyz"}
	//fmt.Println("slicelist is ", sliceList, " size is ", len(sliceList))
	sliceList = append(sliceList, "lmn")
	//fmt.Println("slicelist is ", sliceList, " size is ", len(sliceList))

	sliceList = append(sliceList[:1], sliceList[2:]...) //deleting a value in slice
	numList := make([]int, 3)
	numList[2] = 100
	numList[0] = 101
	numList[1] = 102
	//fmt.Println(numList) //will give arrayoutofbound for numList[3]

	numList = append(numList, 103)
	//fmt.Println(numList)

	sort.Ints(numList)
	n, isAvailable := slices.BinarySearch(numList, 100) // index and true if available
	fmt.Println("Binary search - ", n, isAvailable)

	newList := slices.Clone(numList)
	fmt.Println("NumList is - ", newList)

	numList = slices.Concat(newList, numList)
	fmt.Println(numList)

	newSlice := []int{5}

	newSlice = append(newSlice, 1)
	fmt.Println("new slice is -", newSlice)

}
