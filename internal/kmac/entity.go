package kmac

import (
	"errors"
	"fmt"
)

// Entity represents a KMAC entity definition
type Entity struct {
	id         string
	label      string
	tosidType  string
	properties map[string]string
}

// NewEntity creates a new KMAC entity
func NewEntity(id string, label string, tosidType string) (*Entity, error) {
	if id == "" {
		return nil, errors.New("entity ID cannot be empty")
	}

	if !validateIdentifier(EntityIDPrefix, id) {
		return nil, fmt.Errorf("invalid entity ID format: %s", id)
	}

	return &Entity{
		id:         id,
		label:      label,
		tosidType:  tosidType,
		properties: make(map[string]string),
	}, nil
}

// ID returns the entity's identifier
func (e *Entity) ID() string {
	return e.id
}

// Type returns the statement type
func (e *Entity) Type() string {
	return "DEF_ENTITY"
}

// Label returns the entity's label
func (e *Entity) Label() string {
	return e.label
}

// TOSIDType returns the entity's TOSID type
func (e *Entity) TOSIDType() string {
	return e.tosidType
}

// SetProperty sets a property on the entity
func (e *Entity) SetProperty(key, value string) {
	e.properties[key] = value
}

// GetProperty retrieves a property from the entity
func (e *Entity) GetProperty(key string) (string, bool) {
	val, ok := e.properties[key]
	return val, ok
}

// GetAllProperties returns all properties
func (e *Entity) GetAllProperties() map[string]string {
	result := make(map[string]string)
	for k, v := range e.properties {
		result[k] = v
	}
	return result
}

// RemoveProperty removes a property from the entity
func (e *Entity) RemoveProperty(key string) {
	delete(e.properties, key)
}

// HasProperty checks if the entity has a specific property
func (e *Entity) HasProperty(key string) bool {
	_, exists := e.properties[key]
	return exists
}

// String returns a string representation of the entity in KMAC format
func (e *Entity) String() string {
	return fmt.Sprintf("DEF_ENTITY #%s [%s] type=[%s]", e.id, e.label, e.tosidType)
}

// PropertiesString returns a string representation of all properties
func (e *Entity) PropertiesString() string {
	if len(e.properties) == 0 {
		return ""
	}
	
	result := ""
	for key, value := range e.properties {
		if result != "" {
			result += "\n"
		}
		result += fmt.Sprintf("PROPERTY #%s [%s] value=[%s]", e.id, key, value)
	}
	return result
}