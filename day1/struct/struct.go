package main

import (
	"fmt"
	"time"
)

// type struct_name struct {
// member datatype;
// }

// func acceptor() -> Employeee {
// 	var ID, Sal, ManagerID int
// 	var Name, Address, Position string
// 	var DoB time.Time
	
// 	fmt.Println("Enter ID:");
// 	fmt.Scan("%d", &ID);
// 	fmt.Println("Enter Name:");
// 	fmt.Scan("%s", Name);
// 	fmt.Println("Enter Address:");
// 	fmt.Scan("%s", Address);
// 	fmt.Println("Enter DoB:");
// 	fmt.Scan("%t", DoB);
// 	fmt.Println("Enter Position:");
// 	fmt.Scan("%s", Position);
// 	fmt.Println("Enter Salary:");
// 	fmt.Scan("%d", &Sal);
// 	fmt.Println("Enter Manager ID:");
// 	fmt.Scan("%d", &ManagerID);

// 	return Person{
// 		ID
// 		Name
// 		Address
// 		DoB
// 		Position
// 		Sal
// 		ManagerID
// 	}
// }

// Eg:
type Person struct {
	Name string
	Age int
}

func main() {
	var p1 Person; // Declared a structure variable p1 with default values
	// var p2 Person; // Anynumber of structure variables can be declared
	
	var p2 = Person{Name: "Nihal", Age: 21}
	
	p3 := Person{Name: "Nihal", Age: 21}
	
	p4 := &Person{Name: "Nihal", Age: 21}
	
	fmt.Println(p1);
	fmt.Println(p2);
	fmt.Println(p3);
	fmt.Println(p4);
	
	p5 := struct {
		ID int
		Name string
	}{ID: 1, Name: "Nihal"}
	
	fmt.Println(p5)

	person1 := Person{"Nihal", 21}
	fmt.Println(person1)
	fmt.Println(person1.Age, person1.Name)

	type Employee struct {
		ID int
		Name string
		Address string
		DoB time.Time
		Position string
		Sal int
		ManagerID int
	}

	// var emp1, emp2 Employee;
	
}