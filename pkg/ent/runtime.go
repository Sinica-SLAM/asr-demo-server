// Code generated by entc, DO NOT EDIT.

package ent

import (
	"asr-demo-server/pkg/ent/candidate"
	"asr-demo-server/pkg/ent/schema"
	"asr-demo-server/pkg/ent/segment"
	"asr-demo-server/pkg/ent/session"
	"asr-demo-server/pkg/ent/wordalignment"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	candidateFields := schema.Candidate{}.Fields()
	_ = candidateFields
	// candidateDescConfidence is the schema descriptor for confidence field.
	candidateDescConfidence := candidateFields[0].Descriptor()
	// candidate.ConfidenceValidator is a validator for the "confidence" field. It is called by the builders before save.
	candidate.ConfidenceValidator = candidateDescConfidence.Validators[0].(func(float64) error)
	// candidateDescWord is the schema descriptor for word field.
	candidateDescWord := candidateFields[1].Descriptor()
	// candidate.WordValidator is a validator for the "word" field. It is called by the builders before save.
	candidate.WordValidator = candidateDescWord.Validators[0].(func(string) error)
	segmentFields := schema.Segment{}.Fields()
	_ = segmentFields
	// segmentDescStart is the schema descriptor for start field.
	segmentDescStart := segmentFields[2].Descriptor()
	// segment.StartValidator is a validator for the "start" field. It is called by the builders before save.
	segment.StartValidator = segmentDescStart.Validators[0].(func(float64) error)
	// segmentDescLength is the schema descriptor for length field.
	segmentDescLength := segmentFields[3].Descriptor()
	// segment.LengthValidator is a validator for the "length" field. It is called by the builders before save.
	segment.LengthValidator = segmentDescLength.Validators[0].(func(float64) error)
	// segmentDescID is the schema descriptor for id field.
	segmentDescID := segmentFields[0].Descriptor()
	// segment.IDValidator is a validator for the "id" field. It is called by the builders before save.
	segment.IDValidator = segmentDescID.Validators[0].(func(string) error)
	sessionFields := schema.Session{}.Fields()
	_ = sessionFields
	// sessionDescAudio is the schema descriptor for audio field.
	sessionDescAudio := sessionFields[1].Descriptor()
	// session.AudioValidator is a validator for the "audio" field. It is called by the builders before save.
	session.AudioValidator = sessionDescAudio.Validators[0].(func(string) error)
	// sessionDescCreatedAt is the schema descriptor for created_at field.
	sessionDescCreatedAt := sessionFields[2].Descriptor()
	// session.DefaultCreatedAt holds the default value on creation for the created_at field.
	session.DefaultCreatedAt = sessionDescCreatedAt.Default.(func() time.Time)
	// sessionDescUpdatedAt is the schema descriptor for updated_at field.
	sessionDescUpdatedAt := sessionFields[3].Descriptor()
	// session.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	session.DefaultUpdatedAt = sessionDescUpdatedAt.Default.(func() time.Time)
	// session.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	session.UpdateDefaultUpdatedAt = sessionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sessionDescID is the schema descriptor for id field.
	sessionDescID := sessionFields[0].Descriptor()
	// session.IDValidator is a validator for the "id" field. It is called by the builders before save.
	session.IDValidator = sessionDescID.Validators[0].(func(string) error)
	wordalignmentFields := schema.WordAlignment{}.Fields()
	_ = wordalignmentFields
	// wordalignmentDescStart is the schema descriptor for start field.
	wordalignmentDescStart := wordalignmentFields[0].Descriptor()
	// wordalignment.StartValidator is a validator for the "start" field. It is called by the builders before save.
	wordalignment.StartValidator = wordalignmentDescStart.Validators[0].(func(float64) error)
	// wordalignmentDescLength is the schema descriptor for length field.
	wordalignmentDescLength := wordalignmentFields[1].Descriptor()
	// wordalignment.LengthValidator is a validator for the "length" field. It is called by the builders before save.
	wordalignment.LengthValidator = wordalignmentDescLength.Validators[0].(func(float64) error)
	// wordalignmentDescWord is the schema descriptor for word field.
	wordalignmentDescWord := wordalignmentFields[2].Descriptor()
	// wordalignment.WordValidator is a validator for the "word" field. It is called by the builders before save.
	wordalignment.WordValidator = wordalignmentDescWord.Validators[0].(func(string) error)
	// wordalignmentDescConfidence is the schema descriptor for confidence field.
	wordalignmentDescConfidence := wordalignmentFields[4].Descriptor()
	// wordalignment.ConfidenceValidator is a validator for the "confidence" field. It is called by the builders before save.
	wordalignment.ConfidenceValidator = wordalignmentDescConfidence.Validators[0].(func(float64) error)
}
