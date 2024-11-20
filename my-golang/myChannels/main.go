package main

import (
	"sync"
)

func main() {
	myCh := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		//fmt.Println(<-myCh)
		//fmt.Println(<-ch)
		val, isAva := <-ch

		println("val ", val, isAva)

		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		myCh <- 6
		close(myCh)
		//ch <- 5
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
