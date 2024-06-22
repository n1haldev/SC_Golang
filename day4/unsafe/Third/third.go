package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	id int
	name string
	role string
	salary float32
}

type Company struct {
	name string
	location string
	employees []Employee
}

func main() {
	employees := []Employee{
		{
			id: 3,
			name: "Nihal",
			role: "SDE",
			salary: 987654321,
		},
		{
			id: 1,
			name: "Digvijay",
			role: "Senior Developer",
			salary: 99999999999,
		},
	}

	Random_Company := Company{
		name: "Generic Company",
		location: "Bangalore",
		employees: employees,
	}

	// 1. Accessing embedded structure members
	fmt.Println(reflect.ValueOf(Random_Company.employees[0].id))
	fmt.Println(reflect.ValueOf(Random_Company.employees[0].name))
	fmt.Println(reflect.ValueOf(Random_Company.employees[0].role))
	fmt.Println(reflect.ValueOf(Random_Company.employees[0].salary))

	fmt.Println(reflect.ValueOf(Random_Company.employees[1].id))
	fmt.Println(reflect.ValueOf(Random_Company.employees[1].name))
	fmt.Println(reflect.ValueOf(Random_Company.employees[1].role))
	fmt.Println(reflect.ValueOf(Random_Company.employees[1].salary))
	
	// 2. Access structure members
	fmt.Println(reflect.ValueOf(Random_Company.name));
	fmt.Println(reflect.ValueOf(Random_Company.location));
	fmt.Println(reflect.ValueOf(Random_Company.employees));

	// 3. Accessing structure members randomly
	fmt.Println(reflect.ValueOf(Random_Company.employees[0].id))
	fmt.Println(reflect.ValueOf(Random_Company.employees[0].name))
	fmt.Println(reflect.ValueOf(Random_Company.employees[1].role))
	fmt.Println(reflect.ValueOf(Random_Company.employees[1].salary))

	// 4. Structure members are equal or not
	new_Employee := Employee{
		id: 3,
		name: "Nihal",
		role: "SDE",
		salary: 987654321,
	}

	flag := 0;
	for emp := range Random_Company.employees {
		if reflect.DeepEqual(new_Employee, emp) {
			flag = 1;
			fmt.Println("This person is already in the company bro!");
		}
	}

	if flag == 0 {
		fmt.Println("Welcome to ", Random_Company.name)
	}
}