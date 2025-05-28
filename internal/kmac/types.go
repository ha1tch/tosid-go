package kmac

import "errors"

// StatementCreator is an interface for creating KMAC statements
type StatementCreator interface {
	// CreateEntity creates a new KMAC entity
	CreateEntity(id string, label string, tosidType string) (*Entity, error)
	
	// CreateRelation creates a new KMAC relation
	CreateRelation(id string, label string, relationType string) (*Relation, error)
	
	// CreateAssertion creates a new KMAC assertion
	CreateAssertion(id string, subject string, relation string, object string) (*Assertion, error)
	
	// CreateProperty creates a new KMAC property
	CreateProperty(id string, label string, propertyType string) (*Property, error)
}

// StatementProcessor is an interface for processing KMAC statements
type StatementProcessor interface {
	// ProcessStatement processes a KMAC statement
	ProcessStatement(statement Statement) error
	
	// FindStatements finds statements matching specified criteria
	FindStatements(criteria map[string]string) []Statement
}

// ConfidenceManager is an interface for managing confidence levels in KMAC
type ConfidenceManager interface {
	// SetConfidence sets the confidence level for a statement
	SetConfidence(statementID string, level float64, source string) error
	
	// GetConfidence retrieves the confidence level for a statement
	GetConfidence(statementID string) (float64, string, error)
}

// KMAC public interfaces
var (
	Creator     StatementCreator   = &statementCreator{}
	Processor   StatementProcessor = &statementProcessor{}
	Confidence  ConfidenceManager  = &confidenceManager{}
)

// Internal implementations
type statementCreator struct{}
type statementProcessor struct{}
type confidenceManager struct{}

func (c *statementCreator) CreateEntity(id string, label string, tosidType string) (*Entity, error) {
	return NewEntity(id, label, tosidType)
}

func (c *statementCreator) CreateRelation(id string, label string, relationType string) (*Relation, error) {
	return NewRelation(id, label, relationType)
}

func (c *statementCreator) CreateAssertion(id string, subject string, relation string, object string) (*Assertion, error) {
	return NewAssertion(id, subject, relation, object)
}

func (c *statementCreator) CreateProperty(id string, label string, propertyType string) (*Property, error) {
	return NewProperty(id, label, propertyType)
}

func (p *statementProcessor) ProcessStatement(statement Statement) error {
	return ValidateKMACStatement(statement)
}

func (p *statementProcessor) FindStatements(criteria map[string]string) []Statement {
	// This would be implemented based on specific requirements
	return []Statement{}
}

func (cm *confidenceManager) SetConfidence(statementID string, level float64, source string) error {
	// This would require a global store of statements
	return errors.New("not implemented - requires statement store")
}

func (cm *confidenceManager) GetConfidence(statementID string) (float64, string, error) {
	// This would require a global store of statements
	return 0.0, "", errors.New("not implemented - requires statement store")
}