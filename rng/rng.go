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
    gen := randomNumberGenerator{
        modulus:    1 << 32, // Equivalent to 2^32
        multiplier: 16_645_215,
        increment:  1_013_904_223,
        seed:       time.Now().UnixMicro() & 0x7FFFFFFFFFFFFFFF, // Ensures seed is non-negative
    }

    switch len(prev) {
    case 4:
        gen.modulus = prev[3]
        fallthrough
    case 3:
        gen.increment = prev[2]
        fallthrough
    case 2:
        gen.multiplier = prev[1]
        fallthrough
    case 1:
        gen.seed = prev[0] & 0x7FFFFFFFFFFFFFFF // Ensures seed is non-negative
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
