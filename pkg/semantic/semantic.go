package semantic

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ha1tch/tosid-go/pkg/kmac"
	"github.com/ha1tch/tosid-go/pkg/tosid"
)

// EntityReference represents a reference to an entity with both KMAC and TOSID information
type EntityReference struct {
	KMACEntity *kmac.Entity
	TOSIDObj   *tosid.TOSID
}

// SemanticStore represents a store for semantic entities and relationships
type SemanticStore struct {
	entities    map[string]*EntityReference
	relations   map[string]*kmac.Relation
	assertions  map[string]*kmac.Assertion
	properties  map[string]*kmac.Property
}

// NewSemanticStore creates a new semantic store
func NewSemanticStore() *SemanticStore {
	return &SemanticStore{
		entities:   make(map[string]*EntityReference),
		relations:  make(map[string]*kmac.Relation),
		assertions: make(map[string]*kmac.Assertion),
		properties: make(map[string]*kmac.Property),
	}
}

// AddEntity adds a new entity to the store
func (s *SemanticStore) AddEntity(id string, label string, tosidCode string) error {
	// Create KMAC entity
	entity, err := kmac.NewEntity(id, label, tosidCode)
	if err != nil {
		return fmt.Errorf("failed to create KMAC entity: %v", err)
	}

	// Parse TOSID code if provided
	var tosidObj *tosid.TOSID
	if tosidCode != "" {
		tosidObj, err = tosid.Parse(tosidCode)
		if err != nil {
			return fmt.Errorf("failed to parse TOSID code: %v", err)
		}
	}

	// Create entity reference
	entityRef := &EntityReference{
		KMACEntity: entity,
		TOSIDObj:   tosidObj,
	}

	s.entities[id] = entityRef
	return nil
}

// GetEntity retrieves an entity from the store
func (s *SemanticStore) GetEntity(id string) (*EntityReference, error) {
	entity, exists := s.entities[id]
	if !exists {
		return nil, fmt.Errorf("entity %s not found", id)
	}
	return entity, nil
}

// AddRelation adds a new relation to the store
func (s *SemanticStore) AddRelation(id string, label string, relationType string) error {
	relation, err := kmac.NewRelation(id, label, relationType)
	if err != nil {
		return fmt.Errorf("failed to create relation: %v", err)
	}

	s.relations[id] = relation
	return nil
}

// GetRelation retrieves a relation from the store
func (s *SemanticStore) GetRelation(id string) (*kmac.Relation, error) {
	relation, exists := s.relations[id]
	if !exists {
		return nil, fmt.Errorf("relation %s not found", id)
	}
	return relation, nil
}

// CreateAssertion creates a new assertion between entities
func (s *SemanticStore) CreateAssertion(id string, subjectID string, relationID string, objectID string) error {
	// Verify that subject and object entities exist
	if _, err := s.GetEntity(subjectID); err != nil {
		return fmt.Errorf("subject entity not found: %v", err)
	}

	if _, err := s.GetEntity(objectID); err != nil {
		return fmt.Errorf("object entity not found: %v", err)
	}

	// Create assertion
	assertion, err := kmac.NewAssertion(id, subjectID, relationID, objectID)
	if err != nil {
		return fmt.Errorf("failed to create assertion: %v", err)
	}

	s.assertions[id] = assertion
	return nil
}

// GetAssertion retrieves an assertion from the store
func (s *SemanticStore) GetAssertion(id string) (*kmac.Assertion, error) {
	assertion, exists := s.assertions[id]
	if !exists {
		return nil, fmt.Errorf("assertion %s not found", id)
	}
	return assertion, nil
}

// FindEntitiesByTOSIDPattern finds entities matching a TOSID pattern
func (s *SemanticStore) FindEntitiesByTOSIDPattern(pattern string) []*EntityReference {
	var results []*EntityReference

	for _, entityRef := range s.entities {
		if entityRef.TOSIDObj != nil && entityRef.TOSIDObj.MatchesPattern(pattern) {
			results = append(results, entityRef)
		}
	}

	return results
}

