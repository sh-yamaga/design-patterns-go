package support

import (
	"fmt"

	cor "github.com/yamaga-shu/design-patterns-go/internal/patterns/CoR"
)

type ServerErrorResponseSupport struct {
	HttpResponseSupport
}

func NewServerErrorResponseSupport() *ServerErrorResponseSupport {
	return &ServerErrorResponseSupport{
		HttpResponseSupport{
			name:        "ServerErrorResponseSupport",
			supportFrom: 500,
			supportTo:   599,
			next:        nil,
			handler:     serverErrorHandler{},
		},
	}
}

type serverErrorHandler struct{}

func (seh serverErrorHandler) do(hr *cor.HttpResponse) {
	fmt.Println("Recieved Server Error Response")
	fmt.Println("Do Something...")
	fmt.Println("Done")
}
