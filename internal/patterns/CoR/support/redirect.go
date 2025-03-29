package support

import (
	"fmt"
	"net/http"

	cor "github.com/sh-yamaga/design-patterns-go/internal/patterns/CoR"
)

type RedirectionResponseSupport struct {
	HttpResponseSupport
}

func NewRedirectionResponseSupport() *RedirectionResponseSupport {
	rrs := &RedirectionResponseSupport{
		HttpResponseSupport{
			name:        "RedirectionResponseSupport",
			supportFrom: 300,
			supportTo:   399,
		},
	}
	rrs.ISupport = rrs

	return rrs
}

func (rrs *RedirectionResponseSupport) handle(hr *cor.HttpResponse) {
	fmt.Println("===", rrs.name, "===")
	fmt.Println(
		"Response Status Code:", hr.StatusCode,
		"(", http.StatusText(int(hr.StatusCode)), ")",
	)

	fmt.Println("handle Response...")
	fmt.Println("Done")
}
