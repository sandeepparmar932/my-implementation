package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //pointers
var mut sync.Mutex    // similar to synchronized in java
var signal = []string{"test"}

func main() {
	endPointList := []string{
		"https://google.com",
		"http://localhost:8080/get",
		"https://github.com",
	}
	for _, value := range endPointList {
		go getStatusCode(value)
		wg.Add(1)
	}

	wg.Wait() //doesn't let the wait complete till get the done signal
	fmt.Print("Signal is ", signal)
}

func getStatusCode(endPoint string) {
	res, err := http.Get(endPoint)

	if err != nil {
		fmt.Println(err)
	} else {
		signal = append(signal, endPoint)
		fmt.Println("status is ", res.StatusCode)
	}
	defer wg.Done()

}
