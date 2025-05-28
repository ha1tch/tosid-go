package main

import (
	"testing"

	"github.com/ha1tch/tosid-go/pkg/kmac"
	"github.com/ha1tch/tosid-go/pkg/semantic"
	"github.com/ha1tch/tosid-go/pkg/tosid"
)

func TestBasicTOSIDFunctionality(t *testing.T) {
	// Test TOSID parsing
	sunTosid, err := tosid.Parse("00B2-SOL-STR-SUN:000-000-000-001")
	if err != nil {
		t.Fatalf("Failed to parse Sun TOSID: %v", err)
	}

	if sunTosid.TaxonomyCode != "00" {
		t.Errorf("Expected taxonomy code 00, got %s", sunTosid.TaxonomyCode)
	}

	if sunTosid.NetmaskIndicator != "B" {
		t.Errorf("Expected netmask indicator B, got %s", sunTosid.NetmaskIndicator)
	}

	// Test TOSID creation
	earthTosid, err := tosid.Create("00", "B", "SOL-SYS-ERT:000-000-000-001")
	if err != nil {
		t.Fatalf("Failed to create Earth TOSID: %v", err)
	}

	// Test compatibility
	if !sunTosid.IsCompatibleWith(earthTosid) {
		t.Error("Expected Sun and Earth TOSIDs to be compatible")
	}

	// Test classification description
	description := sunTosid.ClassificationDescription()
	expected := "Celestial/Natural - Physical/Material - Stellar Scale"
	if description != expected {
		t.Errorf("Expected description '%s', got '%s'", expected, description)
	}
}

func TestBasicKMACFunctionality(t *testing.T) {
	// Test entity creation
	nasa, err := kmac.NewEntity("E1001", "NASA", "10C1-ORG-GOV-USA:NASA")
	if err != nil {
		t.Fatalf("Failed to create NASA entity: %v", err)
	}

	if nasa.ID() != "E1001" {
		t.Errorf("Expected entity ID E1001, got %s", nasa.ID())
	}

	if nasa.Label() != "NASA" {
		t.Errorf("Expected entity label NASA, got %s", nasa.Label())
	}

	// Test relation creation
	operates, err := kmac.NewRelation("R1001", "OPERATES", "AGENT_OPERATION")
	if err != nil {
		t.Fatalf("Failed to create OPERATES relation: %v", err)
	}

	// Test assertion creation
	apollo11, err := kmac.NewEntity("E1002", "APOLLO_11", "10B2-SPC-MSN-APL:11")
	if err != nil {
		t.Fatalf("Failed to create Apollo 11 entity: %v", err)
	}

	assertion, err := kmac.NewAssertion("F1001", "E1001", "R1001", "E1002")
	if err != nil {
		t.Fatalf("Failed to create assertion: %v", err)
	}

	// Test confidence setting
	assertion.SetConfidence(0.9999, "HISTORICAL_RECORD")
	confidence, source := assertion.GetConfidence()

	if confidence != 0.9999 {
		t.Errorf("Expected confidence 0.9999, got %f", confidence)
	}

	if source != "HISTORICAL_RECORD" {
		t.Errorf("Expected source HISTORICAL_RECORD, got %s", source)
	}
}

