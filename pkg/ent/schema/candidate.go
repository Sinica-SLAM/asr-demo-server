package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Candidate holds the schema definition for the Candidate entity.
type Candidate struct {
	ent.Schema
}

// Fields of the Candidate.
func (Candidate) Fields() []ent.Field {
	return []ent.Field{
		field.Float("confidence").Positive(),
		field.String("word").NotEmpty(),
	}
}

// Edges of the Candidate.
func (Candidate) Edges() []ent.Edge {
	return nil
}
