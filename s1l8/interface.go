package main

import (
	"fmt"
)

type testiface interface{
	SayHello()
	Say(s string)
	Increment()
	GetInternalValue() int
}
type testConcreteImpl struct{
	i int
}

func(tst testConcreteImpl) SayHello(){
	fmt.Println("Hello")
}
func(tst testConcreteImpl) Say(s string){
	fmt.Println(s)
}
func(tst *testConcreteImpl) Increment(){
	tst.i++
}
func(tst testConcreteImpl) GetInternalValue() int{
	return tst.i
}
func NewTestConcreteImpl() testiface{ // this is constructor
	return new(testConcreteImpl) //same as &testConcreteImpl{}
}
func NewTestConcreteImplWithV(v int) testiface{
	return &testConcreteImpl{i: v}
}

type testEmbedding struct { //want this struct to have all the features of testConcreteImpl. This is called the outer type
	*testConcreteImpl //embedding, this is called the inner type
}

func testEIface(v interface{}){
	fmt.Println(v)
}
func main(){
	
	var tiface testiface
	// tiface = &testConcreteImpl{}
	// tiface = NewTestConcreteImpl()
	tiface = NewTestConcreteImplWithV(5)
	tiface.SayHello()
	tiface.Say("Hello again!!")
	tiface.Increment()
	tiface.Increment()
	fmt.Println(tiface.GetInternalValue())

	
	te:= testEmbedding{testConcreteImpl: &testConcreteImpl{i:50}}
	te.SayHello()
	te.Increment()
	fmt.Println(te.GetInternalValue())
	testEIface(3)
	testEIface("string to empty interface")
	testEIface(tiface)
}