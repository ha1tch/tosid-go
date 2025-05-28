package tosid

import (
	"fmt"
	"sort"
)

// TOSIDCollection represents a collection of TOSID codes
type TOSIDCollection struct {
	tosids map[string]*TOSID
}

// NewTOSIDCollection creates a new TOSID collection
func NewTOSIDCollection() *TOSIDCollection {
	return &TOSIDCollection{
		tosids: make(map[string]*TOSID),
	}
}

// Add adds a TOSID to the collection
func (tc *TOSIDCollection) Add(tosid *TOSID) error {
	if tosid == nil {
		return fmt.Errorf("cannot add nil TOSID")
	}

	validator := NewValidator()
	if valid, warnings := validator.IsWellFormed(tosid); !valid {
		return fmt.Errorf("invalid TOSID: %v", warnings)
	}

	tc.tosids[tosid.String()] = tosid
	return nil
}

// Get retrieves a TOSID by its string representation
func (tc *TOSIDCollection) Get(code string) (*TOSID, bool) {
	tosid, exists := tc.tosids[code]
	return tosid, exists
}

// Remove removes a TOSID by its string representation
func (tc *TOSIDCollection) Remove(code string) bool {
	if _, exists := tc.tosids[code]; exists {
		delete(tc.tosids, code)
		return true
	}
	return false
}

// GetAll returns all TOSIDs
func (tc *TOSIDCollection) GetAll() []*TOSID {
	tosids := make([]*TOSID, 0, len(tc.tosids))
	for _, tosid := range tc.tosids {
		tosids = append(tosids, tosid)
	}
	return tosids
}

// FindByPattern finds TOSIDs matching a pattern
func (tc *TOSIDCollection) FindByPattern(pattern string) []*TOSID {
	var matches []*TOSID
	for _, tosid := range tc.tosids {
		if tosid.MatchesPattern(pattern) {
			matches = append(matches, tosid)
		}
	}
	return matches
}

// GetByTaxonomy returns all TOSIDs with the specified taxonomy code
func (tc *TOSIDCollection) GetByTaxonomy(taxonomyCode string) []*TOSID {
	var matches []*TOSID
	for _, tosid := range tc.tosids {
		if tosid.TaxonomyCode == taxonomyCode {
			matches = append(matches, tosid)
		}
	}
	return matches
}

// GetByNetmask returns all TOSIDs with the specified netmask indicator
func (tc *TOSIDCollection) GetByNetmask(netmaskIndicator string) []*TOSID {
	var matches []*TOSID
	for _, tosid := range tc.tosids {
		if tosid.NetmaskIndicator == netmaskIndicator {
			matches = append(matches, tosid)
		}
	}
	return matches
}

// Count returns the number of TOSIDs in the collection
func (tc *TOSIDCollection) Count() int {
	return len(tc.tosids)
}

// Clear removes all TOSIDs
func (tc *TOSIDCollection) Clear() {
	tc.tosids = make(map[string]*TOSID)
}

// GetStatistics returns statistics about the collection
func (tc *TOSIDCollection) GetStatistics() map[string]int {
	stats := make(map[string]int)

	taxonomyCount := make(map[string]int)
	netmaskCount := make(map[string]int)

	for _, tosid := range tc.tosids {
		taxonomyCount[tosid.TaxonomyCode]++
		netmaskCount[tosid.NetmaskIndicator]++
	}

	stats["total"] = len(tc.tosids)

	for taxonomy, count := range taxonomyCount {
		stats["taxonomy_"+taxonomy] = count
	}

	for netmask, count := range netmaskCount {
		stats["netmask_"+netmask] = count
	}

	return stats
}

// BuildHierarchy builds a hierarchical representation of the TOSIDs
func (tc *TOSIDCollection) BuildHierarchy() map[string]map[string][]*TOSID {
	hierarchy := make(map[string]map[string][]*TOSID)

	for _, tosid := range tc.tosids {
		taxonomyCode := tosid.TaxonomyCode
		netmaskIndicator := tosid.NetmaskIndicator

		if hierarchy[taxonomyCode] == nil {
			hierarchy[taxonomyCode] = make(map[string][]*TOSID)
		}

		hierarchy[taxonomyCode][netmaskIndicator] = append(
			hierarchy[taxonomyCode][netmaskIndicator], tosid)
	}

	return hierarchy
}

// ExportToStrings exports all TOSID codes as strings
func (tc *TOSIDCollection) ExportToStrings() []string {
	codes := make([]string, 0, len(tc.tosids))
	for code := range tc.tosids {
		codes = append(codes, code)
	}
	sort.Strings(codes)
	return codes
}

// TOSIDGenerator helps generate TOSID codes
type TOSIDGenerator struct {
	taxonomyCode     string
	netmaskIndicator string
	baseIdentifier   string
	counter          int
}

// NewTOSIDGenerator creates a new TOSID generator
func NewTOSIDGenerator(taxonomyCode, netmaskIndicator, baseIdentifier string) (*TOSIDGenerator, error) {
	validator := NewValidator()
	if err := validator.ValidateTaxonomyCode(taxonomyCode); err != nil {
		return nil, err
	}
	if err := validator.ValidateNetmaskIndicator(taxonomyCode, netmaskIndicator); err != nil {
		return nil, err
	}

	return &TOSIDGenerator{
		taxonomyCode:     taxonomyCode,
		netmaskIndicator: netmaskIndicator,
		baseIdentifier:   baseIdentifier,
		counter:          1,
	}, nil
}

// Generate generates the next TOSID in sequence
func (tg *TOSIDGenerator) Generate() (*TOSID, error) {
	identifier := fmt.Sprintf("%s:%03d-000-000-001", tg.baseIdentifier, tg.counter)

	tosid := &TOSID{
		TaxonomyCode:     tg.taxonomyCode,
		NetmaskIndicator: tg.netmaskIndicator,
		Identifier:       identifier,
	}

	tg.counter++
	return tosid, nil
}

// GenerateWithSuffix generates a TOSID with a custom suffix
func (tg *TOSIDGenerator) GenerateWithSuffix(suffix string) (*TOSID, error) {
	identifier := fmt.Sprintf("%s:%s", tg.baseIdentifier, suffix)

	validator := NewValidator()
	if err := validator.ValidateIdentifier(identifier); err != nil {
		return nil, err
	}

	tosid := &TOSID{
		TaxonomyCode:     tg.taxonomyCode,
		NetmaskIndicator: tg.netmaskIndicator,
		Identifier:       identifier,
	}

	return tosid, nil
}

// Reset resets the counter
func (tg *TOSIDGenerator) Reset() {
	tg.counter = 1
}

// SetCounter sets the counter to a specific value
func (tg *TOSIDGenerator) SetCounter(counter int) {
	if counter > 0 {
		tg.counter = counter
	}
}
