package main

import "fmt"

func main() {
	var h int32 = 5
	fmt.Printf(`hello rajan,
	can we connect,
	thanks %d`, h)

	var m = int64(h)

	fmt.Printf(`hello rajan,
	can we connect,
	thanks %d`, m)

	name := person{
		"Rajan Singh",
	}
	name.sayHello()
}

type person struct {
	first string
}

func (name person) sayHello() {
	fmt.Println("Hello Mr.", name.first)
}
