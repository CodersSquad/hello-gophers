package main

import (
	"fmt"
	"os"
	"sync"
	//"time"
)

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

func change(s *secret, newPassword string) {
	s.RWM.Lock()
	fmt.Println("change")
	//time.Sleep(3 * time.Second)
	s.password = newPassword
	s.RWM.Unlock()
}

func show(s *secret) string {
	s.RWM.RLock()
	fmt.Println("show")
	//time.Sleep(1 * time.Second)
	defer s.RWM.RUnlock()
	return s.password
}

func showWithLock(s *secret) string {
	s.M.Lock()
	fmt.Println("showWithLock")
	//time.Sleep(1 * time.Second)
	defer s.M.Unlock()
	return s.password
}

func main() {

	password := secret{password: "myPassword"}

	showFunction := func(s *secret) string { return "" }

	if len(os.Args) != 2 {
		fmt.Println("Using Sync.RWMutex")
		showFunction = show
	} else {
		fmt.Println("Using Sync.Mutex")
		showFunction = showWithLock
	}

	fmt.Println("Password: ", showFunction(&password))
	var waitGroup sync.WaitGroup

	for i := 0; i < 15; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			fmt.Println("Password: ", showFunction(&password))
		}()

		go func() {
			waitGroup.Add(1)
			defer waitGroup.Done()
			change(&password, "123456")
		}()
	}

	waitGroup.Wait()
	fmt.Println("Password: ", showFunction(&password))
}
