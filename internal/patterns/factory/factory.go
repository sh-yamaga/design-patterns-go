package factory

import (
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/template/processor"
)

// DataProcessorFactory defines the factory method for creating IDataProcessor instances
type DataProcessorFactory interface {
	CreateCsvProcessor() *processor.CsvProcessor
}

// ProcessorFactory is a concrete factory for creating CsvProcessor instances
type ProcessorFactory struct{}

// CreateCsvProcessor creates a new CsvProcessor instance
func (f *ProcessorFactory) CreateCsvProcessor() *processor.CsvProcessor {
	cp := &processor.CsvProcessor{}
	cp.BaseProcessor.IDataProcessor = cp

	return cp
}
