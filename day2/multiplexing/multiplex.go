package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan string, 10);
	ch2 := make(chan string, 10);

	var wg sync.WaitGroup;
	wg.Add(2);

	go func() {
		defer wg.Done();
		time.Sleep(1*time.Second);
		ch1<-"Hello"
	}()

	go func() {
		defer wg.Done();
		time.Sleep(2*time.Second);
		ch2<-"World"
	}()

	// go func() {
	// 	wg.Wait();
	// 	close(ch1)
	// 	close(ch2)
	// }()

	for i:=0;i < 2; i++ {
		select {
			case msg1 := <-ch1:
				fmt.Println(msg1)
			case msg2 := <-ch2:
				fmt.Println(msg2)
			default:
				fmt.Println("No message received")
		}
	}
}