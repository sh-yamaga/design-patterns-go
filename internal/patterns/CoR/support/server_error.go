package support

import (
	"fmt"
	"net/http"

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
	fmt.Println("===", sers.name, "===")
	fmt.Println(
		"Response Status Code:", hr.StatusCode,
		"(", http.StatusText(int(hr.StatusCode)), ")",
	)

	fmt.Println("handle Response...")
	fmt.Println("Done")
}
