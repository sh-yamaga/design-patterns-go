package factory

import (
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/template/processor"
)

// DataProcessorFactory defines the factory method for creating IDataProcessor instances
type DataProcessorFactory interface {
	CsvProcessor() *processor.CsvProcessor
	XmlProcessor() *processor.XmlProcessor
}

// ProcessorFactory is a concrete factory for creating CsvProcessor instances
type ProcessorFactory struct{}

// CsvProcessor creates a new CsvProcessor instance
func (pf *ProcessorFactory) CsvProcessor() *processor.CsvProcessor {
	cp := &processor.CsvProcessor{}
	cp.BaseProcessor.IDataProcessor = cp

	return cp
}

// XmlProcessor creates a new XmlProcessor instance
func (pf *ProcessorFactory) XmlProcessor() *processor.XmlProcessor {
	xp := &processor.XmlProcessor{}
	xp.BaseProcessor.IDataProcessor = xp

	return xp
}
