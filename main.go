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

func LCG(generator RandomNumberGenerator, prev ...int64) int64 {
	if len(prev) > 0 {
		generator.seed = prev[0]
	}
	return (generator.multiplier*generator.seed + generator.increment) % generator.modulus
}

func getRandomInt(rngFunc func(RandomNumberGenerator, ...int64) int64, rng RandomNumberGenerator, seed int64, min int64, max int64) int64 {
	return rngFunc(rng, seed)%(max-min) + min
}

func main() {
	generator := RandomNumberGenerator{
		modulus:    int64(math.Pow(2, 32)),
		multiplier: 1664525,
		increment:  1013904223,
		seed:       1,
	}
	prev := getRandomInt(LCG, generator, 1000, 1, 10)
	for i := 0; i < 100; i++ {

		fmt.Println(prev)
		prev = getRandomInt(LCG, generator, 1, 1, 10)
	}
}
