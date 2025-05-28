package main

import (
	"fmt"
	"log"

	"github.com/ha1tch/tosid-go/pkg/kmac"
	"github.com/ha1tch/tosid-go/pkg/semantic"
)

func main() {
	fmt.Println("Space Program Management Example")
	fmt.Println("===============================")

	// Create a semantic store for the Jupiter atmospheric energy extraction program
	store := semantic.NewSemanticStore()

	// Add spacecraft and vehicle entities
	fmt.Println("\n1. Spacecraft and Vehicles:")
	fmt.Println("---------------------------")

	entities := []struct {
		id       string
		label    string
		tosidType string
		description string
	}{
		{"E1001", "Jupiter_Orbital_Base_X47", "10B2-SPC-JOB-X47:FNC-ORB-S25", "Jupiter Orbital Base - Function: Orbital - Series 25"},
		{"E1002", "Atmospheric_Entry_D15", "10B2-SPC-JAE-D15:FNC-PEN-S10", "Jupiter Atmospheric Entry - Function: Penetrator - Series 10"},
		{"E1003", "Ion_Propulsion_J09", "10B3-PRO-ION-J09:THR-25K-EFF", "Ion Propulsion Jupiter - Thrust: 25 Kilonewtons - Efficiency Rating"},
		{"E1004", "Aerostatic_Platform_J31", "10B2-PLT-AER-J31:FNC-LFT-P45", "Aerostatic Platform Jupiter - Function: Lift - Pressure 45"},
		{"E1005", "Energy_Extraction_J27", "10B3-ENR-KIN-J27:FNC-TRB-R35", "Energy Kinetic Jupiter - Function: Turbine - Rating 35"},
	}

	for _, e := range entities {
		err := store.AddEntity(e.id, e.label, e.tosidType)
		if err != nil {
			log.Fatalf("Failed to add entity %s: %v", e.id, err)
		}
		fmt.Printf("Added: %s - %s\n", e.label, e.description)
	}

	// Add program phases
	fmt.Println("\n2. Program Phases:")
	fmt.Println("------------------")

	phases := []struct {
		id       string
		label    string
		tosidType string
		timeRange string
	}{
		{"E2001", "Exploration_Phase", "11D1-PRG-JPH-E01:TIM-Y01-Y05", "Years 1-5"},
		{"E2002", "Development_Phase", "11D1-PRG-JPH-D02:TIM-Y06-Y12", "Years 6-12"},
		{"E2003", "Implementation_Phase", "11D1-PRG-JPH-I03:TIM-Y13-Y20", "Years 13-20"},
		{"E2004", "Operation_Phase", "11D1-PRG-JPH-O04:TIM-Y21-Y30", "Years 21-30"},
	}

	for _, p := range phases {
		err := store.AddEntity(p.id, p.label, p.tosidType)
		if err != nil {
			log.Fatalf("Failed to add phase %s: %v", p.id, err)
		}
		fmt.Printf("Phase: %s (%s)\n", p.label, p.timeRange)
	}

	// Create relationships between components
	fmt.Println("\n3. Component Relationships:")
	fmt.Println("---------------------------")

	// Propulsion system powers the orbital base
	err := store.CreateAssertion("F1001", "E1003", "POWERS", "E1001")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	fmt.Println("Ion Propulsion J09 powers Jupiter Orbital Base X47")

	// Entry vehicle uses the propulsion system
	err = store.CreateAssertion("F1002", "E1002", "USES", "E1003")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	fmt.Println("Atmospheric Entry D15 uses Ion Propulsion J09")

	// Platform supports energy extraction
	err = store.CreateAssertion("F1003", "E1004", "SUPPORTS", "E1005")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	fmt.Println("Aerostatic Platform J31 supports Energy Extraction J27")

	// Phase relationships
	err = store.CreateAssertion("F1004", "E2001", "PRECEDES", "E2002")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	fmt.Println("Exploration Phase precedes Development Phase")

	// Component development during phases
	err = store.CreateAssertion("F1005", "E1001", "DEVELOPED_DURING", "E2002")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	fmt.Println("Jupiter Orbital Base developed during Development Phase")

	// Query the semantic store
	fmt.Println("\n4. Semantic Queries:")
	fmt.Println("--------------------")

	// Find all spacecraft components
	spacecraft := store.FindEntitiesByTOSIDPattern("10B2-SPC")
	fmt.Printf("Spacecraft components (%d found):\n", len(spacecraft))
	for _, sc := range spacecraft {
		fmt.Printf("- %s (%s)\n", sc.KMACEntity.Label(), sc.KMACEntity.TOSIDType())
	}

	// Find all propulsion systems
	propulsion := store.FindEntitiesByTOSIDPattern("10B3-PRO")
	fmt.Printf("\nPropulsion systems (%d found):\n", len(propulsion))
	for _, prop := range propulsion {
		fmt.Printf("- %s (%s)\n", prop.KMACEntity.Label(), prop.KMACEntity.TOSIDType())
	}

	// Find all program phases
	programPhases := store.FindEntitiesByTOSIDPattern("11D1-PRG")
	fmt.Printf("\nProgram phases (%d found):\n", len(programPhases))
	for _, phase := range programPhases {
		fmt.Printf("- %s (%s)\n", phase.KMACEntity.Label(), phase.KMACEntity.TOSIDType())
	}

	// Show program statistics
	fmt.Println("\n5. Program Statistics:")
	fmt.Println("----------------------")

	stats := store.GetStatistics()
	fmt.Printf("Total entities: %d\n", stats["entities"])
	fmt.Printf("Total assertions: %d\n", stats["assertions"])
	fmt.Printf("Natural entities: %d\n", stats["taxonomy_00"])
	fmt.Printf("Artificial entities: %d\n", stats["taxonomy_10"])
	fmt.Printf("Conceptual entities: %d\n", stats["taxonomy_11"])

	// Validate the program structure
	fmt.Println("\n6. Program Validation:")
	fmt.Println("----------------------")

	warnings := store.ValidateStore()
	if len(warnings) == 0 {
		fmt.Println("✓ Program structure is valid - no warnings found")
	} else {
		fmt.Printf("⚠ Found %d warnings:\n", len(warnings))
		for _, warning := range warnings {
			fmt.Printf("- %s\n", warning)
		}
	}

	fmt.Println("\nSpace program management example completed successfully!")
	fmt.Println("This demonstrates how TOSID and KMAC can manage complex,")
	fmt.Println("multi-decade programs with precise semantic relationships.")
}