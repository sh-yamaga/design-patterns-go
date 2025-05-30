package support

import (
	"fmt"

	cor "github.com/yamaga-shu/design-patterns-go/internal/patterns/CoR"
)

type RedirectionResponseSupport struct {
	HttpResponseSupport
}

func NewRedirectionResponseSupport() *RedirectionResponseSupport {
	return &RedirectionResponseSupport{
		HttpResponseSupport{
			name:        "RedirectionResponseSupport",
			supportFrom: 300,
			supportTo:   399,
			next:        nil,
			handler:     redirectHandler{},
		},
	}
}

type redirectHandler struct{}

func (rh redirectHandler) handleResponse(hr *cor.HttpResponse) {
	fmt.Println("Recieved Redirection Response")
	fmt.Println("Do Something...")
	fmt.Println("Done")
}
