package main

import "fmt"

func main() {
	var I int = 3

	increment(&I)

	fmt.Println(I)
}

func increment(pI *int) {
	*pI++
}
