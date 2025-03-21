package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/strategy"
)

type Strategy interface {
	GuessNumber(target, max int) int
}

type guess struct {
	strategy Strategy
}

func (g guess) execute(target, max int) int {
	return g.strategy.GuessNumber(target, max)
}

func main() {
	// 乱数の上限値
	max := 1000
	// 乱数生成機の生成
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// strategyの設定
	g := guess{strategy: strategy.BruteForce{}}

	start := time.Now()
	for i := 0; i < 100; i++ {
		// 対象の値
		t := r.Intn(max + 1)
		if rt := g.execute(t, max); t != rt {
			fmt.Println("failed")
		}
	}
	elapsedTime := time.Since(start)
	fmt.Printf("Loop execution time: %s\n", elapsedTime)
}
