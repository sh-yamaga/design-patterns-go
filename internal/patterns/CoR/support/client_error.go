package support

import (
	"fmt"

	cor "github.com/sh-yamaga/design-patterns-go/internal/patterns/CoR"
)

type ClientErrorResponseSupport struct {
	HttpResponseSupport
}

func NewClientErrorResponseSupport() *ClientErrorResponseSupport {
	cers := &ClientErrorResponseSupport{
		HttpResponseSupport{
			name:        "ClientErrorResponseSupport",
			supportFrom: 400,
			supportTo:   499,
		},
	}
	cers.ISupport = cers

	return cers
}

func (cers *ClientErrorResponseSupport) handle(hr *cor.HttpResponse) {
	fmt.Println("handle Response...")
	fmt.Println("Done")
}
