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
	irs := &InformationalResponseSupport{
		HttpResponseSupport{
			name:        "InformationalResponseSupport",
			supportFrom: 100,
			supportTo:   199,
		},
	}
	irs.ISupport = irs

	return irs
}

func (irs *InformationalResponseSupport) handle(hr *cor.HttpResponse) {
	fmt.Println("===", irs.name, "===")
	fmt.Println(
		"Response Status Code:", hr.StatusCode,
		"(", http.StatusText(int(hr.StatusCode)), ")",
	)

	fmt.Println("handle Response...")
	fmt.Println("Done")
}
