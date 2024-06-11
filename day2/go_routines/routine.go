package main

import (
	"fmt"
	"time"
)

func cpc() {
	for i:=1;i<=5;i++ {
		fmt.Println(i)
		time.Sleep(1*time.Second)
	}
}

func main() {
	go cpc()
	for i:=0;i<5;i++ {
		fmt.Println(i)
		time.Sleep(2*time.Second)
	}
}