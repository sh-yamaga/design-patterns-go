package support

import (
	"fmt"
	"net/http"

	cor "github.com/sh-yamaga/design-patterns-go/internal/patterns/CoR"
)

type SuccessfulResponseSupport struct {
	HttpResponseSupport
}

func NewSuccessResponseSupport() *SuccessfulResponseSupport {
	srs := &SuccessfulResponseSupport{
		HttpResponseSupport{
			name:        "SuccessfulResponseSupport",
			supportFrom: 200,
			supportTo:   299,
		},
	}
	srs.ISupport = srs
	return srs
}

func (srs *SuccessfulResponseSupport) handle(hr *cor.HttpResponse) {
	fmt.Println("===", srs.name, "===")
	fmt.Println(
		"Response Status Code:", hr.StatusCode,
		"(", http.StatusText(int(hr.StatusCode)), ")",
	)

	fmt.Println("handle Response...")
	fmt.Println("Done")
}
