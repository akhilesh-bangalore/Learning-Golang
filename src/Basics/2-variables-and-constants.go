package main

import (
	"fmt"
	"runtime"
	"time"
)

var (
	OperatingSystem = runtime.GOOS
	Started         = time.Now()
)

const (
	Copyright = "Copyright (C) Akhilesh Bangalore"
)

func main() {
	var seatNo int
	var age int = 29   // Long declaration and assignment
	name := "Akhilesh" // Short declaration by type inference

	const Country = "India"

	fmt.Println("\nPackage Level Variables: ")
	fmt.Println("Operating System: ", OperatingSystem)
	fmt.Println("Program Started at: ", Started)

	fmt.Println("\nPackage Level Constants: ")
	fmt.Println("Copyright: ", Copyright)

	fmt.Println("\nMethod Level Variables: ")
	fmt.Println("Default variable initialization (seatNo): ", seatNo)
	fmt.Println("age: ", age)
	fmt.Println("name: ", name)

	fmt.Println("\nMethod Level Constants: ")
	fmt.Println("Country: ", Country)
}
