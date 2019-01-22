package main

import (
	"fmt"
)

func main() {
	topics := map[string]string{
		"easy":     "Program Structure",
		"inspired": "Basic and Composite types",
		"what?":    "Functions vs Methods",
		"cool":     "Goroutines",
		"crazy":    "Channels",
	}

	fmt.Printf("Let's have fun with: \n")
	for key, topic := range topics {
		fmt.Printf(" - \"%s\" which is: '%v' \n", topic, key)
	}
}
