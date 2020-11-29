package main

import "fmt"

func do(i interface{}) {

	switch v := i.(type) {
	case int:
		fmt.Println("integer", v)
	case string:
		fmt.Println("string", v)
	case bool:
		fmt.Println("boolean", v)
	default:
		fmt.Println("idk about type!", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}