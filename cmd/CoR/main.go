package main

import (
	cor "github.com/sh-yamaga/design-patterns-go/internal/patterns/CoR"
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/CoR/support"
)

func main() {
	res := []cor.HttpResponse{
		// 100-199
		{StatusCode: 100},
		// 200-299
		{StatusCode: 201},
		// 300-399
		{StatusCode: 301},
		// 400-499
		{StatusCode: 404},
		// 500-599
		{StatusCode: 500},
	}

	info := support.NewInformationalResponseSupport()
	success := support.NewSuccessResponseSupport()
	redirect := support.NewRedirectionResponseSupport()
	clientError := support.NewClientErrorResponseSupport()
	serverError := support.NewServerErrorResponseSupport()

	info.SetNext(success).
		SetNext(redirect).
		SetNext(clientError).
		SetNext(serverError)

	for _, r := range res {
		info.Resolve(&r)
	}
	// Output:
	// === InformationalResponseSupport ===
	// Response Status Code: 100 ( Continue )
	// handle Response...
	// Done
	// === SuccessfulResponseSupport ===
	// Response Status Code: 201 ( Created )
	// handle Response...
	// Done
	// === RedirectionResponseSupport ===
	// Response Status Code: 301 ( Moved Permanently )
	// handle Response...
	// Done
	// === ClientErrorResponseSupport ===
	// Response Status Code: 404 ( Not Found )
	// handle Response...
	// Done
	// === ServerErrorResponseSupport ===
	// Response Status Code: 500 ( Internal Server Error )
	// handle Response...
	// Done
}
