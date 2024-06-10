package main

// strconv - strings to numeric type

import (
	"fmt"
	"strconv"
)

func main() {
	s := "Hello,world!!!";
	fmt.Println(len(s));
	fmt.Println(s[0], s[7]);
	fmt.Println(s[0:5]);
	fmt.Println(s[7:]);
	fmt.Println(s[:5] + " bois");

	x := "123"
	y, err := strconv.Atoi(x);
	if err != nil { 
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(y);
	}

	a := strconv.Itoa(y);
	fmt.Println(a);

	z := "123.45";
	res, err := strconv.ParseFloat(z, 32);
	if err != nil {
		fmt.Println("Error: ", err);
	} else {
		fmt.Println(res);
	}
}