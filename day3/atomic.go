package main 

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var ops atomic.Uint64
	// atomic.Uint64 - Type from the sync/atomic package
	// Integer ops can be accessed and modified atomically(as a single indivisible operation)
}