package main

import (
	"fmt"
)

func main() {
	buffch := make(chan int, 2)
	buffch <- 3
	buffch <- 2
	fmt.Println(<-buffch)
	fmt.Println(<-buffch)
	fmt.Println(<-buffch)
}
