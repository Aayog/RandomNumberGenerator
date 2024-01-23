package main

import (
	"fmt"
	"math"
)

type RandomNumberGenerator struct {
	modulus    int64 // m -- range of random numbers generated, typically chosen in power of 2
	multiplier int64 // a -- affects how the algorithm traverses the range of possible values, a−1 is divisible by all prime factors of mm.
	increment  int64 // c -- ensures the generator doesn't produce cyclical patterns, If c and m are relatively prime (their greatest common divisor is 1), it helps in achieving a full period for all seed values.
	seed       int64 // X_0 -- initial value that starts the random number sequence; should ideally be a random or a semi-random (e.g., time-based) value
}

func (rng *RandomNumberGenerator) Next() int64 {
	rng.seed = (rng.multiplier*rng.seed + rng.increment) % rng.modulus
	return rng.seed
}

func getRandomInt(rng *RandomNumberGenerator, min int64, max int64) int64 {
	return rng.Next()%(max-min+1) + min
}

func main() {
	generator := RandomNumberGenerator{
		modulus:    int64(math.Pow(2, 32)),
		multiplier: 1664525,
		increment:  1013904223,
		seed:       100,
	}
	for i := 0; i < 100; i++ {
		randomInt := getRandomInt(&generator, 1, 100)
		fmt.Println(randomInt)
	}
}
