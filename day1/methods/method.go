package main  

import "fmt"

type Volume func(int, int, int) (int)

type Shape struct {
	length int
	breadth int
	height int
	color string
	cuboid Volume
}

func main() {
	res := Shape{length: 10, breadth: 7, height: 5, color: "Blue", 
	cuboid: func(length int, breadth int, height int) (int) {
		return length * breadth * height
	}};
	fmt.Println(res);


}