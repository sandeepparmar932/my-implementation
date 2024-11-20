package programs

import (
	"fmt"
	"sync"
)

func HelloGo(wg *sync.WaitGroup) {
	fmt.Println("Hello from Goroutine!")
	wg.Done()
}

func AnonymouseGoRoutine(wg *sync.WaitGroup) {
	func() {
		for i := 0; i < 10; i++ {
			go fmt.Print(i)
		}
		wg.Done()
	}()
	wg.Wait()
}
