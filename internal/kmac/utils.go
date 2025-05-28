package kmac

import (
	"fmt"
	"sort"
	"strings"
)

// StatementCollection represents a collection of KMAC statements
type StatementCollection struct {
	statements map[string]Statement
}

// NewStatementCollection creates a new statement collection
func NewStatementCollection() *StatementCollection {
	return &StatementCollection{
		statements: make(map[string]Statement),
	}
}

// Add adds a statement to the collection
func (sc *StatementCollection) Add(statement Statement) error {
	if statement == nil {
		return fmt.Errorf("cannot add nil statement")
	}
	
	if err := ValidateKMACStatement(statement); err != nil {
		return fmt.Errorf("invalid statement: %v", err)
	}
	
	sc.statements[statement.ID()] = statement
	return nil
}

// Get retrieves a statement by ID
func (sc *StatementCollection) Get(id string) (Statement, bool) {
	stmt, exists := sc.statements[id]
	return stmt, exists
}

// Remove removes a statement by ID
func (sc *StatementCollection) Remove(id string) bool {
	if _, exists := sc.statements[id]; exists {
		delete(sc.statements, id)
		return true
	}
	return false
}

// GetAll returns all statements
func (sc *StatementCollection) GetAll() []Statement {
	statements := make([]Statement, 0, len(sc.statements))
	for _, stmt := range sc.statements {
		statements = append(statements, stmt)
	}
	return statements
}

// GetByType returns all statements of a specific type
func (sc *StatementCollection) GetByType(statementType string) []Statement {
	var statements []Statement
	for _, stmt := range sc.statements {
		if stmt.Type() == statementType {
			statements = append(statements, stmt)
		}
	}
	return statements
}

// Count returns the number of statements
func (sc *StatementCollection) Count() int {
	return len(sc.statements)
}

// Clear removes all statements
func (sc *StatementCollection) Clear() {
	sc.statements = make(map[string]Statement)
}

// FilterByPrefix returns statements whose IDs start with the given prefix
func (sc *StatementCollection) FilterByPrefix(prefix string) []Statement {
	var statements []Statement
	for id, stmt := range sc.statements {
		if strings.HasPrefix(id, prefix) {
			statements = append(statements, stmt)
		}
	}
	return statements
}

// GetStatistics returns statistics about the collection
func (sc *StatementCollection) GetStatistics() map[string]int {
	stats := make(map[string]int)
	
	for _, stmt := range sc.statements {
		key := "type_" + stmt.Type()
		stats[key]++
	}
	
	stats["total"] = len(sc.statements)
	return stats
}

// ExportToStrings converts all statements to their string representations
func (sc *StatementCollection) ExportToStrings() []string {
	var strings []string
	
	// Get all statements and sort by ID for consistent output
	ids := make([]string, 0, len(sc.statements))
	for id := range sc.statements {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	
	for _, id := range ids {
		strings = append(strings, sc.statements[id].String())
	}
	
	return strings
}

// Validate checks all statements for consistency
func (sc *StatementCollection) Validate() []string {
	var warnings []string
	
	// Check each statement individually
	for id, stmt := range sc.statements {
		if err := ValidateKMACStatement(stmt); err != nil {
			warnings = append(warnings, fmt.Sprintf("Statement %s: %v", id, err))
		}
	}
	
	// Check for reference consistency
	entityIDs := make(map[string]bool)
	relationIDs := make(map[string]bool)
	
	// Collect all entity and relation IDs
	for _, stmt := range sc.statements {
		switch s := stmt.(type) {
		case *Entity:
			entityIDs[s.ID()] = true
		case *Relation:
			relationIDs[s.ID()] = true
		}
	}
	
	// Check assertions for valid references
	for _, stmt := range sc.statements {
		if assertion, ok := stmt.(*Assertion); ok {
			if !entityIDs[assertion.Subject()] {
				warnings = append(warnings, fmt.Sprintf("Assertion %s references unknown subject %s", assertion.ID(), assertion.Subject()))
			}
			if !entityIDs[assertion.Object()] {
				warnings = append(warnings, fmt.Sprintf("Assertion %s references unknown object %s", assertion.ID(), assertion.Object()))
			}
			if !relationIDs[assertion.Relation()] {
				// Check if it's a built-in relation
				builtInRelations := []string{"AGENT", "LOCATION", "OCCURRED_AT", "INSTANCE_OF"}
				isBuiltIn := false
				for _, builtin := range builtInRelations {
					if assertion.Relation() == builtin {
						isBuiltIn = true
						break
					}
				}
				if !isBuiltIn {
					warnings = append(warnings, fmt.Sprintf("Assertion %s references unknown relation %s", assertion.ID(), assertion.Relation()))
				}
			}
		}
	}
	
	return warnings
}

// KMACBuilder helps build complex KMAC structures
type KMACBuilder struct {
	collection *StatementCollection
	entityCounter    int
	relationCounter  int
	assertionCounter int
}

// NewKMACBuilder creates a new KMAC builder
func NewKMACBuilder() *KMACBuilder {
	return &KMACBuilder{
		collection: NewStatementCollection(),
		entityCounter: 1,
		relationCounter: 1,
		assertionCounter: 1,
	}
}

// AddEntity adds an entity with auto-generated ID
func (kb *KMACBuilder) AddEntity(label string, tosidType string) (*Entity, error) {
	id := fmt.Sprintf("E%04d", kb.entityCounter)
	entity, err := NewEntity(id, label, tosidType)
	if err != nil {
		return nil, err
	}
	
	if err := kb.collection.Add(entity); err != nil {
		return nil, err
	}
	
	kb.entityCounter++
	return entity, nil
}

// AddRelation adds a relation with auto-generated ID
func (kb *KMACBuilder) AddRelation(label string, relationType string) (*Relation, error) {
	id := fmt.Sprintf("R%04d", kb.relationCounter)
	relation, err := NewRelation(id, label, relationType)
	if err != nil {
		return nil, err
	}
	
	if err := kb.collection.Add(relation); err != nil {
		return nil, err
	}
	
	kb.relationCounter++
	return relation, nil
}

// AddAssertion adds an assertion with auto-generated ID
func (kb *KMACBuilder) AddAssertion(subject string, relation string, object string) (*Assertion, error) {
	id := fmt.Sprintf("F%04d", kb.assertionCounter)
	assertion, err := NewAssertion(id, subject, relation, object)
	if err != nil {
		return nil, err
	}
	
	if err := kb.collection.Add(assertion); err != nil {
		return nil, err
	}
	
	kb.assertionCounter++
	return assertion, nil
}

// GetCollection returns the statement collection
func (kb *KMACBuilder) GetCollection() *StatementCollection {
	return kb.collection
}

// Build returns all statements as a slice
func (kb *KMACBuilder) Build() []Statement {
	return kb.collection.GetAll()
}

// Reset clears the builder
func (kb *KMACBuilder) Reset() {
	kb.collection = NewStatementCollection()
	kb.entityCounter = 1
	kb.relationCounter = 1
	kb.assertionCounter = 1
}

// Validate validates the built structure
func (kb *KMACBuilder) Validate() []string {
	return kb.collection.Validate()
}