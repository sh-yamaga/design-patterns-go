package support

import (
	"fmt"

	cor "github.com/yamaga-shu/design-patterns-go/internal/patterns/CoR"
)

type SuccessfulResponseSupport struct {
	HttpResponseSupport
}

func NewSuccessResponseSupport() *SuccessfulResponseSupport {
	return &SuccessfulResponseSupport{
		HttpResponseSupport{
			name:        "SuccessfulResponseSupport",
			supportFrom: 200,
			supportTo:   299,
			next:        nil,
			handler:     successHandler{},
		},
	}
}

type successHandler struct{}

func (sh successHandler) do(hr *cor.HttpResponse) {
	fmt.Println("Recieved Success Response")
	fmt.Println("Do Something...")
	fmt.Println("Done")
}
