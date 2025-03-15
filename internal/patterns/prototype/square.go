package prototype

import (
	"fmt"
	"strings"
)

type Square struct {
	Side int
}

func (s *Square) Draw() {
	slice := make([]struct{}, s.Side)
	for i := range slice {
		if i == 0 || i == s.Side-1 {
			fmt.Println(strings.Repeat("*", s.Side))
		} else {
			fmt.Printf("*%s*\n", strings.Repeat(" ", s.Side-2))
		}
	}
}

func (s *Square) Clone() Shape {
	// 新しいSquareを作成し、元のSquareの値をコピーします。
	return &Square{Side: s.Side}
}
