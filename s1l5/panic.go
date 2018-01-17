package main

import (
	"fmt"
	"iobuf"
)

func main() {
	// defer fmt.Println("World 1")
	// defer fmt.Println("World 1")

	fmt.Println("Hello")
	testpanic()
	fmt.Println("LALAL WORLD")
}

func testpanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("We recovered from a panic")
		}
	}()
	panic("lalala")
}

// type SchoolTime struct {
// 	From time.Time
// 	To   time.Time
// }

// func main() {
// 	var primary = SchoolTime{time.Now(), time.Now()}

// 	fmt.Println(primary)
// }
