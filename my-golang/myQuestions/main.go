package main

import (
	"fmt"
	"sync"
)

func main() {

	// arr := []int{1, 2, 3, 4,
	// test.A(arr, 1, 3)
	// fmt.Println("Array is ", arr)

	// fmt.Println(isFabNum(1000000))
	// fmt.Println(test.SquareOfOdd([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	// fmt.Println(isFabNum(1000000))

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	wg := sync.WaitGroup{}
	oddChan := make(chan []int)
	evenChan := make(chan []int)
	wg.Add(2)

	go splitOdd(arr, &wg, oddChan)
	go splitEven(arr, &wg, evenChan)
	defer wg.Wait()
	//oddArray := <-oddChan
	//evenArray := <-evenChan
	fmt.Println("oddArray", <-oddChan)
	fmt.Println("EvenArray", <-evenChan)

}

func isFabNum(input int) bool {
	if input < 0 {
		return false
	}
	firstNum := 0
	secondNum := 1
	fibNum := firstNum + secondNum
	for input >= fibNum {
		if input == fibNum {
			return true
		}
		firstNum = secondNum
		secondNum = fibNum
		fibNum = firstNum + secondNum
	}
	return false
}

func splitOdd(arr []int, wg *sync.WaitGroup, ch chan []int) {
	var oddArray []int
	for _, val := range arr {
		if val%2 == 1 {
			oddArray = append(oddArray, val)
		}
	}
	defer wg.Done()
	ch <- oddArray

}

func splitEven(arr []int, wg *sync.WaitGroup, ch chan []int) {
	var evenArray []int
	for _, val := range arr {
		if val%2 == 0 {
			evenArray = append(evenArray, val)
		}
	}
	defer wg.Done()
	ch <- evenArray
}
