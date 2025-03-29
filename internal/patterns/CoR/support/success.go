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
	return &SuccessfulResponseSupport{
		HttpResponseSupport{
			name:        "SuccessfulResponseSupport",
			supportFrom: 200,
			supportTo:   299,
		},
	}
}

func (srs *SuccessfulResponseSupport) Resolve(hr *cor.HttpResponse) {
	if srs.isAvailable(hr) {
		fmt.Println("===", srs.name, "===")
		fmt.Println(
			"Response Status Code:", hr.StatusCode,
			"(", http.StatusText(int(hr.StatusCode)), ")",
		)

		fmt.Println("Logging...")
		fmt.Println("handle Response...")
		fmt.Println("Done")
	} else {
		srs.next.Resolve(hr)
	}
}
