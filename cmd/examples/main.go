package main

import (
	"fmt"
	"log"

	"github.com/ha1tch/tosid-go/pkg/kmac"
	"github.com/ha1tch/tosid-go/pkg/semantic"
	"github.com/ha1tch/tosid-go/pkg/tosid"
)

func main() {
	fmt.Println("TOSID and KMAC Examples")
	fmt.Println("======================")

	// TOSID Examples
	fmt.Println("\nTOSID Examples:")
	fmt.Println("---------------")
	
	// Example 1: Parse a TOSID code for the Sun
	sunTosid, err := tosid.Parse("00B2-SOL-STR-SUN:000-000-000-001")
	if err != nil {
		log.Fatalf("Failed to parse Sun TOSID: %v", err)
	}
	
	fmt.Printf("Sun TOSID: %s\n", sunTosid.String())
	fmt.Printf("Classification: %s\n", sunTosid.ClassificationDescription())
	
	// Example 2: Create a TOSID code for Earth
	earthTosid, err := tosid.Create("00", "B", "SOL-SYS-ERT:000-000-000-001")
	if err != nil {
		log.Fatalf("Failed to create Earth TOSID: %v", err)
	}
	
	fmt.Printf("Earth TOSID: %s\n", earthTosid.String())
	fmt.Printf("Classification: %s\n", earthTosid.ClassificationDescription())
	
	// Example 3: Check compatibility
	fmt.Printf("Sun and Earth are compatible: %v\n", sunTosid.IsCompatibleWith(earthTosid))
	
	// KMAC Examples
	fmt.Println("\nKMAC Examples:")
	fmt.Println("--------------")
	
	// Example 1: Create entities
	nasa, err := kmac.NewEntity("E1001", "NASA", "10C1-ORG-GOV-USA:NASA")
	if err != nil {
		log.Fatalf("Failed to create NASA entity: %v", err)
	}
	
	apollo11, err := kmac.NewEntity("E1002", "APOLLO_11", "10B2-SPC-MSN-APL:11")
	if err != nil {
		log.Fatalf("Failed to create Apollo 11 entity: %v", err)
	}
	
	// Example 2: Create relation
	operates, err := kmac.NewRelation("R1001", "OPERATES", "AGENT_OPERATION")
	if err != nil {
		log.Fatalf("Failed to create OPERATES relation: %v", err)
	}
	
	// Example 3: Create assertion
	assertion, err := kmac.NewAssertion("F1001", "E1001", "R1001", "E1002")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	
	// Set confidence
	assertion.SetConfidence(0.9999, "HISTORICAL_RECORD")
	
	// Output KMAC statements
	fmt.Println(nasa.String())
	fmt.Println(apollo11.String())
	fmt.Println(operates.String())
	fmt.Println(assertion.String())
	fmt.Println(assertion.ConfidenceString())
	
	// Integrated TOSID-KMAC Examples
	fmt.Println("\nIntegrated TOSID-KMAC Examples:")
	fmt.Println("-------------------------------")
	
	// Create a semantic store
	store := semantic.NewSemanticStore()
	
	// Add entities with TOSID classification
	err = store.AddEntity("E1001", "NASA", "10C1-ORG-GOV-USA:NASA")
	if err != nil {
		log.Fatalf("Failed to add NASA to store: %v", err)
	}
	
	err = store.AddEntity("E1002", "APOLLO_11", "10B2-SPC-MSN-APL:11")
	if err != nil {
		log.Fatalf("Failed to add Apollo 11 to store: %v", err)
	}
	
	err = store.AddEntity("E1003", "LUNAR_SURFACE", "00B2-CEL-MON-SFC:000-000-000-001")
	if err != nil {
		log.Fatalf("Failed to add Lunar Surface to store: %v", err)
	}
	
	// Create relation between entities
	err = store.CreateAssertion("F1001", "E1001", "R1001", "E1002")
	if err != nil {
		log.Fatalf("Failed to create assertion in store: %v", err)
	}
	
	// Find entities by TOSID pattern
	fmt.Println("\nFind all celestial bodies (pattern '00B*'):")
	celestialBodies := store.FindEntitiesByTOSIDPattern("00B")
	for _, entity := range celestialBodies {
		fmt.Printf("- %s (%s)\n", entity.KMACEntity.Label(), entity.KMACEntity.TOSIDType())
	}
	
	// Find all assertions for NASA
	fmt.Println("\nFind all assertions involving NASA:")
	nasaAssertions := store.FindAssertionsForEntity("E1001")
	for _, assertion := range nasaAssertions {
		fmt.Println(assertion.String())
	}
	
	// Example of a complex knowledge representation
	fmt.Println("\nComplex Knowledge Representation Example (Apollo 11 Mission):")
	createApollo11Example()
}

