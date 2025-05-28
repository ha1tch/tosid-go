package semantic

import (
	"fmt"
	"testing"
)

func TestSemanticStoreBasicOperations(t *testing.T) {
	store := NewSemanticStore()

	// Test adding entities
	err := store.AddEntity("E1001", "Sun", "00B2-SOL-STR-SUN:000-000-000-001")
	if err != nil {
		t.Fatalf("Failed to add entity: %v", err)
	}

	err = store.AddEntity("E1002", "Earth", "00B3-SOL-SYS-ERT:000-000-000-001")
	if err != nil {
		t.Fatalf("Failed to add entity: %v", err)
	}

	// Test entity retrieval
	entity, err := store.GetEntity("E1001")
	if err != nil {
		t.Fatalf("Failed to get entity: %v", err)
	}

	if entity.KMACEntity.Label() != "Sun" {
		t.Errorf("Expected label 'Sun', got %s", entity.KMACEntity.Label())
	}

	// Test TOSID pattern matching
	celestialBodies := store.FindEntitiesByTOSIDPattern("00B")
	if len(celestialBodies) != 2 {
		t.Errorf("Expected 2 celestial bodies, got %d", len(celestialBodies))
	}

	// Test assertion creation
	err = store.CreateAssertion("F1001", "E1002", "ORBITS", "E1001")
	if err != nil {
		t.Fatalf("Failed to create assertion: %v", err)
	}

	// Test finding assertions
	assertions := store.FindAssertionsForEntity("E1001")
	if len(assertions) != 1 {
		t.Errorf("Expected 1 assertion, got %d", len(assertions))
	}
}

func TestSemanticStoreStatistics(t *testing.T) {
	store := NewSemanticStore()

	store.AddEntity("E1001", "Sun", "00B2-SOL-STR-SUN:000-000-000-001")
	store.AddEntity("E1002", "Earth", "00B3-SOL-SYS-ERT:000-000-000-001")
	store.AddEntity("E1003", "NASA", "10C1-ORG-GOV-USA:NASA")

	stats := store.GetStatistics()

	if stats["entities"] != 3 {
		t.Errorf("Expected 3 entities, got %d", stats["entities"])
	}

	if stats["taxonomy_00"] != 2 {
		t.Errorf("Expected 2 natural entities, got %d", stats["taxonomy_00"])
	}

	if stats["taxonomy_10"] != 1 {
		t.Errorf("Expected 1 artificial entity, got %d", stats["taxonomy_10"])
	}
}

func TestSemanticStoreValidation(t *testing.T) {
	store := NewSemanticStore()

	store.AddEntity("E1001", "Sun", "00B2-SOL-STR-SUN:000-000-000-001")
	store.CreateAssertion("F1001", "E1001", "R1001", "E9999") // Non-existent entity

	warnings := store.ValidateStore()
	if len(warnings) == 0 {
		t.Error("Expected validation warnings, got none")
	}
}

func BenchmarkSemanticStore(b *testing.B) {
	store := NewSemanticStore()

	b.Run("AddEntity", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			id := fmt.Sprintf("E%d", i)
			label := fmt.Sprintf("Entity_%d", i)
			store.AddEntity(id, label, "00B2-SOL-STR-SUN:000-000-000-001")
		}
	})

	// Setup some entities for other benchmarks
	for i := 0; i < 1000; i++ {
		id := fmt.Sprintf("E%d", i)
		label := fmt.Sprintf("Entity_%d", i)
		store.AddEntity(id, label, "00B2-SOL-STR-SUN:000-000-000-001")
	}

	b.Run("FindByPattern", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			store.FindEntitiesByTOSIDPattern("00B")
		}
	})

	b.Run("GetEntity", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			store.GetEntity("E500")
		}
	})
}