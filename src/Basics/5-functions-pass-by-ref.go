package main

import (
	"fmt"
)

func main() {
	name := "Akhilesh"
	fmt.Println("\nName before modify01: ", name)

	modify01(name)
	fmt.Println("\nName after modify01: ", name)

	modify02(&name)
	fmt.Println("\nName after modify02: ", name)

	// Slices are passed by reference
	// while arrays are passed by value
	data := []int{3, 4, 5, 6}
	fmt.Println("Before: ", data)
	modifySlice(data)
	fmt.Println("After: ", data)
}

func modify01(name string) {
	name = "Akhilesh 007"
}

func modify02(name *string) {
	*name = "Akhilesh 786"
}

func modifySlice(data []int) {
	data[0] = 35
}
