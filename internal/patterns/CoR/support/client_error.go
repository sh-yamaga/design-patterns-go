package support

import (
	"fmt"
	"net/http"

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
	fmt.Println("===", cers.name, "===")
	fmt.Println(
		"Response Status Code:", hr.StatusCode,
		"(", http.StatusText(int(hr.StatusCode)), ")",
	)

	fmt.Println("handle Response...")
	fmt.Println("Done")
}
