package main

import (
	"fmt"
	"time"
)

func goroutine(ch chan int) {
	for {
		fmt.Printf("%d\n", <-ch)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int)

	go goroutine(ch)

	for i := 0; i < 10; i++ {
		ch <- i
	}
	fmt.Println(<-ch)
}
