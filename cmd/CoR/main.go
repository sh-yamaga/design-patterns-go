package main

import (
	cor "github.com/sh-yamaga/design-patterns-go/internal/patterns/CoR"
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/CoR/support"
)

func main() {
	res := []cor.HttpResponse{
		// 100-199
		{StatusCode: 100},
		// 200-299
		{StatusCode: 201},
	}

	info := support.NewInformationalResponseSupport()
	success := support.NewSuccessResponseSupport()

	info.SetNext(success)

	for _, r := range res {
		info.Resolve(&r)
	}
}
