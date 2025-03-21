package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/strategy"
)

type Strategy interface {
	GuessNumber(target, max int) int
	Name() string
}

type guess struct {
	strategy Strategy
}

func (g guess) execute(target, max int) int {
	return g.strategy.GuessNumber(target, max)
}

func main() {
	max := 100_000
	var start time.Time
	// Random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var gs = [2]guess{
		{strategy: strategy.BruteForce{}},
		{strategy: strategy.BinarySearch{}},
	}

	for _, g := range gs {
		start = time.Now()
		for i := 0; i < 100; i++ {
			t := r.Intn(max + 1)
			if rt := g.execute(t, max); t != rt {
				fmt.Println("failed")
			}
		}
		fmt.Printf("%s Time: %s\n", g.strategy.Name(), time.Since(start))
	}
}
