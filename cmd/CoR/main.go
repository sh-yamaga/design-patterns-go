package main

import (
	cor "github.com/yamaga-shu/design-patterns-go/internal/patterns/CoR"
	"github.com/yamaga-shu/design-patterns-go/internal/patterns/CoR/support"
)

func main() {
	// create sample Http Response
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

	// create Response Supporter
	entry := support.NewEntry()
	info := support.NewInformationalResponseSupport()
	success := support.NewSuccessResponseSupport()
	redirect := support.NewRedirectionResponseSupport()
	clientError := support.NewClientErrorResponseSupport()
	serverError := support.NewServerErrorResponseSupport()

	// create Chain of Responsibility
	entry.SetNext(info).
		SetNext(success).
		SetNext(redirect).
		SetNext(clientError).
		SetNext(serverError)

	// resolve Http Response
	for _, r := range res {
		entry.Resolve(&r)
	}
	// Output:
	// === InformationalResponseSupport ===
	// Response Status Code: 100 (Continue)
	// Recieved Informational Response
	// Do Something...
	// Done
	// === SuccessfulResponseSupport ===
	// Response Status Code: 201 (Created)
	// Recieved Success Response
	// Do Something...
	// Done
	// === RedirectionResponseSupport ===
	// Response Status Code: 301 (Moved Permanently)
	// Recieved Redirection Response
	// Do Something...
	// Done
	// === ClientErrorResponseSupport ===
	// Response Status Code: 404 (Not Found)
	// Recieved Client Error Response
	// Do Something...
	// Done
	// === ServerErrorResponseSupport ===
	// Response Status Code: 500 (Internal Server Error)
	// Recieved Server Error Response
	// Do Something...
	// Done
}
