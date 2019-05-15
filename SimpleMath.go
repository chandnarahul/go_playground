package main

import (
	"fmt"
)

func main() {
	first := 10
	second := 20
	fmt.Printf("%d + %d = %d", first, second, (first + second))
	fmt.Printf("%d - %d = %d", first, second, (first - second))
	fmt.Printf("%d * %d = %d", first, second, (first * second))
	fmt.Printf("%d / %d = %d", first, second, (first / second))
}
