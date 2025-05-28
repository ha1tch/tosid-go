package tosid

// TOSIDParser is an interface for parsing TOSID codes
type TOSIDParser interface {
	// Parse creates a TOSID from a string representation
	Parse(code string) (*TOSID, error)
	
	// ParseBatch parses multiple TOSID codes
	ParseBatch(codes []string) ([]*TOSID, []error)
	
	// ValidateFormat checks if a string matches TOSID format
	ValidateFormat(code string) bool
}

// TOSIDCreator is an interface for creating TOSID codes
type TOSIDCreator interface {
	// Create creates a new TOSID with the specified components
	Create(taxonomyCode, netmaskIndicator, identifier string) (*TOSID, error)
	
	// CreateFromTemplate creates a TOSID from a template
	CreateFromTemplate(template string, values map[string]string) (*TOSID, error)
	
	// GenerateNext generates the next TOSID in a sequence
	GenerateNext(base *TOSID) (*TOSID, error)
}

// TOSIDClassifier is an interface for working with TOSID classifications
type TOSIDClassifier interface {
	// ClassificationDescription returns a human-readable description
	ClassificationDescription() string
	
	// IsCompatibleWith checks if this TOSID is compatible with another
	IsCompatibleWith(other *TOSID) bool
	
	// MatchesPattern checks if a TOSID matches a pattern
	MatchesPattern(pattern string) bool
	
	// GetHierarchy returns the hierarchical levels
	GetHierarchy() []string
}

// TOSIDValidator is an interface for validating TOSID codes
type TOSIDValidator interface {
	// ValidateFormat validates the basic format
	ValidateFormat(code string) error
	
	// ValidateComponents validates individual components
	ValidateComponents(taxonomyCode, netmaskIndicator, identifier string) error
	
	// ValidateSemanticConsistency checks semantic consistency
	ValidateSemanticConsistency(tosid *TOSID) []string
	
	// IsWellFormed checks if a TOSID is well-formed
	IsWellFormed(tosid *TOSID) (bool, []string)
}

// TOSIDRepository is an interface for storing and retrieving TOSID codes
type TOSIDRepository interface {
	// Store stores a TOSID
	Store(tosid *TOSID) error
	
	// Retrieve retrieves a TOSID by its string representation
	Retrieve(code string) (*TOSID, error)
	
	// FindByPattern finds TOSIDs matching a pattern
	FindByPattern(pattern string) ([]*TOSID, error)
	
	// ListAll lists all stored TOSIDs
	ListAll() ([]*TOSID, error)
	
	// Delete deletes a TOSID
	Delete(code string) error
}

// TOSIDAnalyzer is an interface for analyzing TOSID relationships
type TOSIDAnalyzer interface {
	// FindRelated finds TOSIDs related to a given TOSID
	FindRelated(tosid *TOSID) ([]*TOSID, error)
	
	// BuildHierarchy builds a hierarchical structure
	BuildHierarchy(tosids []*TOSID) (*TOSIDHierarchy, error)
	
	// CompareClassifications compares two TOSIDs
	CompareClassifications(first, second *TOSID) (*ComparisonResult, error)
}

// TOSIDHierarchy represents a hierarchical structure of TOSIDs
type TOSIDHierarchy struct {
	Root     *TOSID
	Children []*TOSIDHierarchy
	Level    int
}

// ComparisonResult represents the result of comparing two TOSIDs
type ComparisonResult struct {
	Compatible    bool
	SharedLevels  int
	Differences   []string
	Relationship  string // "parent", "child", "sibling", "unrelated"
}