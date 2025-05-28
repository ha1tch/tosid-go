package kmac

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
	
	// ValidateStatement validates a KMAC statement
	ValidateStatement(statement Statement) error
}

// KnowledgeBase is an interface for KMAC knowledge bases
type KnowledgeBase interface {
	// AddStatement adds a statement to the knowledge base
	AddStatement(statement Statement) error
	
	// GetStatement retrieves a statement by ID
	GetStatement(id string) (Statement, error)
	
	// RemoveStatement removes a statement from the knowledge base
	RemoveStatement(id string) error
	
	// Query performs a query on the knowledge base
	Query(query string) ([]Statement, error)
	
	// GetAllStatements returns all statements in the knowledge base
	GetAllStatements() []Statement
}

// ConfidenceManager is an interface for managing confidence levels in KMAC
type ConfidenceManager interface {
	// SetConfidence sets the confidence level for a statement
	SetConfidence(statementID string, level float64, source string) error
	
	// GetConfidence retrieves the confidence level for a statement
	GetConfidence(statementID string) (float64, string, error)
	
	// UpdateConfidence updates the confidence level for a statement
	UpdateConfidence(statementID string, level float64, source string) error
}

// ReasoningEngine is an interface for reasoning over KMAC statements
type ReasoningEngine interface {
	// Infer performs inference to derive new statements
	Infer(statements []Statement) ([]Statement, error)
	
	// CheckConsistency checks for consistency in a set of statements
	CheckConsistency(statements []Statement) (bool, []string)
	
	// FindConflicts finds conflicting statements
	FindConflicts(statements []Statement) ([]StatementPair, error)
}

// StatementPair represents a pair of related statements
type StatementPair struct {
	First  Statement
	Second Statement
	Relationship string
}

// Serializer is an interface for serializing KMAC statements
type Serializer interface {
	// Serialize converts statements to a serialized format
	Serialize(statements []Statement) ([]byte, error)
	
	// Deserialize converts serialized data back to statements
	Deserialize(data []byte) ([]Statement, error)
	
	// SerializeToString converts statements to string format
	SerializeToString(statements []Statement) (string, error)
	
	// DeserializeFromString converts string data back to statements
	DeserializeFromString(data string) ([]Statement, error)
}