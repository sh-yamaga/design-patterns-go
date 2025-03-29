package support

import (
	"fmt"

	cor "github.com/sh-yamaga/design-patterns-go/internal/patterns/CoR"
)

type ClientErrorResponseSupport struct {
	HttpResponseSupport
}

func NewClientErrorResponseSupport() *ClientErrorResponseSupport {
	return &ClientErrorResponseSupport{
		HttpResponseSupport{
			name:        "ClientErrorResponseSupport",
			supportFrom: 400,
			supportTo:   499,
			next:        nil,
			handler:     clientErrorHandler{},
		},
	}
}

type clientErrorHandler struct{}

func (ceh clientErrorHandler) do(hr *cor.HttpResponse) {
	fmt.Println("handle Response...")
	fmt.Println("Done")
}
