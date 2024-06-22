package main 

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup, n string) {
	fmt.Println("hello "+n)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup;
	wg.Add(2)
	go sayHello(&wg, "Alice");
	go sayHello(&wg, "Bob");
	wg.Wait()
	fmt.Println("Main is done with its job!")
}