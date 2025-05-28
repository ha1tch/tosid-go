package integration

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ha1tch/tosid-go/pkg/kmac"
	"github.com/ha1tch/tosid-go/pkg/tosid"
)

// KMACTOSIDConverter provides functionality to convert between TOSID and KMAC systems
type KMACTOSIDConverter struct{}

// ConvertTOSIDToKMACEntity converts a TOSID code to a KMAC entity
func (c *KMACTOSIDConverter) ConvertTOSIDToKMACEntity(id string, label string, tosidCode string) (*kmac.Entity, error) {
	// Validate the TOSID code first
	_, err := tosid.Parse(tosidCode)
	if err != nil {
		return nil, fmt.Errorf("invalid TOSID code: %v", err)
	}

	// Create a KMAC entity with the TOSID type
	return kmac.NewEntity(id, label, tosidCode)
}

// ExtractTOSIDFromKMACEntity extracts a TOSID code from a KMAC entity
func (c *KMACTOSIDConverter) ExtractTOSIDFromKMACEntity(entity *kmac.Entity) (*tosid.TOSID, error) {
	tosidType := entity.TOSIDType()
	if tosidType == "" {
		return nil, errors.New("entity does not have a TOSID type")
	}

	return tosid.Parse(tosidType)
}

// FindEntitiesByTOSIDPattern finds KMAC entities matching a TOSID pattern
func (c *KMACTOSIDConverter) FindEntitiesByTOSIDPattern(entities []*kmac.Entity, pattern string) []*kmac.Entity {
	var results []*kmac.Entity

	for _, entity := range entities {
		tosidObj, err := c.ExtractTOSIDFromKMACEntity(entity)
		if err != nil {
			continue
		}

		if tosidObj.MatchesPattern(pattern) {
			results = append(results, entity)
		}
	}

	return results
}

// GenerateKMACFromTOSIDHierarchy generates KMAC statements from a TOSID hierarchy
func (c *KMACTOSIDConverter) GenerateKMACFromTOSIDHierarchy(tosidCodes []string, labels map[string]string) ([]kmac.Statement, error) {
	var statements []kmac.Statement
	entityMap := make(map[string]*kmac.Entity)
	idCounter := 1

	// First pass: create entities
	for _, code := range tosidCodes {
		// Parse the TOSID code
		tosidObj, err := tosid.Parse(code)
		if err != nil {
			return nil, fmt.Errorf("invalid TOSID code %s: %v", code, err)
		}

		// Get or generate a label
		label, ok := labels[code]
		if !ok {
			// Generate a label from the TOSID code
			parts := strings.Split(tosidObj.Identifier, "-")
			if len(parts) > 0 {
				label = parts[len(parts)-1]
			} else {
				label = fmt.Sprintf("Entity%d", idCounter)
			}
		}

		// Create a KMAC entity
		entityID := fmt.Sprintf("E%04d", idCounter)
		entity, err := kmac.NewEntity(entityID, label, code)
		if err != nil {
			return nil, fmt.Errorf("failed to create entity for %s: %v", code, err)
		}

		statements = append(statements, entity)
		entityMap[code] = entity
		idCounter++
	}

	// Second pass: create part-of relationships based on TOSID hierarchy
	relationCounter := 1
	for _, code := range tosidCodes {
		tosidObj, err := tosid.Parse(code)
		if err != nil {
			continue
		}

		// Check if this TOSID has a parent in our collection
		for parentCode, parentEntity := range entityMap {
			parentTosid, err := tosid.Parse(parentCode)
			if err != nil {
				continue
			}

			// A TOSID is a "child" if it is more specific than another TOSID
			// and shares the same prefix
			if strings.HasPrefix(tosidObj.String(), parentTosid.TaxonomyCode+parentTosid.NetmaskIndicator) &&
				tosidObj.String() != parentTosid.String() {
				
				// Create a part-of relationship
				partOf, err := kmac.NewPartOf(entityMap[code].ID(), parentEntity.ID())
				if err != nil {
					continue
				}
				
				statements = append(statements, partOf)
			}
		}
	}

	return statements, nil
}

