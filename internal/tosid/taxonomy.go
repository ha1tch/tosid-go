package tosid

// TaxonomyDomains maps the first digit to domain descriptions
var TaxonomyDomains = map[string]string{
	"0": "Celestial/Natural",
	"1": "Artificial/Intelligent",
}

// TaxonomyTypes maps the second digit to type descriptions
var TaxonomyTypes = map[string]string{
	"0": "Physical/Material",
	"1": "Conceptual/Abstract",
}

// NetmaskDescriptions provides descriptions for netmask scopes by taxonomy prefix
var NetmaskDescriptions = map[string]map[string]string{
	"00": { // Natural Material
		"A": "Cosmic Scale",
		"B": "Stellar Scale",
		"C": "Planetary Scale",
		"D": "Regional Scale",
		"E": "Local Scale",
		"F": "Microscopic Scale",
	},
	"01": { // Natural Conceptual
		"A": "Universal Laws",
		"B": "Systemic Principles",
		"C": "Historical Events",
		"D": "Cyclic Phenomena",
		"E": "Emergent Patterns",
	},
	"10": { // Artificial Material
		"A": "Megastructures",
		"B": "Buildings",
		"C": "Complex Objects",
		"D": "Tools/Devices",
		"E": "Components",
		"F": "Manufactured Materials",
	},
	"11": { // Artificial Conceptual
		"A": "Civilizational Systems",
		"B": "Organized Knowledge",
		"C": "Cultural Expressions",
		"D": "Designed Processes",
		"E": "Linguistic Constructs",
	},
}

// BiologicalHierarchyScopes provides scope indicators for biological entities
var BiologicalHierarchyScopes = map[string]string{
	"1": "Ecosystem/Population",
	"2": "Organism",
	"3": "Organ System",
	"4": "Organ",
	"5": "Tissue",
	"6": "Cell",
	"7": "Organelle",
	"8": "Molecular",
	"9": "Genomic",
}

// TaxonomyClassifier provides classification utilities
type TaxonomyClassifier struct{}

// NewTaxonomyClassifier creates a new taxonomy classifier
func NewTaxonomyClassifier() *TaxonomyClassifier {
	return &TaxonomyClassifier{}
}

// GetDomainDescription returns the domain description for a taxonomy code
func (tc *TaxonomyClassifier) GetDomainDescription(taxonomyCode string) string {
	if len(taxonomyCode) < 1 {
		return "Unknown Domain"
	}
	
	if desc, exists := TaxonomyDomains[taxonomyCode[:1]]; exists {
		return desc
	}
	return "Unknown Domain"
}

// GetTypeDescription returns the type description for a taxonomy code
func (tc *TaxonomyClassifier) GetTypeDescription(taxonomyCode string) string {
	if len(taxonomyCode) < 2 {
		return "Unknown Type"
	}
	
	if desc, exists := TaxonomyTypes[taxonomyCode[1:2]]; exists {
		return desc
	}
	return "Unknown Type"
}

// GetScopeDescription returns the scope description for a taxonomy code and netmask
func (tc *TaxonomyClassifier) GetScopeDescription(taxonomyCode, netmaskIndicator string) string {
	if scopes, exists := NetmaskDescriptions[taxonomyCode]; exists {
		if desc, exists := scopes[netmaskIndicator]; exists {
			return desc
		}
	}
	return "Unknown Scope"
}

// GetFullClassification returns the complete classification description
func (tc *TaxonomyClassifier) GetFullClassification(taxonomyCode, netmaskIndicator string) string {
	domain := tc.GetDomainDescription(taxonomyCode)
	typeDesc := tc.GetTypeDescription(taxonomyCode)
	scope := tc.GetScopeDescription(taxonomyCode, netmaskIndicator)
	
	return domain + " - " + typeDesc + " - " + scope
}

// IsValidTaxonomyCode checks if a taxonomy code is valid
func (tc *TaxonomyClassifier) IsValidTaxonomyCode(taxonomyCode string) bool {
	if len(taxonomyCode) != 2 {
		return false
	}
	
	_, domainExists := TaxonomyDomains[taxonomyCode[:1]]
	_, typeExists := TaxonomyTypes[taxonomyCode[1:2]]
	
	return domainExists && typeExists
}

// IsValidNetmaskIndicator checks if a netmask indicator is valid for a taxonomy code
func (tc *TaxonomyClassifier) IsValidNetmaskIndicator(taxonomyCode, netmaskIndicator string) bool {
	if scopes, exists := NetmaskDescriptions[taxonomyCode]; exists {
		_, exists := scopes[netmaskIndicator]
		return exists
	}
	return false
}

// GetCompatibleScopes returns all valid netmask indicators for a taxonomy code
func (tc *TaxonomyClassifier) GetCompatibleScopes(taxonomyCode string) []string {
	var scopes []string
	if scopeMap, exists := NetmaskDescriptions[taxonomyCode]; exists {
		for scope := range scopeMap {
			scopes = append(scopes, scope)
		}
	}
	return scopes
}