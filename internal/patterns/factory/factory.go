package factory

import "github.com/sh-yamaga/design-patterns-go/internal/patterns/template"

// DataProcessorFactory defines the factory method for creating IDataProcessor instances
type DataProcessorFactory interface {
	CreateCsvProcessor() *template.CsvProcessor
}

// ProcessorFactory is a concrete factory for creating CsvProcessor instances
type ProcessorFactory struct{}

// CreateCsvProcessor creates a new CsvProcessor instance
func (f *ProcessorFactory) CreateCsvProcessor() *template.CsvProcessor {
	cp := &template.CsvProcessor{}
	cp.BaseProcessor.IDataProcessor = cp

	return cp
}
