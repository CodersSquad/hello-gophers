// Netcat2 is a read-only TCP client with channels
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan int)
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- 2 // signal the main goroutine
	}()

	x := 1
	x = <-done // wait for background goroutine to finish
	log.Println("Channel Closed with value: ", x)
	close(done)
}
