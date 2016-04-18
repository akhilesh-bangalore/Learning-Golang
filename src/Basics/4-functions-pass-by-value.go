package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 45
	b := 0

	fmt.Printf("\nMath operations on a=%d and b=%d\n", a, b)
	fmt.Println("Add: ", add(a, b))
	fmt.Println("Sub: ", subtract(a, b))
	fmt.Println("Multiply: ", multiply(a, b))

	var c, err = divide(a, b)
	fmt.Println("Divide: Q=", c, "Error=", err)

	fmt.Println("Variadic Add: ", addVariadic(10, 20, 30, 40, 50, 60, 70, 80, 90, 100))
}

// func <name>(<parameters>) <return type>
func add(a int, b int) int {
	return a + b
}

func addVariadic(a int, b int, more ...int) int {
	var result = add(a, b)
	for _, val := range more {
		result = add(result, val)
	}
	return result
}

func subtract(a, b int) int {
	return a - b
}

// func with named return value
func multiply(a, b int) (result int) {
	return a * b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by 0")
	}

	return a / b, nil
}
