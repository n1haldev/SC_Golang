package main

import "fmt"

func main() {
	// create a slice
	s1 := []int{10, 20, 30, 40, 50, 60};
	fmt.Println(s1);
	fmt.Println(s1[1:4]);

	// creating a slice with make
	slice1 := make([]int, 3, 5);
	fmt.Println(slice1);

	// len(), cap()
	// Appending elements to a slice
	slice1 = append(slice1, 1, 2, 3, 4, 5);
	fmt.Println(slice1);
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice1), cap(slice1));

	slice1[2] = 999
	fmt.Println(slice1)

	// method 1:
	for index, val := range slice1 {
		fmt.Println(index, val)
	}

	for _, val := range slice1 {
		fmt.Println(val)
	}

	// method 2:
	for i:=0;i<len(slice1);i++ {
		fmt.Println(slice1[i]);
	}

	// Appending slices
	slice2 := []int{1, 2, 3, 4, 5};
	slice3 := []int{6, 7, 8, 9, 10};
	slice2 = append(slice2, slice3...);
	// ... is used to unpack the elements of the slice slice3
	fmt.Println(slice2);

	var stack []string;
	y := "apple"
	stack = append(stack, y)
	fmt.Println(stack)

	complex := [...][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}
	fmt.Println(complex)

	aggregate := []int{};
	aggregate = append(aggregate, 1, 2, 3);


}