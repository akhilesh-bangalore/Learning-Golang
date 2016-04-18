package main

import (
	"fmt"
)

func main() {
	fmt.Println("\nEven numbers:")
	printEven(1, 10)

	numbers := []int{5, 8, 12, 9, 4, 10}
	fmt.Println("\nMax of ", numbers, " is ", findMax(numbers))
	fmt.Println("\nSequence:")
	sequence(1, 10)

	fmt.Println("\nOdd numbers:")
	printOdd(1, 10)

	fmt.Println("\nBreaking outerloop")
	breakOuterLoop()
}

// looping with <initialization> <condition> <increment/decrement>
func printEven(start int, limit int) {
	for num := start; num < start+limit; num++ {
		if num%2 == 0 {
			fmt.Println(num)
		}
	}
}

// looping with range
func findMax(values []int) int {
	max := values[0]
	for _, val := range values {
		if val > max {
			max = val
		}
	}

	return max
}

// infinite loop - using break
func sequence(start int, limit int) {
	initial := start
	for {
		if initial == start+limit {
			break
		}
		fmt.Println(initial)
		initial++
	}
}

// using continue
func printOdd(start int, limit int) {
	for num := start; num < start+limit; num++ {
		if num%2 == 0 {
			continue
		}
		fmt.Println(num)
	}
}

// Shows how to break outer loops
// A thing to note is unused label will be detected as error
// Don't bother about the logic. Its all about syntax
func breakOuterLoop() {
Level1:
	for i := 0; i < 10; i++ {
	Level2:
		for j := 0; j < 10; j++ {
			if j > i {
				break Level1
			}
			for k := 0; k < 10; k++ {
				break Level2
			}
		}
	}
}
