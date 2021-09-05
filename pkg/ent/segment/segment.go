// Code generated by entc, DO NOT EDIT.

package segment

import (
	"fmt"
)

const (
	// Label holds the string label denoting the segment type in the database.
	Label = "segment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldStart holds the string denoting the start field in the database.
	FieldStart = "start"
	// FieldLength holds the string denoting the length field in the database.
	FieldLength = "length"
	// EdgeWordAlignments holds the string denoting the wordalignments edge name in mutations.
	EdgeWordAlignments = "wordAlignments"
	// Table holds the table name of the segment in the database.
	Table = "segments"
	// WordAlignmentsTable is the table that holds the wordAlignments relation/edge.
	WordAlignmentsTable = "word_alignments"
	// WordAlignmentsInverseTable is the table name for the WordAlignment entity.
	// It exists in this package in order to avoid circular dependency with the "wordalignment" package.
	WordAlignmentsInverseTable = "word_alignments"
	// WordAlignmentsColumn is the table column denoting the wordAlignments relation/edge.
	WordAlignmentsColumn = "segment_word_alignments"
)

// Columns holds all SQL columns for segment fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldStart,
	FieldLength,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "segments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"session_segments",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// StartValidator is a validator for the "start" field. It is called by the builders before save.
	StartValidator func(float64) error
	// LengthValidator is a validator for the "length" field. It is called by the builders before save.
	LengthValidator func(float64) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeRealtime  Type = "realtime"
	TypeUpload    Type = "upload"
	TypePost      Type = "post"
	TypeTranslate Type = "translate"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeRealtime, TypeUpload, TypePost, TypeTranslate:
		return nil
	default:
		return fmt.Errorf("segment: invalid enum value for type field: %q", _type)
	}
}
