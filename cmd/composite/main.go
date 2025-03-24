package main

import "fmt"

type Area interface {
	GetPopulation() uint
}

func main() {
	jp := NewCountry()
	jp.AddPrefecture(&Prefecture{
		name:       "tokyo",
		population: 14_200_331,
	})
	jp.AddPrefecture(&Prefecture{
		name:       "hokkaido",
		population: 5_039_100,
	})

	fmt.Println(jp.GetPopulation())
}