// FindAssertionsForEntity finds all assertions where the given entity is either subject or object
func (s *SemanticStore) FindAssertionsForEntity(entityID string) []*kmac.Assertion {
	var results []*kmac.Assertion

	for _, assertion := range s.assertions {
		if assertion.Subject() == entityID || assertion.Object() == entityID {
			results = append(results, assertion)
		}
	}

	return results
}

// FindEntitiesByLabel finds entities by label (case-insensitive partial match)
func (s *SemanticStore) FindEntitiesByLabel(labelPattern string) []*EntityReference {
	var results []*EntityReference
	pattern := strings.ToLower(labelPattern)

	for _, entityRef := range s.entities {
		entityLabel := strings.ToLower(entityRef.KMACEntity.Label())
		if strings.Contains(entityLabel, pattern) {
			results = append(results, entityRef)
		}
	}

	return results
}

// FindRelatedEntities finds entities related to a given entity through assertions
func (s *SemanticStore) FindRelatedEntities(entityID string) map[string][]*EntityReference {
	results := make(map[string][]*EntityReference)

	for _, assertion := range s.assertions {
		var relatedID string
		var direction string

		if assertion.Subject() == entityID {
			relatedID = assertion.Object()
			direction = "outbound"
		} else if assertion.Object() == entityID {
			relatedID = assertion.Subject()
			direction = "inbound"
		}

		if relatedID != "" {
			if relatedEntity, exists := s.entities[relatedID]; exists {
				if results[direction] == nil {
					results[direction] = make([]*EntityReference, 0)
				}
				results[direction] = append(results[direction], relatedEntity)
			}
		}
	}

	return results
}

// GetStatistics returns statistics about the semantic store
func (s *SemanticStore) GetStatistics() map[string]int {
	stats := make(map[string]int)
	stats["entities"] = len(s.entities)
	stats["relations"] = len(s.relations)
	stats["assertions"] = len(s.assertions)
	stats["properties"] = len(s.properties)

	// Count entities by taxonomy
	taxonomyCount := make(map[string]int)
	for _, entityRef := range s.entities {
		if entityRef.TOSIDObj != nil {
			taxonomy := entityRef.TOSIDObj.TaxonomyCode
			taxonomyCount[taxonomy]++
		}
	}

	for taxonomy, count := range taxonomyCount {
		stats["taxonomy_"+taxonomy] = count
	}

	return stats
}

// ValidateStore performs consistency checks on the semantic store
func (s *SemanticStore) ValidateStore() []string {
	var warnings []string

	// Check for assertions with missing entities
	for assertionID, assertion := range s.assertions {
		if _, exists := s.entities[assertion.Subject()]; !exists {
			warnings = append(warnings, fmt.Sprintf("assertion %s references non-existent subject %s", assertionID, assertion.Subject()))
		}
		if _, exists := s.entities[assertion.Object()]; !exists {
			warnings = append(warnings, fmt.Sprintf("assertion %s references non-existent object %s", assertionID, assertion.Object()))
		}
	}

	// Check for orphaned entities (entities with no assertions)
	for entityID := range s.entities {
		hasAssertions := false
		for _, assertion := range s.assertions {
			if assertion.Subject() == entityID || assertion.Object() == entityID {
				hasAssertions = true
				break
			}
		}
		if !hasAssertions {
			warnings = append(warnings, fmt.Sprintf("entity %s has no assertions", entityID))
		}
	}

	return warnings
}

// Clear removes all data from the semantic store
func (s *SemanticStore) Clear() {
	s.entities = make(map[string]*EntityReference)
	s.relations = make(map[string]*kmac.Relation)
	s.assertions = make(map[string]*kmac.Assertion)
	s.properties = make(map[string]*kmac.Property)
}