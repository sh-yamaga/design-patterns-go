package support

import (
	"fmt"
	"net/http"

	cor "github.com/sh-yamaga/design-patterns-go/internal/patterns/CoR"
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
		},
	}
}

func (irs *InformationalResponseSupport) Resolve(hr *cor.HttpResponse) {
	if irs.isAvailable(hr) {
		fmt.Println("===", irs.name, "===")
		fmt.Println(
			"Response Status Code:", hr.StatusCode,
			"(", http.StatusText(int(hr.StatusCode)), ")",
		)

		fmt.Println("Logging...")
		fmt.Println("Done")
	} else {
		irs.next.Resolve(hr)
	}
}
