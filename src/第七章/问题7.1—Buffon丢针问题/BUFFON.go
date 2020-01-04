package main

import (
	"fmt"
	"math/rand"
	"math"
)

func main() {
	var (
		n int
		count int
		x, theta, dist float64
		length = 1.0
		pi = math.Pi
	)

	fmt.Println("Buffon Needle Program by Simulation")
	fmt.Println("===================================")
	fmt.Print("Number of Simulation Points --> ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		x = float64(rand.Intn(10)) / 10
		theta = x * 2.0 * pi
		dist = x + length * math.Sin(theta)
		if 0.0 < dist && dist < length {
			count++
		}
	}
	fmt.Printf("Crossing Probability --> %.2f\n", 1.0 - float64(count) / float64(n));
    fmt.Printf("True Probability ------> %.2f\n", 2.0 / pi)
}