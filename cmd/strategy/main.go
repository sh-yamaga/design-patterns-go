package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/strategy"
)

type Strategy interface {
	GuessNumber(target, max int) int
	String() string
}

type guess struct {
	strategy Strategy
}

func (g guess) execute(target, max int) int {
	return g.strategy.GuessNumber(target, max)
}

func main() {
	max := 10_000
	var start time.Time
	// Random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Array of strategies for guessing random numbers
	var gs = [2]guess{
		{strategy: strategy.BruteForce{}},
		{strategy: strategy.BinarySearch{}},
	}

	for _, g := range gs {
		start = time.Now()
		for i := 0; i < 10_000; i++ {
			// generate random number
			target := r.Intn(max + 1)
			// g.execute returns an integer that is expected to be the target
			if result := g.execute(target, max); target != result {
				fmt.Println("failed")
			}
		}
		fmt.Printf("%s Time: %s\n", g.strategy.String(), time.Since(start))
	}
	// Output:
	// BruteForce Time: 13.972542ms
	// BinarySearch Time: 555.542µs
}
