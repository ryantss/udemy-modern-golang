package main

import (
	"fmt"
	"time"
)

func main() {
	qs := make(chan bool)
	go func() {
		fmt.Println("Hello from a new goroutine")
		qs <- true
	}()
	// go sayHelloFromGoroutine(qs)
	fmt.Println("Hello from main")
	v := <-qs
	fmt.Println(v)

	ic := make(chan int)
	go periodicsSend(ic)
	for i := range ic {
		fmt.Println(i)
	}
	_, ok := <-ic
	fmt.Println(ok)
}

// func sayHelloFromGoroutine(qs chan bool) {
// 	fmt.Println("Hello from a new goroutine")
// 	qs <- false
// }

func periodicsSend(ic chan int) {
	i := 0
	for i <= 10 {
		ic <- i
		i++
		time.Sleep(1 * time.Second)
	}
	close(ic)
}
