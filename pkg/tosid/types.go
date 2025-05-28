package tosid

// Types provides interface definitions for TOSID types

// TOSIDParser is an interface for parsing TOSID codes
type TOSIDParser interface {
	// Parse creates a TOSID from a string representation
	Parse(code string) (*TOSID, error)
}

// TOSIDCreator is an interface for creating TOSID codes
type TOSIDCreator interface {
	// Create creates a new TOSID with the specified components
	Create(taxonomyCode, netmaskIndicator, identifier string) (*TOSID, error)
}

// TOSIDClassifier is an interface for working with TOSID classifications
type TOSIDClassifier interface {
	// ClassificationDescription returns a human-readable description of the TOSID classification
	ClassificationDescription() string
	
	// IsCompatibleWith checks if this TOSID is compatible with another TOSID
	IsCompatibleWith(other *TOSID) bool
	
	// MatchesPattern checks if a TOSID matches a pattern with wildcards
	MatchesPattern(pattern string) bool
}

// TOSID public interfaces
var (
	Parser    TOSIDParser    = tosidParser{}
	Creator   TOSIDCreator   = tosidCreator{}
	Classifier TOSIDClassifier = tosidClassifier{}
)

// Internal implementations
type tosidParser struct{}
type tosidCreator struct{}
type tosidClassifier struct{}

func (p tosidParser) Parse(code string) (*TOSID, error) {
	return Parse(code)
}

func (c tosidCreator) Create(taxonomyCode, netmaskIndicator, identifier string) (*TOSID, error) {
	return Create(taxonomyCode, netmaskIndicator, identifier)
}
