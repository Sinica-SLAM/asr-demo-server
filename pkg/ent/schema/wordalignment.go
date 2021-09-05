package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// WordAlignment holds the schema definition for the WordAlignment entity.
type WordAlignment struct {
	ent.Schema
}

// Fields of the WordAlignment.
func (WordAlignment) Fields() []ent.Field {
	return []ent.Field{
		field.Float("start").Positive(),
		field.Float("length").Positive(),
		field.String("word").NotEmpty(),
		field.String("token"),
		field.Float("confidence").Negative(),
	}
}

// Edges of the WordAlignment.
func (WordAlignment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("candidates", Candidate.Type),
	}
}
