package main

import "fmt"

func main() {
	// Declaring variabels
	// var name data_type = value
	// var s string = "Computer Science"
	// var s2 = "Automatic Assignment"
	// s3 := "this is called a walrus operator"
	var a int
	var b bool
	var c string
	var d float32
	fmt.Println(a, b, c, d)

	var rr, ec, hn = "ring road", "electronic city", "hanumanthnagar"
	fmt.Println(rr, ec, hn)

	var1 := "Hi"
	// variable shadowing
	fmt.Println(var1)
	var1 = "Yo"
}