// createApollo11Example demonstrates a more complex knowledge representation for the Apollo 11 mission
func createApollo11Example() {
	// Create entities
	entities := []struct {
		id       string
		label    string
		tosidType string
	}{
		{"E1001", "NASA", "10C1-ORG-GOV-USA:NASA"},
		{"E1002", "APOLLO_11", "10B2-SPC-MSN-APL:11"},
		{"E1003", "LUNAR_SURFACE", "00B2-CEL-MON-SFC:000-000-000-001"},
		{"E1004", "MANKIND", "11B1-COL-HUM-ALL:000-000-000-001"},
		{"E1005", "MOON", "00B2-CEL-MON-LUN:000-000-000-001"},
	}
	
	// Create events
	events := []struct {
		id       string
		label    string
		tosidType string
	}{
		{"V1001", "LANDING", "11B3-EVT-TRV-LND:000-000-000-001"},
		{"V1002", "FIRST_STEPS", "11B3-EVT-HST-FST:000-000-000-001"},
	}
	
	// Create relations
	relations := []struct {
		id       string
		label    string
		relationType string
	}{
		{"R1001", "OPERATES", "AGENT_OPERATION"},
		{"R1002", "DESTINATION", "TRAVEL_ENDPOINT"},
		{"R1003", "ACHIEVED_BY", "AGENT_ACCOMPLISHMENT"},
		{"R1004", "HISTORICAL_PRECEDENCE", "TEMPORAL_FIRST"},
	}
	
	// Output the KMAC representation
	for _, e := range entities {
		entity, _ := kmac.NewEntity(e.id, e.label, e.tosidType)
		fmt.Println(entity.String())
	}
	
	for _, e := range events {
		event, _ := kmac.NewEntity(e.id, e.label, e.tosidType)
		fmt.Println(event.String())
	}
	
	for _, r := range relations {
		relation, _ := kmac.NewRelation(r.id, r.label, r.relationType)
		fmt.Println(relation.String())
	}
	
	// Create assertions
	assertions := []struct {
		id      string
		subject string
		relation string
		object  string
	}{
		{"F1001", "E1001", "R1001", "E1002"}, // NASA operates Apollo 11
		{"F1003", "V1001", "AGENT", "E1002"}, // Landing was done by Apollo 11
		{"F1004", "V1001", "R1002", "E1003"}, // Landing destination was lunar surface
		{"F1007", "V1002", "AGENT", "E1004"}, // First steps were by mankind
		{"F1008", "V1002", "LOCATION", "E1005"}, // First steps were on the moon
		{"F1009", "V1002", "R1003", "V1001"}, // First steps achieved by landing
	}
	
	for _, a := range assertions {
		assertion, _ := kmac.NewAssertion(a.id, a.subject, a.relation, a.object)
		assertion.SetConfidence(0.9999, "HISTORICAL_RECORD")
		fmt.Println(assertion.String())
		fmt.Println(assertion.ConfidenceString())
	}
	
	// Create a time reference
	fmt.Println("DEF_TIME #T1001 type=[TIMESTAMP] value=[1969-07-20T20:17:40Z]")
	
	// Create temporal qualification
	fmt.Println("TEMPORAL #F1003 state=[POINT_IN_TIME] timestamp=[#T1001]")
	
	// Create a part-of relationship
	fmt.Println("PART_OF #E1003 whole=[#E1005]")
}
