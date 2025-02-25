package factory

import (
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/template/processor"
)

// DataProcessorFactory defines the factory method for creating IDataProcessor instances
type DataProcessorFactory interface {
	NewCsvProcessor() *processor.IDataProcessor
	NewXmlProcessor() *processor.IDataProcessor
}

// ProcessorFactory is a concrete factory for creating CsvProcessor instances
type ProcessorFactory struct{}

// NewCsvProcessor creates a new CsvProcessor instance, implements IDataProcessor
func (pf *ProcessorFactory) NewCsvProcessor() processor.IDataProcessor {
	cp := &processor.CsvProcessor{}
	cp.BaseProcessor.IDataProcessor = cp

	return cp
}

// NewXmlProcessor creates a new XmlProcessor instance, implements IDataProcessor
func (pf *ProcessorFactory) NewXmlProcessor() processor.IDataProcessor {
	xp := &processor.XmlProcessor{}
	xp.BaseProcessor.IDataProcessor = xp

	return xp
}
