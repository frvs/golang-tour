package main
import (
	"fmt"
	"math"
)

type Interface interface {
	Method()
}

type Type struct {
	String string
}

type MyFloat float64

func (f MyFloat) Method() {
	fmt.Println(f)
}

func (t Type) Method() {
	fmt.Println(t.String)
}

func describe(i Interface) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i Interface 
	
	i = &Type{"Hello"}
	describe(i)
	i.Method()

	i = MyFloat(math.Pi)
	describe(i)
	i.Method()
}