package main

import "fmt"

func main() {
	pow := make([]int, 10)

	fmt.Printf("len=%d cap=%d %v \n", len(pow), cap(pow), pow)

	for i:= range pow {
		pow[i] = 1 << uint(i)
	}
	fmt.Printf("len=%d cap=%d %v \n", len(pow), cap(pow), pow)
	
	for _, value := range pow {
		fmt.Printf("%d \n", value)
	}

}