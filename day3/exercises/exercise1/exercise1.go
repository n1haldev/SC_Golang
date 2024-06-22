package main 

import (
	"fmt"
	"sync"
)

func main() {
	var counter int 
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0;i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0;j < 1000; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("The final value is: ", counter);
}