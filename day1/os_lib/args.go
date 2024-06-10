package main

import (
	"fmt"
	"os"
)

func main() {
	args1 := os.Args
	args2 := os.Args[1:]
	fmt.Println("All args are: ", args2, args1)

	fmt.Println(args1[1]);
	fmt.Println(args1[2]);
	fmt.Println(args1[1:3]);
	fmt.Println(args1[:]);
}