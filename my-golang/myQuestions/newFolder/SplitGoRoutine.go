package newFolder

import "sync"

func splitOdd(arr []int, wg *sync.WaitGroup) []int {
	var oddArray []int
	for _, val := range arr {
		if val%2 == 1 {
			oddArray = append(oddArray, val)
		}
	}
	defer wg.Done()
	return oddArray
}

func splitEven(arr []int, wg *sync.WaitGroup) []int {
	var evenArray []int
	for _, val := range arr {
		if val%2 == 0 {
			evenArray = append(evenArray, val)
		}
	}
	defer wg.Done()
	return evenArray
}
