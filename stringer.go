package main

import "fmt"

type Person struct {
	Name string 
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Lucas", 21}
	z := Person{"Leela", 30}
	fmt.Println(a,z)
}