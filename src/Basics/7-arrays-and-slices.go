package main

import (
	"fmt"
)

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("I am an array", array, "with length", len(array), "and capacity", cap(array))

	// 1:3 => 1 is inclusive and 3 is exclusive
	// The below slice will have [2, 3, 4]
	slice := array[1:4]
	fmt.Println("I am a slice", slice, "with length", len(slice), "and capacity", cap(slice))

	// Making slice without explicitly defining the array.
	// make(type, length, capacity)
	// Go will create an array automatically. When append is called on an array
	// that has reached the capacity limit, Go will create a new array which is double the
	// the times of its initial size
	slice2 := make([]int, 5)
	for i := 1; i < 10; i++ {
		fmt.Println("len=", len(slice2), "cap=", cap(slice2))
		slice2 = append(slice2, i)
	}

	// appending slices
	combinedSlice := append(slice, slice2...)
	fmt.Println("Combined Slice: ", combinedSlice)
}
