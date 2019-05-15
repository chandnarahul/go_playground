package main

import (
	"fmt"
	"time"
)

func main() {
	currentAge := 30
	retirementAge := 60
	fmt.Printf("You have %d  years left until you can retire.", (retirementAge - currentAge))

	currentYear := time.Now().Year()
	fmt.Println("")
	fmt.Printf("It's %d, so you can retire in %d.", currentYear, currentYear+(retirementAge-currentAge))
}
