package rng

import "time"

type randomNumberGenerator struct {
	modulus    int64 // m -- range of random numbers generated, typically chosen in power of 2
	multiplier int64 // a -- affects how the algorithm traverses the range of possible values, a−1 is divisible by all prime factors of mm.
	increment  int64 // c -- ensures the generator doesn't produce cyclical patterns, If c and m are relatively prime (their greatest common divisor is 1), it helps in achieving a full period for all seed values.
	seed       int64 // X_0 -- initial value that starts the random number sequence; should ideally be a random or a semi-random (e.g., time-based) value
}

// NewRandomNumberGenerator creates a RandomNumberGenerator with optional parameters.
// The expected order of parameters is modulus, multiplier, increment, seed.
func NewRandomNumberGenerator(prev ...int64) *randomNumberGenerator {
	// Set defaults
	gen := randomNumberGenerator{
		modulus:    1 << 32, // 2^31
		multiplier: 1664525,
		increment:  1013904223,
		seed:       time.Now().UnixMilli(),
	}

	// Override defaults based on provided arguments
	if len(prev) > 0 {
		gen.modulus = prev[0]
	}
	if len(prev) > 1 {
		gen.multiplier = prev[1]
	}
	if len(prev) > 2 {
		gen.increment = prev[2]
	}
	if len(prev) > 3 {
		gen.seed = prev[3]
	}
	return &gen
}

func (rng *randomNumberGenerator) Next() int64 {
	rng.seed = LCG(rng)
	return rng.seed
}

func (rng *randomNumberGenerator) GetRandomInt(min int64, max int64) int64 {
	return rng.Next()%(max-min+1) + min
}

func LCG(rng *randomNumberGenerator) int64 {
	return (rng.multiplier*rng.seed + rng.increment) % rng.modulus
}