func TestSemanticStoreIntegration(t *testing.T) {
	// Create a semantic store
	store := semantic.NewSemanticStore()

	// Add entities
	err := store.AddEntity("E1001", "NASA", "10C1-ORG-GOV-USA:NASA")
	if err != nil {
		t.Fatalf("Failed to add NASA to store: %v", err)
	}

	err = store.AddEntity("E1002", "APOLLO_11", "10B2-SPC-MSN-APL:11")
	if err != nil {
		t.Fatalf("Failed to add Apollo 11 to store: %v", err)
	}

	err = store.AddEntity("E1003", "LUNAR_SURFACE", "00B2-CEL-MON-SFC:000-000-000-001")
	if err != nil {
		t.Fatalf("Failed to add Lunar Surface to store: %v", err)
	}

	// Test entity retrieval
	nasaEntity, err := store.GetEntity("E1001")
	if err != nil {
		t.Fatalf("Failed to retrieve NASA entity: %v", err)
	}

	if nasaEntity.KMACEntity.Label() != "NASA" {
		t.Errorf("Expected NASA label, got %s", nasaEntity.KMACEntity.Label())
	}

	// Test TOSID pattern matching
	celestialBodies := store.FindEntitiesByTOSIDPattern("00B")
	if len(celestialBodies) != 1 {
		t.Errorf("Expected 1 celestial body, got %d", len(celestialBodies))
	}

	// Test assertion creation
	err = store.CreateAssertion("F1001", "E1001", "R1001", "E1002")
	if err != nil {
		t.Fatalf("Failed to create assertion in store: %v", err)
	}

	// Test assertion retrieval
	assertions := store.FindAssertionsForEntity("E1001")
	if len(assertions) != 1 {
		t.Errorf("Expected 1 assertion for NASA, got %d", len(assertions))
	}

	// Test store statistics
	stats := store.GetStatistics()
	if stats["entities"] != 3 {
		t.Errorf("Expected 3 entities in stats, got %d", stats["entities"])
	}

	if stats["assertions"] != 1 {
		t.Errorf("Expected 1 assertion in stats, got %d", stats["assertions"])
	}
}

func TestTOSIDPatternMatching(t *testing.T) {
	sunTosid, err := tosid.Parse("00B2-SOL-STR-SUN:000-000-000-001")
	if err != nil {
		t.Fatalf("Failed to parse Sun TOSID: %v", err)
	}

	testCases := []struct {
		pattern  string
		expected bool
	}{
		{"00B", true},
		{"00B2", true},
		{"00*", true},
		{"*SOL*", true},
		{"11*", false},
		{"00C*", false},
		{"", true}, // Empty pattern matches everything
	}

	for _, tc := range testCases {
		result := sunTosid.MatchesPattern(tc.pattern)
		if result != tc.expected {
			t.Errorf("Pattern '%s' expected %v, got %v", tc.pattern, tc.expected, result)
		}
	}
}

func TestKMACConfidenceLevels(t *testing.T) {
	assertion, err := kmac.NewAssertion("F1001", "E1001", "R1001", "E1002")
	if err != nil {
		t.Fatalf("Failed to create assertion: %v", err)
	}

	// Test default confidence
	confidence, source := assertion.GetConfidence()
	if confidence != 1.0 {
		t.Errorf("Expected default confidence 1.0, got %f", confidence)
	}

	// Test setting confidence
	assertion.SetConfidence(0.75, "STATISTICAL_ANALYSIS")
	confidence, source = assertion.GetConfidence()

	if confidence != 0.75 {
		t.Errorf("Expected confidence 0.75, got %f", confidence)
	}

	if source != "STATISTICAL_ANALYSIS" {
		t.Errorf("Expected source STATISTICAL_ANALYSIS, got %s", source)
	}

	// Test confidence bounds
	assertion.SetConfidence(-0.5, "TEST") // Should be clamped to 0
	confidence, _ = assertion.GetConfidence()
	if confidence != 0.0 {
		t.Errorf("Expected confidence to be clamped to 0.0, got %f", confidence)
	}

	assertion.SetConfidence(1.5, "TEST") // Should be clamped to 1
	confidence, _ = assertion.GetConfidence()
	if confidence != 1.0 {
		t.Errorf("Expected confidence to be clamped to 1.0, got %f", confidence)
	}
}

func TestSemanticStoreValidation(t *testing.T) {
	store := semantic.NewSemanticStore()

	// Add some entities
	store.AddEntity("E1001", "NASA", "10C1-ORG-GOV-USA:NASA")
	store.AddEntity("E1002", "APOLLO_11", "10B2-SPC-MSN-APL:11")

	// Create assertion with non-existent entity
	store.CreateAssertion("F1001", "E1001", "R1001", "E9999") // E9999 doesn't exist

	// Validate store
	warnings := store.ValidateStore()
	if len(warnings) == 0 {
		t.Error("Expected validation warnings, got none")
	}

	// Check that the warning mentions the non-existent entity
	found := false
	for _, warning := range warnings {
		if strings.Contains(warning, "E9999") {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected warning about non-existent entity E9999")
	}
}