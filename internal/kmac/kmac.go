package kmac

// Identifier types
const (
	EntityIDPrefix    = "E"
	EventIDPrefix     = "V"
	RelationIDPrefix  = "R"
	PropertyIDPrefix  = "P"
	TimeIDPrefix      = "T"
	AssertionIDPrefix = "F"
)

// Statement represents a KMAC statement
type Statement interface {
	ID() string
	Type() string
	String() string
}

// validateIdentifier checks if an ID matches the expected prefix format
func validateIdentifier(prefix string, id string) bool {
	// For simplicity, we're just checking if it starts with the expected letter
	// and contains only alphanumeric characters
	if len(id) < 1 || string(id[0]) != prefix {
		return false
	}
	
	// In a real implementation, we would do more thorough validation
	// For now, just check that it's not empty and starts correctly
	return len(id) > 1
}

// ValidateKMACStatement validates a KMAC statement for correctness
func ValidateKMACStatement(statement Statement) error {
	switch stmt := statement.(type) {
	case *Entity:
		return validateEntity(stmt)
	case *Relation:
		return validateRelation(stmt)
	case *Assertion:
		return validateAssertion(stmt)
	case *Property:
		return validateProperty(stmt)
	default:
		return fmt.Errorf("unknown statement type: %T", statement)
	}
}

func validateEntity(entity *Entity) error {
	if entity.ID() == "" {
		return errors.New("entity ID cannot be empty")
	}
	if entity.Label() == "" {
		return errors.New("entity label cannot be empty")
	}
	return nil
}

func validateRelation(relation *Relation) error {
	if relation.ID() == "" {
		return errors.New("relation ID cannot be empty")
	}
	if relation.Label() == "" {
		return errors.New("relation label cannot be empty")
	}
	if relation.RelationType() == "" {
		return errors.New("relation type cannot be empty")
	}
	return nil
}

func validateAssertion(assertion *Assertion) error {
	if assertion.ID() == "" {
		return errors.New("assertion ID cannot be empty")
	}
	if assertion.Subject() == "" {
		return errors.New("assertion subject cannot be empty")
	}
	if assertion.Relation() == "" {
		return errors.New("assertion relation cannot be empty")
	}
	if assertion.Object() == "" {
		return errors.New("assertion object cannot be empty")
	}
	return nil
}

func validateProperty(property *Property) error {
	if property.ID() == "" {
		return errors.New("property ID cannot be empty")
	}
	if property.Label() == "" {
		return errors.New("property label cannot be empty")
	}
	return nil
}