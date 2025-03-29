package support

import (
	"fmt"

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
	fmt.Println("handle Response...")
	fmt.Println("Done")
}
