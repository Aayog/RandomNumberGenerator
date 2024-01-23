package main

import (
	"fmt"
	"math"
)

func main() {
	generator := RandomNumberGenerator{
		modulus:    int64(math.Pow(2, 32)),
		multiplier: 1664525,
		increment:  1013904223,
		seed:       100,
	}
	for i := 0; i < 100; i++ {
		randomInt := generator.getRandomInt(1, 100)
		fmt.Println(randomInt)
	}
}
