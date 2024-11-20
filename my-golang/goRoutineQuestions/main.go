package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// wg := &sync.WaitGroup{}
	// myCh := make(chan int)
	// mut := &sync.Mutex{}
	//go programs.HelloGo(&wg)

	// wg.Add(10)
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		fmt.Println(i)
	// 		wg.Done()
	// 	}()
	// }
	// defer wg.Wait()
	// fmt.Println("Program end")

	//passing channel to 2 threads
	// wg.Add(3)
	// go routineWithChannel(myCh, wg, mut)
	// go firstChannel(myCh, wg, mut)
	// go thirdfun(myCh, wg, mut)
	// wg.Wait()

	// wg.Add(2)
	// go func() {
	// 	defer wg.Done()
	// 	for i := 1; i <= 10; i++ {
	// 		myCh <- i
	// 		mut.Lock()
	// 		fmt.Printf("Sent: %d\n", i)
	// 		mut.Unlock()
	// 	}
	// 	close(myCh)
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	for v := range myCh {
	// 		mut.Lock()
	// 		fmt.Println("Value is ", v)
	// 		mut.Unlock()
	// 	}
	// }()
	// wg.Wait()

	pingChan := make(chan string)
	pongChan := make(chan string)

	// Goroutine to send "ping" every 1 second
	go func() {
		for {
			time.Sleep(1 * time.Second)
			pingChan <- "ping"
		}
	}()

	// Goroutine to send "pong" every 2 seconds
	go func() {
		for {
			time.Sleep(2 * time.Second)
			pongChan <- "pong"
		}
	}()

	// Main loop to listen to the channels using select
	for {
		select {
		case msg := <-pingChan:
			fmt.Println(msg)
		case msg := <-pongChan:
			fmt.Println(msg)
		}
	}

}

func firstChannel(myCh chan int, wg *sync.WaitGroup, mut *sync.Mutex) {
	mut.Lock()
	myCh <- 2
	mut.Unlock()
	close(myCh)
	defer wg.Done()
	println("Value of firstChannel is ", <-myCh)
}

func routineWithChannel(myCh chan int, wg *sync.WaitGroup, mut *sync.Mutex) {
	defer wg.Done()
	val := <-myCh
	mut.Lock()
	myCh <- 10
	mut.Unlock()
	close(myCh)
	println("Value is ", val)

}

func thirdfun(myCh chan int, wg *sync.WaitGroup, mut *sync.Mutex) {
	defer wg.Done()
	mut.Lock()
	myCh <- 3
	mut.Unlock()
	close(myCh)
	println("Value thirdfun is ", <-myCh)

}
