package main 

import (
	"fmt"
	"time"
)

func execute(n string) {
	for i := 0;i < 10; i++ {
		fmt.Println(n, i);
		time.Sleep(time.Second)
	}
}

func main() {
	go execute("First")
	go execute("Second")
	fmt.Println("Main Exited!")
}