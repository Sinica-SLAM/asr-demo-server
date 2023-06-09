// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CandidatesColumns holds the columns for the "candidates" table.
	CandidatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "confidence", Type: field.TypeFloat64},
		{Name: "word", Type: field.TypeString},
		{Name: "word_alignment_candidates", Type: field.TypeInt, Nullable: true},
	}
	// CandidatesTable holds the schema information for the "candidates" table.
	CandidatesTable = &schema.Table{
		Name:       "candidates",
		Columns:    CandidatesColumns,
		PrimaryKey: []*schema.Column{CandidatesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "candidates_word_alignments_candidates",
				Columns:    []*schema.Column{CandidatesColumns[3]},
				RefColumns: []*schema.Column{WordAlignmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SegmentsColumns holds the columns for the "segments" table.
	SegmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"realtime", "upload", "post", "translate"}},
		{Name: "start", Type: field.TypeFloat64},
		{Name: "length", Type: field.TypeFloat64},
		{Name: "session_segments", Type: field.TypeString, Nullable: true},
	}
	// SegmentsTable holds the schema information for the "segments" table.
	SegmentsTable = &schema.Table{
		Name:       "segments",
		Columns:    SegmentsColumns,
		PrimaryKey: []*schema.Column{SegmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "segments_sessions_segments",
				Columns:    []*schema.Column{SegmentsColumns[4]},
				RefColumns: []*schema.Column{SessionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "audio", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
	}
	// WordAlignmentsColumns holds the columns for the "word_alignments" table.
	WordAlignmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "start", Type: field.TypeFloat64},
		{Name: "length", Type: field.TypeFloat64},
		{Name: "word", Type: field.TypeString},
		{Name: "token", Type: field.TypeString},
		{Name: "confidence", Type: field.TypeFloat64},
		{Name: "segment_word_alignments", Type: field.TypeString, Nullable: true},
	}
	// WordAlignmentsTable holds the schema information for the "word_alignments" table.
	WordAlignmentsTable = &schema.Table{
		Name:       "word_alignments",
		Columns:    WordAlignmentsColumns,
		PrimaryKey: []*schema.Column{WordAlignmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "word_alignments_segments_wordAlignments",
				Columns:    []*schema.Column{WordAlignmentsColumns[6]},
				RefColumns: []*schema.Column{SegmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CandidatesTable,
		SegmentsTable,
		SessionsTable,
		WordAlignmentsTable,
	}
)

func init() {
	CandidatesTable.ForeignKeys[0].RefTable = WordAlignmentsTable
	SegmentsTable.ForeignKeys[0].RefTable = SessionsTable
	WordAlignmentsTable.ForeignKeys[0].RefTable = SegmentsTable
}
