package main

import (
	"fmt"
	"math"
)

func sum(a,b int) (int) {
	return a + b
}

func sayHello(name string) {
	fmt.Println("Hello ", name);
}

func names(n[]string, fun func(string)) {
	for _, v := range n {
		fun(v)
	}
}

func circleArea(n float32) (float32) {
	return math.Pi * n * n;
}

func main() {
	result := sum(1, 2)
	fmt.Println("1 + 2 = ", result)

	// fmt.Println("Enter your name");
	// var name string;
	// fmt.Scanf("%s", &name);
	// sayHello(name);

	names([]string{"John", "Doe", "Jane"}, sayHello);
	fmt.Printf("The area of the circle is: %f", circleArea(6.0));
}