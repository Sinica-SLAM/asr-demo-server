// Code generated by entc, DO NOT EDIT.

package candidate

const (
	// Label holds the string label denoting the candidate type in the database.
	Label = "candidate"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldConfidence holds the string denoting the confidence field in the database.
	FieldConfidence = "confidence"
	// FieldWord holds the string denoting the word field in the database.
	FieldWord = "word"
	// Table holds the table name of the candidate in the database.
	Table = "candidates"
)

// Columns holds all SQL columns for candidate fields.
var Columns = []string{
	FieldID,
	FieldConfidence,
	FieldWord,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "candidates"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"word_alignment_candidates",
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
	// ConfidenceValidator is a validator for the "confidence" field. It is called by the builders before save.
	ConfidenceValidator func(float64) error
	// WordValidator is a validator for the "word" field. It is called by the builders before save.
	WordValidator func(string) error
)
