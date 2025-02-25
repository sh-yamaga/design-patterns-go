package factory

import (
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/template/processor"
)

// DataProcessorFactory defines the factory method for creating IDataProcessor instances
type DataProcessorFactory interface {
	NewCsvProcessor() *processor.CsvProcessor
	NewXmlProcessor() *processor.XmlProcessor
}

// ProcessorFactory is a concrete factory for creating CsvProcessor instances
type ProcessorFactory struct{}

// NewCsvProcessor creates a new CsvProcessor instance
func (pf *ProcessorFactory) NewCsvProcessor() *processor.CsvProcessor {
	cp := &processor.CsvProcessor{}
	cp.BaseProcessor.IDataProcessor = cp

	return cp
}

// NewXmlProcessor creates a new XmlProcessor instance
func (pf *ProcessorFactory) NewXmlProcessor() *processor.XmlProcessor {
	xp := &processor.XmlProcessor{}
	xp.BaseProcessor.IDataProcessor = xp

	return xp
}
