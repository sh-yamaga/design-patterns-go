package support

import cor "github.com/sh-yamaga/design-patterns-go/internal/patterns/CoR"

type ISupport interface {
	SetNext(ISupport) ISupport
	isAvailable(*cor.HttpResponse) bool
	Resolve(*cor.HttpResponse)
}

type HttpResponseSupport struct {
	name        string
	supportFrom cor.HttpStatusCode
	supportTo   cor.HttpStatusCode
	next        ISupport
}

func (hrs *HttpResponseSupport) SetNext(next ISupport) ISupport {
	hrs.next = next

	return next
}

func (hrs *HttpResponseSupport) isAvailable(hr *cor.HttpResponse) bool {
	return hrs.supportFrom <= hr.StatusCode && hr.StatusCode <= hrs.supportTo
}
