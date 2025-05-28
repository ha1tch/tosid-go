package kmac

import (
	"testing"
)

func TestEntityCreation(t *testing.T) {
	entity, err := NewEntity("E1001", "Test Entity", "00B2-SOL-STR-SUN:000-000-000-001")
	if err != nil {
		t.Fatalf("Failed to create entity: %v", err)
	}

	if entity.ID() != "E1001" {
		t.Errorf("Expected ID E1001, got %s", entity.ID())
	}

	if entity.Label() != "Test Entity" {
		t.Errorf("Expected label 'Test Entity', got %s", entity.Label())
	}

	if entity.TOSIDType() != "00B2-SOL-STR-SUN:000-000-000-001" {
		t.Errorf("Expected TOSID type '00B2-SOL-STR-SUN:000-000-000-001', got %s", entity.TOSIDType())
	}
}

func TestRelationCreation(t *testing.T) {
	relation, err := NewRelation("R1001", "Test Relation", "TEST_TYPE")
	if err != nil {
		t.Fatalf("Failed to create relation: %v", err)
	}

	if relation.ID() != "R1001" {
		t.Errorf("Expected ID R1001, got %s", relation.ID())
	}

	if relation.Label() != "Test Relation" {
		t.Errorf("Expected label 'Test Relation', got %s", relation.Label())
	}

	if relation.RelationType() != "TEST_TYPE" {
		t.Errorf("Expected relation type 'TEST_TYPE', got %s", relation.RelationType())
	}
}

func TestAssertionCreation(t *testing.T) {
	assertion, err := NewAssertion("F1001", "E1001", "R1001", "E1002")
	if err != nil {
		t.Fatalf("Failed to create assertion: %v", err)
	}

	if assertion.Subject() != "E1001" {
		t.Errorf("Expected subject E1001, got %s", assertion.Subject())
	}

	if assertion.Relation() != "R1001" {
		t.Errorf("Expected relation R1001, got %s", assertion.Relation())
	}

	if assertion.Object() != "E1002" {
		t.Errorf("Expected object E1002, got %s", assertion.Object())
	}
}

func TestConfidenceLevels(t *testing.T) {
	assertion, _ := NewAssertion("F1001", "E1001", "R1001", "E1002")
	
	// Test default confidence
	confidence, source := assertion.GetConfidence()
	if confidence != 1.0 {
		t.Errorf("Expected default confidence 1.0, got %f", confidence)
	}

	// Test setting confidence
	assertion.SetConfidence(0.85, "TEST_SOURCE")
	confidence, source = assertion.GetConfidence()

	if confidence != 0.85 {
		t.Errorf("Expected confidence 0.85, got %f", confidence)
	}

	if source != "TEST_SOURCE" {
		t.Errorf("Expected source 'TEST_SOURCE', got %s", source)
	}
}

func BenchmarkEntityCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := NewEntity("E1001", "Test Entity", "00B2-SOL-STR-SUN:000-000-000-001")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAssertionCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := NewAssertion("F1001", "E1001", "R1001", "E1002")
		if err != nil {
			b.Fatal(err)
		}
	}
}