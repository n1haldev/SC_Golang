package main 

import (
	"fmt"
	"sync"
)

type Counter struct{
	count int
	mutex sync.Mutex
	wg sync.WaitGroup
}

func worker(counter *Counter) {
	defer counter.wg.Done()

	for i := 0;i < 1000; i++ {
		counter.mutex.Lock()
		counter.count++;
		counter.mutex.Unlock()
	}
}

func main() {
	counter1 := Counter{}
	counter2 := Counter{}

	counter1.wg.Add(5)
	counter2.wg.Add(5)

	for i := 0;i < 5; i++ {
		go worker(&counter1)
		go worker(&counter2)
	}

	counter1.wg.Wait()
	counter2.wg.Wait()

	fmt.Println("Final value of counter 1: ", counter1.count);
	fmt.Println("Final value of counter 2: ", counter2.count);
}