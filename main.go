package main

import (
	"fmt"

	p "github.com/arnald/adventofcode/challenges/dayOne/partOne"
)

func main() {
	totalDistance, err := p.TotalDistance()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Total Distance: %d\n", totalDistance)
}
