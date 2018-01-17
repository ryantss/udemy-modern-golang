package main

import (
	"fmt"
)

func main(){
	s := struct{
		i int
		f float32
	}{i:3, f:3.3}
	fmt.Printf("%v\n", s)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%T\n", s)
	fmt.Printf("i is %d, while f is %f \n", s.i, s.f)

	fs := fmt.Sprintf("%+v\n", s)

	fmt.Println(fs)

}