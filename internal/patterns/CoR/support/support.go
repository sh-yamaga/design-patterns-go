package support

import (
	"fmt"
	"net/http"

	cor "github.com/sh-yamaga/design-patterns-go/internal/patterns/CoR"
)

type ISupport interface {
	SetNext(ISupport) ISupport
	Resolve(*cor.HttpResponse)
	isAvailable(*cor.HttpResponse) bool
	handle(*cor.HttpResponse)
}

type HttpResponseSupport struct {
	name        string
	supportFrom cor.HttpStatusCode
	supportTo   cor.HttpStatusCode
	next        ISupport
	ISupport
}

func (hrs *HttpResponseSupport) SetNext(next ISupport) ISupport {
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
			"Response Status Code: %d, (%s)\n",
			hr.StatusCode,
			http.StatusText(int(hr.StatusCode)),
		)
		hrs.ISupport.handle(hr)
		return
	}

	if hrs.next != nil {
		hrs.next.Resolve(hr)
	}
}

func (hrs *HttpResponseSupport) handle(hr *cor.HttpResponse) {
	// will override on embed struct
}
