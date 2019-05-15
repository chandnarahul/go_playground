package main

import (
	"fmt"
)

func main() {
	first := 10
	second := 20
	fmt.Printf("%d + %d = %d\n", first, second, (first + second))
	fmt.Printf("%d - %d = %d\n", first, second, (first - second))
	fmt.Printf("%d * %d = %d\n", first, second, (first * second))
	fmt.Printf("%d / %d = %d\n", first, second, (first / second))
}
