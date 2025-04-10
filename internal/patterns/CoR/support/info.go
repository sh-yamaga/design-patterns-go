package support

import (
	"fmt"

	cor "github.com/yamaga-shu/design-patterns-go/internal/patterns/CoR"
)

type InformationalResponseSupport struct {
	HttpResponseSupport
}

func NewInformationalResponseSupport() *InformationalResponseSupport {
	return &InformationalResponseSupport{
		HttpResponseSupport{
			name:        "InformationalResponseSupport",
			supportFrom: 100,
			supportTo:   199,
			next:        nil,
			handler:     informationalHandler{},
		},
	}
}

type informationalHandler struct{}

func (ih informationalHandler) do(hr *cor.HttpResponse) {
	fmt.Println("Recieved Informational Response")
	fmt.Println("Do Something...")
	fmt.Println("Done")
}
