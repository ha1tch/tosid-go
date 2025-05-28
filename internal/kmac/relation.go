package kmac

import (
	"errors"
	"fmt"
)

// Relation represents a KMAC relation definition
type Relation struct {
	id           string
	label        string
	relationType string
	properties   map[string]string
	domain       string // Subject domain constraint
	range_       string // Object domain constraint
}

// NewRelation creates a new KMAC relation
func NewRelation(id string, label string, relationType string) (*Relation, error) {
	if id == "" {
		return nil, errors.New("relation ID cannot be empty")
	}

	if !validateIdentifier(RelationIDPrefix, id) {
		return nil, fmt.Errorf("invalid relation ID format: %s", id)
	}

	return &Relation{
		id:           id,
		label:        label,
		relationType: relationType,
		properties:   make(map[string]string),
	}, nil
}

// ID returns the relation's identifier
func (r *Relation) ID() string {
	return r.id
}

// Type returns the statement type
func (r *Relation) Type() string {
	return "DEF_RELATION"
}

// Label returns the relation's label
func (r *Relation) Label() string {
	return r.label
}

// RelationType returns the relation's type
func (r *Relation) RelationType() string {
	return r.relationType
}

// SetDomain sets the domain constraint for this relation
func (r *Relation) SetDomain(domain string) {
	r.domain = domain
}

// GetDomain returns the domain constraint
func (r *Relation) GetDomain() string {
	return r.domain
}

// SetRange sets the range constraint for this relation
func (r *Relation) SetRange(range_ string) {
	r.range_ = range_
}

// GetRange returns the range constraint
func (r *Relation) GetRange() string {
	return r.range_
}

// SetProperty sets a property on the relation
func (r *Relation) SetProperty(key, value string) {
	r.properties[key] = value
}

// GetProperty retrieves a property from the relation
func (r *Relation) GetProperty(key string) (string, bool) {
	val, ok := r.properties[key]
	return val, ok
}

// IsSymmetric checks if this relation is symmetric
func (r *Relation) IsSymmetric() bool {
	symmetric, exists := r.properties["symmetric"]
	return exists && symmetric == "true"
}

// IsTransitive checks if this relation is transitive
func (r *Relation) IsTransitive() bool {
	transitive, exists := r.properties["transitive"]
	return exists && transitive == "true"
}

// IsReflexive checks if this relation is reflexive
func (r *Relation) IsReflexive() bool {
	reflexive, exists := r.properties["reflexive"]
	return exists && reflexive == "true"
}

// String returns a string representation of the relation in KMAC format
func (r *Relation) String() string {
	return fmt.Sprintf("DEF_RELATION #%s [%s] type=[%s]", r.id, r.label, r.relationType)
}