package support

import (
	"fmt"

	cor "github.com/sh-yamaga/design-patterns-go/internal/patterns/CoR"
)

type RedirectionResponseSupport struct {
	HttpResponseSupport
}

func NewRedirectionResponseSupport() *RedirectionResponseSupport {
	rrs := &RedirectionResponseSupport{
		HttpResponseSupport{
			name:        "RedirectionResponseSupport",
			supportFrom: 300,
			supportTo:   399,
		},
	}
	rrs.ISupport = rrs

	return rrs
}

func (rrs *RedirectionResponseSupport) handle(hr *cor.HttpResponse) {
	fmt.Println("handle Response...")
	fmt.Println("Done")
}
