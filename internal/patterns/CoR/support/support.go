package support

import (
	"fmt"
	"net/http"

	cor "github.com/sh-yamaga/design-patterns-go/internal/patterns/CoR"
)

type Support interface {
	SetNext(Support) Support
	Resolve(*cor.HttpResponse)
}

type SupportHandler interface {
	do(*cor.HttpResponse)
}

type HttpResponseSupport struct {
	name        string
	supportFrom cor.HttpStatusCode
	supportTo   cor.HttpStatusCode
	next        Support
	handler     SupportHandler
}

func (hrs *HttpResponseSupport) SetNext(next Support) Support {
	hrs.next = next

	return next
}

func (hrs *HttpResponseSupport) isAvailable(hr *cor.HttpResponse) bool {
	return hrs.supportFrom <= hr.StatusCode && hr.StatusCode <= hrs.supportTo
}

func (hrs *HttpResponseSupport) Resolve(hr *cor.HttpResponse) {
	if hrs.isAvailable(hr) {
		fmt.Println("===", hrs.name, "===")
		fmt.Printf(
			"Response Status Code: %d (%s)\n",
			hr.StatusCode,
			http.StatusText(int(hr.StatusCode)),
		)
		hrs.handler.do(hr)
		return
	}

	if hrs.next != nil {
		hrs.next.Resolve(hr)
	}
}
