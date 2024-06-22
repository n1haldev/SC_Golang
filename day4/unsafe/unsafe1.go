package main

import (
	"fmt"
	"unsafe"
)

type Student struct{
	name string
	id int
}

func main() {
	w := []int{4,5,3,6,8,3,2};
	ptr := &w[0];
	// Conversion of pointer to unsafe.Pointer and perform arithmetic operations
	nextptr := unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(w[0]))
	fmt.Println(*(*int)(nextptr))

	student1 := Student{id:3, name: "qwertyuiopASGHJKL;ZXCVBNM"};
	// student2 := Student{id:3, name:"Nihal"};

	// fmt.Println("Align of usage: ", unsafe.Alignof(student1));
	fmt.Println("Align of usage: ", unsafe.Alignof(student1.name));
	fmt.Println("Align of usage: ", unsafe.Alignof(student1.id));

	// fmt.Println("Offset of usage: ", unsafe.Offsetof(student1));
	fmt.Println("Offset of usage: ", unsafe.Offsetof(student1.name));
	fmt.Println("Offset of usage: ", unsafe.Offsetof(student1.id));
}