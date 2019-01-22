package main

import (
	"fmt"
)

func counter(naturals chan int) {
	for x := 0; x < 100; x++ {
		naturals <- x
	}
	close(naturals)
}

func squarer(naturals, squares chan int) {
	for x := range naturals {
		squares <- x * x
	}
	close(squares)
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go counter(naturals)

	// Squarer
	go squarer(naturals, squares)

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}
