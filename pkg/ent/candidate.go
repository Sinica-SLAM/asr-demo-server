// Code generated by entc, DO NOT EDIT.

package ent

import (
	"asr-demo-recognize/pkg/ent/candidate"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Candidate is the model entity for the Candidate schema.
type Candidate struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Confidence holds the value of the "confidence" field.
	Confidence float64 `json:"confidence,omitempty"`
	// Word holds the value of the "word" field.
	Word                      string `json:"word,omitempty"`
	word_alignment_candidates *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Candidate) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case candidate.FieldConfidence:
			values[i] = new(sql.NullFloat64)
		case candidate.FieldID:
			values[i] = new(sql.NullInt64)
		case candidate.FieldWord:
			values[i] = new(sql.NullString)
		case candidate.ForeignKeys[0]: // word_alignment_candidates
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Candidate", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Candidate fields.
func (c *Candidate) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case candidate.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case candidate.FieldConfidence:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field confidence", values[i])
			} else if value.Valid {
				c.Confidence = value.Float64
			}
		case candidate.FieldWord:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field word", values[i])
			} else if value.Valid {
				c.Word = value.String
			}
		case candidate.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field word_alignment_candidates", value)
			} else if value.Valid {
				c.word_alignment_candidates = new(int)
				*c.word_alignment_candidates = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Candidate.
// Note that you need to call Candidate.Unwrap() before calling this method if this Candidate
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Candidate) Update() *CandidateUpdateOne {
	return (&CandidateClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Candidate entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Candidate) Unwrap() *Candidate {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Candidate is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Candidate) String() string {
	var builder strings.Builder
	builder.WriteString("Candidate(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", confidence=")
	builder.WriteString(fmt.Sprintf("%v", c.Confidence))
	builder.WriteString(", word=")
	builder.WriteString(c.Word)
	builder.WriteByte(')')
	return builder.String()
}

// Candidates is a parsable slice of Candidate.
type Candidates []*Candidate

func (c Candidates) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
