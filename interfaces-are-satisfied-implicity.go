package main
import "fmt"

type Interface interface {
	Method()
}

type Type struct {
	String string
}

func (t Type) Method() {
	fmt.Println(t.String)
}

func main() {
	var i Interface = Type{"hello"}
	i.Method()
}