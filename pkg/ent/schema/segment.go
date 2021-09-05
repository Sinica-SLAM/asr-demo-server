package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Segment holds the schema definition for the Segment entity.
type Segment struct {
	ent.Schema
}

// Fields of the Segment.
func (Segment) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique().Immutable(),
		field.Enum("type").Values("realtime", "upload", "post", "translate"),
		field.Float("start").Positive(),
		field.Float("length").Positive(),
	}
}

// Edges of the Segment.
func (Segment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("wordAlignments", WordAlignment.Type),
	}
}
