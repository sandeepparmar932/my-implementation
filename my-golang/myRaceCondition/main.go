package main

import (
	"fmt"
	"sync"
)

func main() {
	var intArr = []int{0}
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("This is first")
		mut.Lock()
		intArr = append(intArr, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("This is second")
		mut.Lock()
		intArr = append(intArr, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("This is third")
		mut.Lock()
		intArr = append(intArr, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	wg.Wait()
	fmt.Println("int array is ", intArr)
}
