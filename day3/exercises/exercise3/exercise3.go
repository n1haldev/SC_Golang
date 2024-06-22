package main 

import (
	"fmt"
	"sync"
)

var arr[] int = []int{1,2,3,4,5}
var mutex sync.Mutex
var wg sync.WaitGroup

func reader(a,b int) {
	defer wg.Done();
	mutex.Lock();
	defer mutex.Unlock();
	fmt.Println(arr[a:b]);
}

func writer(ele int, index int) {
	defer wg.Done();
	mutex.Lock();
	defer mutex.Unlock();
	arr[index] = ele;
}

func main() {

	wg.Add(5);
	go func() {
		reader(0,5)
		writer(1,1)
		reader(0,5)
		writer(4,4)
		reader(0,5)
	}()

	wg.Wait();
}