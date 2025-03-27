package decorator

import "strings"

type UpperCaseFilter struct {
	Wrapped IStream
}

func (ucf *UpperCaseFilter) Read() string {
	wrappedData := ucf.Wrapped.Read()

	return strings.ToUpper(wrappedData)
}
