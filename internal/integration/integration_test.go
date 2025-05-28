package integration_test

import (
	"testing"
	"strings"

	"github.com/ha1tch/tosid-go/internal/integration"
	"github.com/ha1tch/tosid-go/pkg/kmac"
	"github.com/ha1tch/tosid-go/pkg/tosid"
)

func TestConvertTOSIDToKMACEntity(t *testing.T) {
	converter := integration.NewKMACTOSIDConverter()
	
	// Test valid conversion
	entity, err := converter.ConvertTOSIDToKMACEntity("E1001", "Sun", "00B2-SOL-STR-SUN:000-000-000-001")
	if err != nil {
		t.Fatalf("Failed to convert TOSID to KMAC entity: %v", err)
	}
	
	if entity.ID() != "E1001" {
		t.Errorf("Expected entity ID E1001, got %s", entity.ID())
	}
	
	if entity.Label() != "Sun" {
		t.Errorf("Expected entity label Sun, got %s", entity.Label())
	}
	
	if entity.TOSIDType() != "00B2-SOL-STR-SUN:000-000-000-001" {
		t.Errorf("Expected TOSID type 00B2-SOL-STR-SUN:000-000-000-001, got %s", entity.TOSIDType())
	}
	
	// Test invalid TOSID code
	_, err = converter.ConvertTOSIDToKMACEntity("E1002", "Invalid", "INVALID-TOSID-CODE")
	if err == nil {
		t.Error("Expected error for invalid TOSID code, got nil")
	}
}

func TestExtractTOSIDFromKMACEntity(t *testing.T) {
	converter := integration.NewKMACTOSIDConverter()
	
	// Create a KMAC entity with a TOSID type
	entity, err := kmac.NewEntity("E1001", "Sun", "00B2-SOL-STR-SUN:000-000-000-001")
	if err != nil {
		t.Fatalf("Failed to create KMAC entity: %v", err)
	}
	
	// Extract TOSID from the entity
	tosidObj, err := converter.ExtractTOSIDFromKMACEntity(entity)
	if err != nil {
		t.Fatalf("Failed to extract TOSID from KMAC entity: %v", err)
	}
	
	if tosidObj.TaxonomyCode != "00" {
		t.Errorf("Expected taxonomy code 00, got %s", tosidObj.TaxonomyCode)
	}
	
	if tosidObj.NetmaskIndicator != "B" {
		t.Errorf("Expected netmask indicator B, got %s", tosidObj.NetmaskIndicator)
	}
	
	if !strings.Contains(tosidObj.Identifier, "SOL-STR-SUN") {
		t.Errorf("Expected identifier to contain SOL-STR-SUN, got %s", tosidObj.Identifier)
	}
	
	// Test entity without TOSID type
	entityWithoutTOSID, err := kmac.NewEntity("E1002", "NoTOSID", "")
	if err != nil {
		t.Fatalf("Failed to create KMAC entity without TOSID: %v", err)
	}
	
	_, err = converter.ExtractTOSIDFromKMACEntity(entityWithoutTOSID)
	if err == nil {
		t.Error("Expected error for entity without TOSID type, got nil")
	}
}

func TestFindEntitiesByTOSIDPattern(t *testing.T) {
	converter := integration.NewKMACTOSIDConverter()
	
	// Create a set of entities with different TOSID types
	sun, _ := kmac.NewEntity("E1001", "Sun", "00B2-SOL-STR-SUN:000-000-000-001")
	earth, _ := kmac.NewEntity("E1002", "Earth", "00B3-SOL-SYS-ERT:000-000-000-001")
	mars, _ := kmac.NewEntity("E1003", "Mars", "00B3-SOL-SYS-MRS:000-000-000-001")
	nasa, _ := kmac.NewEntity("E1004", "NASA", "10C1-ORG-GOV-USA:NASA")
	
	entities := []*kmac.Entity{sun, earth, mars, nasa}
	
	// Find all celestial bodies (00B*)
	celestialBodies := converter.FindEntitiesByTOSIDPattern(entities, "00B")
	if len(celestialBodies) != 3 {
		t.Errorf("Expected 3 celestial bodies, got %d", len(celestialBodies))
	}
	
	// Find all planets (00B3*)
	planets := converter.FindEntitiesByTOSIDPattern(entities, "00B3")
	if len(planets) != 2 {
		t.Errorf("Expected 2 planets, got %d", len(planets))
	}
	
	// Find all organizations (10C*)
	organizations := converter.FindEntitiesByTOSIDPattern(entities, "10C")
	if len(organizations) != 1 {
		t.Errorf("Expected 1 organization, got %d", len(organizations))
	}
}

func TestGenerateKMACFromTOSIDHierarchy(t *testing.T) {
	converter := integration.NewKMACTOSIDConverter()
	
	// Define a TOSID hierarchy
	tosidCodes := []string{
		"00B2-SOL-STR-SUN:000-000-000-001", // Sun
		"00B3-SOL-SYS-ERT:000-000-000-001", // Earth
		"00B3-SOL-SYS-MRS:000-000-000-001", // Mars
	}
	
	labels := map[string]string{
		"00B2-SOL-STR-SUN:000-000-000-001": "Sun",
		"00B3-SOL-SYS-ERT:000-000-000-001": "Earth",
		"00B3-SOL-SYS-MRS:000-000-000-001": "Mars",
	}
	
	// Generate KMAC statements
	statements, err := converter.GenerateKMACFromTOSIDHierarchy(tosidCodes, labels)
	if err != nil {
		t.Fatalf("Failed to generate KMAC from TOSID hierarchy: %v", err)
	}
	
	// Check that we have the expected number of statements
	// We should have 3 entity statements and 0 part-of statements in this simple example
	if len(statements) != 3 {
		t.Errorf("Expected 3 statements, got %d", len(statements))
	}
	
	// Check that all entities were created
	entityCount := 0
	for _, stmt := range statements {
		if stmt.Type() == "DEF_ENTITY" {
			entityCount++
		}
	}
	
	if entityCount != 3 {
		t.Errorf("Expected 3 entities, got %d", entityCount)
	}
}

func TestExtractSemanticInfoFromTOSID(t *testing.T) {
	converter := integration.NewKMACTOSIDConverter()
	
	// Extract semantic info from the Sun's TOSID
	info, err := converter.ExtractSemanticInfoFromTOSID("00B2-SOL-STR-SUN:000-000-000-001")
	if err != nil {
		t.Fatalf("Failed to extract semantic info: %v", err)
	}
	
	// Check the extracted information
	if info["domain"] != "Celestial/Natural" {
		t.Errorf("Expected domain Celestial/Natural, got %s", info["domain"])
	}
	
	if info["type"] != "Physical/Material" {
		t.Errorf("Expected type Physical/Material, got %s", info["type"])
	}
	
	if info["scope"] != "Stellar Scale" {
		t.Errorf("Expected scope Stellar Scale, got %s", info["scope"])
	}
	
	if info["category1"] != "SOL" {
		t.Errorf("Expected category1 SOL, got %s", info["category1"])
	}
	
	if info["category2"] != "STR" {
		t.Errorf("Expected category2 STR, got %s", info["category2"])
	}
	
	if info["category3"] != "SUN" {
		t.Errorf("Expected category3 SUN, got %s", info["category3"])
	}
	
	if info["specific_identifier"] != "000-000-000-001" {
		t.Errorf("Expected specific_identifier 000-000-000-001, got %s", info["specific_identifier"])
	}
}
