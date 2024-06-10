package main

import "fmt"

func print(args ...interface{}) (int, error) {
    return fmt.Println(args...)
}

func main() {
	var a [5]int
	fmt.Println("emp: ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	fmt.Println("len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// declare an array
	// 1. var array_name = [length]datatype{values} - length is fixed (OR)
	// var array_name = [...]datatype{values} - length is inferred

	// 2. array_name := [length]datatype{values} - length is fixed (OR
	// array_name := [...]datatype{values} - length is inferred
	var array = [10]int{5,3,2,1,5,2,5,2,4,2};
	fmt.Println("Array: ", array);

	// 3. var array_name [length]datatype - length is fixed
	// array_name = [values]
	print("Hello")
	k := [...]int{99:100};
	print(len(k))
	print(k)
	print("2" == "2")
	a1 := [2]int{1,2}
	b1 := [...]int{1,2}
	c1 := [2]int{1,3}

	print(a1 == b1)
	print(a1 == c1)
	print(b1 == c1)
}