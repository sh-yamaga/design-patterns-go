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
	return &RedirectionResponseSupport{
		HttpResponseSupport{
			name:        "RedirectionResponseSupport",
			supportFrom: 300,
			supportTo:   399,
		},
	}
}

func (rrs *RedirectionResponseSupport) Resolve(hr *cor.HttpResponse) {
	if rrs.isAvailable(hr) {
		fmt.Println("===", rrs.name, "===")
		fmt.Println(
			"Response Status Code:", hr.StatusCode,
			"(", http.StatusText(int(hr.StatusCode)), ")",
		)

		fmt.Println("Logging...")
		fmt.Println("Redirection Handling")
		fmt.Println("Done")
	} else {
		rrs.next.Resolve(hr)
	}
}
