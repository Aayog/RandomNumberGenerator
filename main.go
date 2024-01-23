package main

import (
	"fmt"

	"aayogkoirala.com/rng"
)

func main() {
	generator := rng.NewRandomNumberGenerator(1) //1 is seed
	fmt.Println("Testing out the random generator.")
	fmt.Println("6 sided die roll 10 times:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d  ", generator.GetRandomInt(1, 6))
	}
	fmt.Println()
	fmt.Println("10 coin tosses")
	for i := 0; i < 10; i++ {
		coin := "T"
		if generator.GetRandomInt(1, 2) == 1 {
			coin = "H"
		}
		fmt.Printf("%s  ", coin)
	}
	fmt.Println()
	fmt.Println("100 random 1 digit numbers")
	for i := 0; i < 100; i++ {
		fmt.Printf("%d  ", generator.GetRandomInt(0, 9))
	}
	fmt.Println()
}
