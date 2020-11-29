package main

import (
	"fmt"
)

type Interface interface {
	Method()
}

type Type struct {
	String string
}


func (t *Type) Method() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.String)
}

func describe(i Interface) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i Interface

	var t *Type
	i = t
	describe(i)
	i.Method()
}