package main 

import (
	"fmt"
	"unsafe"
)

type College struct {
	courses []string
	rank int
	location string
}

func main() {
	PES := College{
		courses: []string{"CSE", "ECE"},
		rank: 83,
		location: "Electronic City",
	}

	// 1. Access struct members
	fmt.Println("Courses offered: ", PES.courses);
	fmt.Println("Rank: ", PES.rank);
	fmt.Println("Location: ", PES.location);

	// 2. Find the size of the structure
	fmt.Println(unsafe.Sizeof(PES));

	// 3. Performing Alignof and Offsestof of the struct members
	fmt.Println("Align of usage for courses: ", unsafe.Alignof(PES.courses));
	fmt.Println("Align of usage for rank: ", unsafe.Alignof(PES.rank));
	fmt.Println("Align of usage for location: ", unsafe.Alignof(PES.location));

	fmt.Println("Offset of usage for courses: ", unsafe.Offsetof(PES.courses));
	fmt.Println("Offset of usage for rank: ", unsafe.Offsetof(PES.rank));
	fmt.Println("Offset of usage for location: ", unsafe.Offsetof(PES.location));
}