package support

import (
	"fmt"

	cor "github.com/sh-yamaga/design-patterns-go/internal/patterns/CoR"
)

type ServerErrorResponseSupport struct {
	HttpResponseSupport
}

func NewServerErrorResponseSupport() *ServerErrorResponseSupport {
	sers := &ServerErrorResponseSupport{
		HttpResponseSupport{
			name:        "ServerErrorResponseSupport",
			supportFrom: 500,
			supportTo:   599,
		},
	}
	sers.ISupport = sers

	return sers
}

func (sers *ServerErrorResponseSupport) handle(hr *cor.HttpResponse) {
	fmt.Println("handle Response...")
	fmt.Println("Done")
}
