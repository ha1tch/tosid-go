package kmac

import (
	"errors"
	"fmt"
)

// Property represents a KMAC property definition
type Property struct {
	id           string
	label        string
	propertyType string
	domain       string // What entities can have this property
	range_       string // What values this property can take
	functional   bool   // Whether this property is functional (single-valued)
}

// NewProperty creates a new KMAC property
func NewProperty(id string, label string, propertyType string) (*Property, error) {
	if id == "" {
		return nil, errors.New("property ID cannot be empty")
	}

	if !validateIdentifier(PropertyIDPrefix, id) {
		return nil, fmt.Errorf("invalid property ID format: %s", id)
	}

	return &Property{
		id:           id,
		label:        label,
		propertyType: propertyType,
		functional:   false,
	}, nil
}

// ID returns the property's identifier
func (p *Property) ID() string {
	return p.id
}

// Type returns the statement type
func (p *Property) Type() string {
	return "DEF_PROPERTY"
}

// Label returns the property's label
func (p *Property) Label() string {
	return p.label
}

// PropertyType returns the property's type
func (p *Property) PropertyType() string {
	return p.propertyType
}

// SetDomain sets the domain for this property
func (p *Property) SetDomain(domain string) {
	p.domain = domain
}

// GetDomain returns the domain
func (p *Property) GetDomain() string {
	return p.domain
}

// SetRange sets the range for this property
func (p *Property) SetRange(range_ string) {
	p.range_ = range_
}

// GetRange returns the range
func (p *Property) GetRange() string {
	return p.range_
}

// SetFunctional sets whether this property is functional
func (p *Property) SetFunctional(functional bool) {
	p.functional = functional
}

// IsFunctional returns whether this property is functional
func (p *Property) IsFunctional() bool {
	return p.functional
}

// String returns a string representation of the property in KMAC format
func (p *Property) String() string {
	return fmt.Sprintf("DEF_PROPERTY #%s [%s] type=[%s]", p.id, p.label, p.propertyType)
}

// PropertyAssertion represents a property assertion about an entity
type PropertyAssertion struct {
	id         string
	entity     string
	property   string
	value      string
	confidence float64
	source     string
}

// NewPropertyAssertion creates a new property assertion
func NewPropertyAssertion(id string, entity string, property string, value string) (*PropertyAssertion, error) {
	if id == "" || entity == "" || property == "" || value == "" {
		return nil, errors.New("all fields are required for property assertion")
	}

	return &PropertyAssertion{
		id:         id,
		entity:     entity,
		property:   property,
		value:      value,
		confidence: 1.0,
	}, nil
}

// ID returns the property assertion's identifier
func (pa *PropertyAssertion) ID() string {
	return pa.id
}

// Type returns the statement type
func (pa *PropertyAssertion) Type() string {
	return "PROPERTY_ASSERT"
}

// Entity returns the entity ID
func (pa *PropertyAssertion) Entity() string {
	return pa.entity
}

// Property returns the property ID
func (pa *PropertyAssertion) Property() string {
	return pa.property
}

// Value returns the property value
func (pa *PropertyAssertion) Value() string {
	return pa.value
}

// SetConfidence sets the confidence for this property assertion
func (pa *PropertyAssertion) SetConfidence(level float64, source string) {
	if level < 0.0 {
		level = 0.0
	} else if level > 1.0 {
		level = 1.0
	}
	pa.confidence = level
	pa.source = source
}

// GetConfidence returns the confidence and source
func (pa *PropertyAssertion) GetConfidence() (float64, string) {
	return pa.confidence, pa.source
}

// String returns a string representation of the property assertion
func (pa *PropertyAssertion) String() string {
	return fmt.Sprintf("ASSERT #%s subject=[#%s] property=[#%s] value=[%s]", 
		pa.id, pa.entity, pa.property, pa.value)
}