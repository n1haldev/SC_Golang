package main

import (
	"fmt"
	"reflect"
)

func main() {
	strings_arr := []string{"Go programming", "Golang"};

	copy_arr := make([]string, len(strings_arr))
	// var copy_arr []string;

	reflect.Copy(reflect.ValueOf(copy_arr), reflect.ValueOf(strings_arr));
	fmt.Println(copy_arr);
}