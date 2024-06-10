package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, " is negative")
	} else if num < 10 {
		fmt.Println(num, " has 1 digit")
	} else {
		fmt.Println(num, " has multiple digits")
	}

	// switch case
	exp := 5%3
	switch exp {
	case 0:
		fmt.Println("0");
		break;
	case 1:
		fmt.Println("1");
		break;
	case 2:
		fmt.Println("2");
		break;
	}

	// multicase switch
	switch exp {
	case 0, 1:
		fmt.Println("0 or 1");
		break;
	case 2, 3:
		fmt.Println("2 or 3");
		break;
	default: 
		fmt.Println("Default");
		break;
	}
}