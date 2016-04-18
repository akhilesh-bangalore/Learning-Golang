package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("\nHello from", runtime.GOOS)
}
