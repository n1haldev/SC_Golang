package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	int_arr1 := []int{1, 2, 3, 4, 5, 6}
	int_arr2 := []int{1, 2, 3, 4, 5, 6}
	float_arr1 := []float32{1.2, 2.5, 3.6, 4.1}
	// float_arr2 := []float32{1.2, 2.5, 3.6, 4.1}
	string_arr1 := []string{"Hello", "Nice person", "Student"}
	string_arr2 := []string{"Hello", "Nice person", "Student"}

	// first op
	if reflect.DeepEqual(string_arr1, string_arr2) {
		fmt.Println("The two arrays are equal!")
	}

	// second op
	fmt.Println(unsafe.Sizeof(float_arr1));

	// third op
	fmt.Println(reflect.ValueOf(int_arr1).Len());

	// fourth op
	// for i := 0;i < len(int_arr2)/2 - 1; i++ {
	// 	temp := int_arr2[i]
	// 	int_arr2[i] = int_arr2[len(int_arr2) - 1 - i];
	// 	int_arr2[len(int_arr2) - 1 - i] = temp;
	// }

	// using reflect
	swap := reflect.Swapper(int_arr2);
	for i := 0;i < len(int_arr2)/2; i++ {
		swap(i, len(int_arr2) - 1 - i);
	}

	fmt.Println(int_arr2);
}