// GenerateClassHierarchyFromTOSID generates a class hierarchy for a given TOSID
func (c *KMACTOSIDConverter) GenerateClassHierarchyFromTOSID(tosidCode string) ([]string, error) {
	// Parse the TOSID code
	tosidObj, err := tosid.Parse(tosidCode)
	if err != nil {
		return nil, fmt.Errorf("invalid TOSID code: %v", err)
	}

	// Generate hierarchy levels
	var hierarchy []string

	// First level: Domain level (e.g., 00 - Natural Material)
	domainCode := tosidObj.TaxonomyCode
	hierarchy = append(hierarchy, domainCode)

	// Second level: Add netmask (e.g., 00B - Natural Material, Stellar Scale)
	domainWithNetmask := domainCode + tosidObj.NetmaskIndicator
	hierarchy = append(hierarchy, domainWithNetmask)

	// Extract parts of the identifier
	identifier := tosidObj.Identifier
	parts := strings.Split(identifier, "-")
	if len(parts) < 3 {
		return hierarchy, nil // Return what we have if identifier format is unexpected
	}

	// Third level: Add first category (e.g., 00B-SOL)
	level3 := domainWithNetmask + "-" + parts[0]
	hierarchy = append(hierarchy, level3)

	// Fourth level: Add second category (e.g., 00B-SOL-STR)
	level4 := level3 + "-" + parts[1]
	hierarchy = append(hierarchy, level4)

	// Fifth level: Complete category (e.g., 00B-SOL-STR-SUN)
	level5 := level4 + "-" + parts[2]
	hierarchy = append(hierarchy, level5)

	// Add the full TOSID if it has specific identifier after the colon
	if strings.Contains(identifier, ":") {
		hierarchy = append(hierarchy, tosidCode)
	}

	return hierarchy, nil
}

// CreateKMACAssertionsFromTOSIDRelationships creates KMAC assertions based on TOSID relationships
func (c *KMACTOSIDConverter) CreateKMACAssertionsFromTOSIDRelationships(
	relationships []struct {
		Subject string
		Relation string
		Object string
	},
	entityMap map[string]*kmac.Entity,
	relationMap map[string]*kmac.Relation,
) ([]*kmac.Assertion, error) {
	var assertions []*kmac.Assertion
	idCounter := 1

	for _, rel := range relationships {
		// Get the entities and relation
		subjectEntity, subjectExists := entityMap[rel.Subject]
		objectEntity, objectExists := entityMap[rel.Object]
		relation, relationExists := relationMap[rel.Relation]

		if !subjectExists {
			return nil, fmt.Errorf("subject entity not found: %s", rel.Subject)
		}
		if !objectExists {
			return nil, fmt.Errorf("object entity not found: %s", rel.Object)
		}
		if !relationExists {
			return nil, fmt.Errorf("relation not found: %s", rel.Relation)
		}

		// Create the assertion
		assertionID := fmt.Sprintf("F%04d", idCounter)
		assertion, err := kmac.NewAssertion(assertionID, subjectEntity.ID(), relation.ID(), objectEntity.ID())
		if err != nil {
			return nil, fmt.Errorf("failed to create assertion: %v", err)
		}

		assertions = append(assertions, assertion)
		idCounter++
	}

	return assertions, nil
}

// ExtractSemanticInfoFromTOSID extracts semantic information from a TOSID code
func (c *KMACTOSIDConverter) ExtractSemanticInfoFromTOSID(tosidCode string) (map[string]string, error) {
	// Parse the TOSID code
	tosidObj, err := tosid.Parse(tosidCode)
	if err != nil {
		return nil, fmt.Errorf("invalid TOSID code: %v", err)
	}

	// Extract semantic information
	info := make(map[string]string)

	// Extract domain information
	domainDigit := string(tosidObj.TaxonomyCode[0])
	if domain, ok := tosid.TaxonomyDomains[domainDigit]; ok {
		info["domain"] = domain
	}

	// Extract type information
	typeDigit := string(tosidObj.TaxonomyCode[1])
	if typeDesc, ok := tosid.TaxonomyTypes[typeDigit]; ok {
		info["type"] = typeDesc
	}

	// Extract scope information
	if scopes, ok := tosid.NetmaskDescriptions[tosidObj.TaxonomyCode]; ok {
		if scope, ok := scopes[tosidObj.NetmaskIndicator]; ok {
			info["scope"] = scope
		}
	}

	// Extract identifier parts
	identifier := tosidObj.Identifier
	parts := strings.Split(identifier, "-")
	if len(parts) >= 3 {
		info["category1"] = parts[0]
		info["category2"] = parts[1]
		info["category3"] = parts[2]
	}

	// Extract specific identifier if present
	if strings.Contains(identifier, ":") {
		specificParts := strings.Split(identifier, ":")
		if len(specificParts) > 1 {
			info["specific_identifier"] = specificParts[1]
		}
	}

	return info, nil
}

// NewKMACTOSIDConverter creates a new KMAC-TOSID converter
func NewKMACTOSIDConverter() *KMACTOSIDConverter {
	return &KMACTOSIDConverter{}
}
