package builder

import (
	"strings"
)

// ISQLBuilder is the builder interface
type ISQLBuilder interface {
	Select(columns ...string) ISQLBuilder
	From(table string) ISQLBuilder
	Where(condition string) ISQLBuilder
	Build() string
}

// SQLBuilder is a concrete builder
type SQLBuilder struct {
	columns   []string
	table     string
	condition string
}

// NewSQLBuilder creates a new instance of SQLBuilder
func NewSQLBuilder() *SQLBuilder {
	return &SQLBuilder{}
}

func (b *SQLBuilder) Select(columns ...string) *SQLBuilder {
	b.columns = columns
	return b
}

func (b *SQLBuilder) From(table string) *SQLBuilder {
	b.table = table
	return b
}

func (b *SQLBuilder) Where(condition string) *SQLBuilder {
	b.condition = condition
	return b
}

func (b *SQLBuilder) Build() string {
	var sb strings.Builder
	sb.WriteString("SELECT ")
	if len(b.columns) > 0 {
		sb.WriteString(strings.Join(b.columns, ", "))
	} else {
		sb.WriteString("*")
	}
	sb.WriteString(" FROM ")
	sb.WriteString(b.table)
	if b.condition != "" {
		sb.WriteString(" WHERE ")
		sb.WriteString(b.condition)
	}
	sb.WriteString(";")
	return sb.String()
}
