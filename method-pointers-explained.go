package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X,Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3,4}
	Scale(&v, 10) // passa a referência do objeto, modificando-o no contexto da função main
	fmt.Println(Abs(v))
}