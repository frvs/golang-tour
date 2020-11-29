package main

import "fmt"

type Interface interface {
	Method()
}

func describe(i Interface) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i Interface
	describe(i)
	i.Method()
}

