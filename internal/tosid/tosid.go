package tosid

import (
	"fmt"
	"regexp"
	"strings"
)

// TOSID represents a Taxonomic Ontological Semantic IDentification
type TOSID struct {
	TaxonomyCode     string // TT
	NetmaskIndicator string // N
	Identifier       string // XXX-XXX-XXX:XXX-XXX-XXX-XXX
}

// String returns the string representation of the TOSID
func (t *TOSID) String() string {
	return fmt.Sprintf("%s%s-%s", t.TaxonomyCode, t.NetmaskIndicator, t.Identifier)
}

// ClassificationDescription returns a human-readable description of the TOSID classification
func (t *TOSID) ClassificationDescription() string {
	classifier := NewTaxonomyClassifier()
	return classifier.GetFullClassification(t.TaxonomyCode, t.NetmaskIndicator)
}

// IsCompatibleWith checks if this TOSID is compatible with another TOSID
// Two TOSIDs are compatible if they share the same taxonomy and netmask
func (t *TOSID) IsCompatibleWith(other *TOSID) bool {
	return t.TaxonomyCode == other.TaxonomyCode && t.NetmaskIndicator == other.NetmaskIndicator
}

// MatchesPattern checks if a TOSID matches a pattern with wildcards
// Example pattern: "00B*"
func (t *TOSID) MatchesPattern(pattern string) bool {
	if len(pattern) == 0 {
		return true
	}
	
	tosidStr := t.String()
	
	// Convert pattern to regex
	regexPattern := "^"
	for _, c := range pattern {
		if c == '*' {
			regexPattern += ".*"
		} else {
			regexPattern += regexp.QuoteMeta(string(c))
		}
	}
	regexPattern += ".*$" // Match the rest of the string
	
	matched, _ := regexp.MatchString(regexPattern, tosidStr)
	return matched
}

// GetHierarchy returns the hierarchical levels of this TOSID
func (t *TOSID) GetHierarchy() []string {
	var hierarchy []string
	
	// Level 1: Domain
	hierarchy = append(hierarchy, t.TaxonomyCode)
	
	// Level 2: Domain + Netmask
	hierarchy = append(hierarchy, t.TaxonomyCode+t.NetmaskIndicator)
	
	// Extract identifier parts
	identifierParts := strings.Split(t.Identifier, ":")
	categoryPart := identifierParts[0]
	
	categories := strings.Split(categoryPart, "-")
	if len(categories) >= 3 {
		// Level 3: Domain + Netmask + First category
		level3 := t.TaxonomyCode + t.NetmaskIndicator + "-" + categories[0]
		hierarchy = append(hierarchy, level3)
		
		// Level 4: Domain + Netmask + First two categories
		level4 := level3 + "-" + categories[1]
		hierarchy = append(hierarchy, level4)
		
		// Level 5: Complete category part
		level5 := level4 + "-" + categories[2]
		hierarchy = append(hierarchy, level5)
	}
	
	// Level 6: Full TOSID if it has specific identifier
	if len(identifierParts) > 1 {
		hierarchy = append(hierarchy, t.String())
	}
	
	return hierarchy
}

// GetParent returns the parent TOSID at the next higher hierarchical level
func (t *TOSID) GetParent() *TOSID {
	hierarchy := t.GetHierarchy()
	if len(hierarchy) <= 1 {
		return nil // No parent
	}
	
	// Get the parent level (second to last in hierarchy)
	parentStr := hierarchy[len(hierarchy)-2]
	
	// Parse the parent string to create a TOSID
	parser := NewParser()
	parent, err := parser.Parse(parentStr)
	if err != nil {
		return nil
	}
	
	return parent
}

// GetChildren returns potential child TOSID patterns
func (t *TOSID) GetChildren() []string {
	// This would return patterns that could match children
	// Implementation depends on specific use case
	return []string{t.String() + "*"}
}

// IsParentOf checks if this TOSID is a parent of another TOSID
func (t *TOSID) IsParentOf(other *TOSID) bool {
	thisStr := t.String()
	otherStr := other.String()
	
	// A TOSID is a parent if the other starts with this TOSID's string
	// and is more specific
	return strings.HasPrefix(otherStr, thisStr) && len(otherStr) > len(thisStr)
}

// IsChildOf checks if this TOSID is a child of another TOSID
func (t *TOSID) IsChildOf(other *TOSID) bool {
	return other.IsParentOf(t)
}

// GetDepth returns the hierarchical depth of this TOSID
func (t *TOSID) GetDepth() int {
	return len(t.GetHierarchy())
}