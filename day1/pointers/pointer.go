package main 

import "fmt"

func main() {
	x := 1;
	// y := 0;
	p := &x;
	fmt.Println(p);
	fmt.Println(*p);

	*p = 999;
	fmt.Println(*p);
	fmt.Println(x);

	q := &x;
	fmt.Println(p == q);
}

// In Go, pointer arithmetic is not supported