// Alternative way of assigning a mutex
package main

import (
	"sync"
)

var count int;

func worker(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock();
	count += 1;
	mutex.Unlock();
	wg.Done();
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i:=0;i < 100; i++ {
		wg.Add(1);
		go worker(&wg, &mutex);
		}
	wg.Wait();
}
