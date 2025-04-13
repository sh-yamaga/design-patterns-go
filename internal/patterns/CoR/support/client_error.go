package support

import (
	"fmt"

	cor "github.com/yamaga-shu/design-patterns-go/internal/patterns/CoR"
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

func (ceh clientErrorHandler) handleResponse(hr *cor.HttpResponse) {
	fmt.Println("Recieved Client Error Response")
	fmt.Println("Do Something...")
	fmt.Println("Done")
}
