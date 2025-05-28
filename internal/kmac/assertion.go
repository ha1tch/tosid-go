package kmac

import (
	"errors"
	"fmt"
)

// Assertion represents a KMAC assertion
type Assertion struct {
	id               string
	subject          string
	relation         string
	object           string
	confidence       float64
	confidenceSource string
	properties       map[string]string
	negated          bool
}

// NewAssertion creates a new KMAC assertion
func NewAssertion(id string, subject string, relation string, object string) (*Assertion, error) {
	if id == "" {
		return nil, errors.New("assertion ID cannot be empty")
	}

	if !validateIdentifier(AssertionIDPrefix, id) {
		return nil, fmt.Errorf("invalid assertion ID format: %s", id)
	}

	if subject == "" || relation == "" || object == "" {
		return nil, errors.New("subject, relation, and object cannot be empty")
	}

	return &Assertion{
		id:         id,
		subject:    subject,
		relation:   relation,
		object:     object,
		confidence: 1.0, // Default to full confidence
		properties: make(map[string]string),
		negated:    false,
	}, nil
}

// ID returns the assertion's identifier
func (a *Assertion) ID() string {
	return a.id
}

// Type returns the statement type
func (a *Assertion) Type() string {
	return "ASSERT"
}

// Subject returns the assertion's subject
func (a *Assertion) Subject() string {
	return a.subject
}

// Relation returns the assertion's relation
func (a *Assertion) Relation() string {
	return a.relation
}

// Object returns the assertion's object
func (a *Assertion) Object() string {
	return a.object
}

// SetConfidence sets the confidence level and source for this assertion
func (a *Assertion) SetConfidence(level float64, source string) {
	if level < 0.0 {
		level = 0.0
	} else if level > 1.0 {
		level = 1.0
	}
	a.confidence = level
	a.confidenceSource = source
}

// GetConfidence returns the confidence level and source for this assertion
func (a *Assertion) GetConfidence() (float64, string) {
	return a.confidence, a.confidenceSource
}

// SetNegated sets whether this assertion is negated
func (a *Assertion) SetNegated(negated bool) {
	a.negated = negated
}

// IsNegated returns whether this assertion is negated
func (a *Assertion) IsNegated() bool {
	return a.negated
}

// SetProperty sets a property on the assertion
func (a *Assertion) SetProperty(key, value string) {
	a.properties[key] = value
}

// GetProperty retrieves a property from the assertion
func (a *Assertion) GetProperty(key string) (string, bool) {
	val, ok := a.properties[key]
	return val, ok
}

// String returns a string representation of the assertion in KMAC format
func (a *Assertion) String() string {
	prefix := "ASSERT"
	if a.negated {
		prefix = "NEGATE"
	}
	return fmt.Sprintf("%s #%s subject=[#%s] relation=[#%s] object=[#%s]", 
		prefix, a.id, a.subject, a.relation, a.object)
}

// ConfidenceString returns the confidence statement for this assertion
func (a *Assertion) ConfidenceString() string {
	if a.confidenceSource == "" {
		return ""
	}
	return fmt.Sprintf("CONFIDENCE #%s level=[%.4f] source=[%s]", 
		a.id, a.confidence, a.confidenceSource)
}

// IsEquivalent checks if this assertion is logically equivalent to another
func (a *Assertion) IsEquivalent(other *Assertion) bool {
	return a.subject == other.subject &&
		a.relation == other.relation &&
		a.object == other.object &&
		a.negated == other.negated
}

// Conflicts checks if this assertion conflicts with another
func (a *Assertion) Conflicts(other *Assertion) bool {
	return a.subject == other.subject &&
		a.relation == other.relation &&
		a.object == other.object &&
		a.negated != other.negated
}