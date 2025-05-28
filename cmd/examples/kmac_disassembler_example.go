package main

import (
	"fmt"
	"log"

	"github.com/ha1tch/tosid-go/internal/kmac"
	"github.com/ha1tch/tosid-go/pkg/semantic"
	"github.com/ha1tch/tosid-go/pkg/tosid"
)

func main() {
	fmt.Println("KMAC Disassembler Example - Exoplanetary Classification")
	fmt.Println("=====================================================")

	// Create the disassembler
	disassembler := kmac.NewDisassembler(nil) // nil means use stdout

	// Create a semantic store for testing
	store := semantic.NewSemanticStore()

	// Create star systems
	sol, err := kmac.NewEntity("E1001", "Sol", "00B2-SOL-STR-SGL:SPT-G2V-001")
	if err != nil {
		log.Fatalf("Failed to create entity: %v", err)
	}
	disassembler.RegisterEntity(sol)
	
	kepler186, err := kmac.NewEntity("E1002", "Kepler-186", "00B2-SOL-STR-SGL:SPT-M1V-186")
	if err != nil {
		log.Fatalf("Failed to create entity: %v", err)
	}
	disassembler.RegisterEntity(kepler186)
	
	trappist1, err := kmac.NewEntity("E1003", "TRAPPIST-1", "00B2-SOL-STR-SGL:SPT-M8V-T01")
	if err != nil {
		log.Fatalf("Failed to create entity: %v", err)
	}
	disassembler.RegisterEntity(trappist1)

	// Create exoplanets
	earth, err := kmac.NewEntity("E2001", "Earth", "00B3-EXO-TE-P01:RAD-1.0E-M1")
	if err != nil {
		log.Fatalf("Failed to create entity: %v", err)
	}
	earth.SetProperty("diameter", "12,742 km")
	earth.SetProperty("mass", "5.97×10^24 kg")
	earth.SetProperty("has_liquid_water", "true")
	disassembler.RegisterEntity(earth)
	
	kepler186f, err := kmac.NewEntity("E2002", "Kepler-186f", "00B3-EXO-TE-P02:RAD-1.17E-M1.4")
	if err != nil {
		log.Fatalf("Failed to create entity: %v", err)
	}
	kepler186f.SetProperty("discovery_year", "2014")
	kepler186f.SetProperty("estimated_surface_temperature", "unknown")
	disassembler.RegisterEntity(kepler186f)
	
	trappist1e, err := kmac.NewEntity("E2003", "TRAPPIST-1e", "00B3-EXO-TE-P03:RAD-0.91E-M0.77")
	if err != nil {
		log.Fatalf("Failed to create entity: %v", err)
	}
	trappist1e.SetProperty("discovery_year", "2017")
	trappist1e.SetProperty("tidally_locked", "likely")
	disassembler.RegisterEntity(trappist1e)

	// Create orbital characteristics
	earthOrbit, err := kmac.NewEntity("E3001", "Earth_Orbit", "00B3-EXO-HZ-P01:ORB-1.0A-P365")
	if err != nil {
		log.Fatalf("Failed to create entity: %v", err)
	}
	disassembler.RegisterEntity(earthOrbit)
	
	kepler186fOrbit, err := kmac.NewEntity("E3002", "Kepler-186f_Orbit", "00B3-EXO-HZ-P02:ORB-0.38A-P130")
	if err != nil {
		log.Fatalf("Failed to create entity: %v", err)
	}
	disassembler.RegisterEntity(kepler186fOrbit)
	
	trappist1eOrbit, err := kmac.NewEntity("E3003", "TRAPPIST-1e_Orbit", "00B3-EXO-HZ-P03:ORB-0.03A-P6.1")
	if err != nil {
		log.Fatalf("Failed to create entity: %v", err)
	}
	disassembler.RegisterEntity(trappist1eOrbit)

	// Create atmosphere entities
	earthAtmosphere, err := kmac.NewEntity("E4001", "Earth_Atmosphere", "00B3-EXO-AT-P01:ATM-N2O-H67")
	if err != nil {
		log.Fatalf("Failed to create entity: %v", err)
	}
	earthAtmosphere.SetProperty("composition", "78% nitrogen, 21% oxygen, 1% other gases")
	earthAtmosphere.SetProperty("pressure", "1013 hPa")
	disassembler.RegisterEntity(earthAtmosphere)
	
	kepler186fAtmosphere, err := kmac.NewEntity("E4002", "Kepler-186f_Atmosphere", "00B3-EXO-AT-P02:ATM-UKN-H00")
	if err != nil {
		log.Fatalf("Failed to create entity: %v", err)
	}
	disassembler.RegisterEntity(kepler186fAtmosphere)
	
	trappist1eAtmosphere, err := kmac.NewEntity("E4003", "TRAPPIST-1e_Atmosphere", "00B3-EXO-AT-P03:ATM-UKN-H00")
	if err != nil {
		log.Fatalf("Failed to create entity: %v", err)
	}
	disassembler.RegisterEntity(trappist1eAtmosphere)

	// Define relations
	orbitsRelation, err := kmac.NewRelation("R1001", "ORBITS", "CELESTIAL_MOTION")
	if err != nil {
		log.Fatalf("Failed to create relation: %v", err)
	}
	disassembler.RegisterRelation(orbitsRelation)
	
	hasOrbitRelation, err := kmac.NewRelation("R1002", "HAS_ORBIT", "ORBITAL_CHARACTERISTICS")
	if err != nil {
		log.Fatalf("Failed to create relation: %v", err)
	}
	disassembler.RegisterRelation(hasOrbitRelation)
	
	hasAtmosphereRelation, err := kmac.NewRelation("R1003", "HAS_ATMOSPHERE", "ATMOSPHERIC_CHARACTERISTICS")
	if err != nil {
		log.Fatalf("Failed to create relation: %v", err)
	}
	disassembler.RegisterRelation(hasAtmosphereRelation)
	
	habitabilityRelation, err := kmac.NewRelation("R1004", "HABITABILITY_POTENTIAL", "ASTROBIOLOGICAL_ASSESSMENT")
	if err != nil {
		log.Fatalf("Failed to create relation: %v", err)
	}
	disassembler.RegisterRelation(habitabilityRelation)

	// Create assertions for planetary orbits
	earthOrbitsAssertion, err := kmac.NewAssertion("F1001", "E2001", "R1001", "E1001")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	earthOrbitsAssertion.SetConfidence(1.0, "DIRECT_OBSERVATION")
	disassembler.RegisterAssertion(earthOrbitsAssertion)
	
	kepler186fOrbitsAssertion, err := kmac.NewAssertion("F1002", "E2002", "R1001", "E1002")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	kepler186fOrbitsAssertion.SetConfidence(0.999, "TRANSIT_OBSERVATIONS")
	disassembler.RegisterAssertion(kepler186fOrbitsAssertion)
	
	trappist1eOrbitsAssertion, err := kmac.NewAssertion("F1003", "E2003", "R1001", "E1003")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	trappist1eOrbitsAssertion.SetConfidence(0.999, "TRANSIT_OBSERVATIONS")
	disassembler.RegisterAssertion(trappist1eOrbitsAssertion)

	// Create assertions for orbital characteristics
	earthOrbitAssertion, err := kmac.NewAssertion("F2001", "E2001", "R1002", "E3001")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	earthOrbitAssertion.SetConfidence(1.0, "DIRECT_OBSERVATION")
	disassembler.RegisterAssertion(earthOrbitAssertion)
	
	kepler186fOrbitAssertion, err := kmac.NewAssertion("F2002", "E2002", "R1002", "E3002")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	kepler186fOrbitAssertion.SetConfidence(0.95, "ORBITAL_CALCULATIONS")
	disassembler.RegisterAssertion(kepler186fOrbitAssertion)
	
	trappist1eOrbitAssertion, err := kmac.NewAssertion("F2003", "E2003", "R1002", "E3003")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	trappist1eOrbitAssertion.SetConfidence(0.95, "ORBITAL_CALCULATIONS")
	disassembler.RegisterAssertion(trappist1eOrbitAssertion)

	// Create assertions for atmospheric characteristics
	earthAtmosphereAssertion, err := kmac.NewAssertion("F3001", "E2001", "R1003", "E4001")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	earthAtmosphereAssertion.SetConfidence(1.0, "DIRECT_OBSERVATION")
	disassembler.RegisterAssertion(earthAtmosphereAssertion)
	
	kepler186fAtmosphereAssertion, err := kmac.NewAssertion("F3002", "E2002", "R1003", "E4002")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	kepler186fAtmosphereAssertion.SetConfidence(0.7, "SPECTROSCOPIC_INFERENCE")
	disassembler.RegisterAssertion(kepler186fAtmosphereAssertion)
	
	trappist1eAtmosphereAssertion, err := kmac.NewAssertion("F3003", "E2003", "R1003", "E4003")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	trappist1eAtmosphereAssertion.SetConfidence(0.75, "SPECTROSCOPIC_INFERENCE")
	disassembler.RegisterAssertion(trappist1eAtmosphereAssertion)

	// Create assertions for habitability potential
	earthHabitabilityAssertion, err := kmac.NewAssertion("F4001", "E2001", "R1004", "HIGH")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	earthHabitabilityAssertion.SetConfidence(1.0, "DIRECT_OBSERVATION")
	disassembler.RegisterAssertion(earthHabitabilityAssertion)
	
	kepler186fHabitabilityAssertion, err := kmac.NewAssertion("F4002", "E2002", "R1004", "MEDIUM")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	kepler186fHabitabilityAssertion.SetConfidence(0.75, "SPECTROSCOPIC_ANALYSIS")
	disassembler.RegisterAssertion(kepler186fHabitabilityAssertion)
	
	trappist1eHabitabilityAssertion, err := kmac.NewAssertion("F4003", "E2003", "R1004", "MEDIUM_HIGH")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	trappist1eHabitabilityAssertion.SetConfidence(0.85, "SPECTROSCOPIC_ANALYSIS")
	disassembler.RegisterAssertion(trappist1eHabitabilityAssertion)

	// Define discovery events
	earthKnowledge, err := kmac.NewEvent("V1001", "Earth_Knowledge", "11B3-EVT-HST-FST:000-000-000-001")
	if err != nil {
		log.Fatalf("Failed to create event: %v", err)
	}
	disassembler.RegisterEvent(earthKnowledge)
	
	kepler186fDiscovery, err := kmac.NewEvent("V1002", "Kepler-186f_Discovery", "11B3-EVT-AST-DSC:000-000-000-001")
	if err != nil {
		log.Fatalf("Failed to create event: %v", err)
	}
	disassembler.RegisterEvent(kepler186fDiscovery)
	
	trappist1eDiscovery, err := kmac.NewEvent("V1003", "TRAPPIST-1e_Discovery", "11B3-EVT-AST-DSC:000-000-000-002")
	if err != nil {
		log.Fatalf("Failed to create event: %v", err)
	}
	disassembler.RegisterEvent(trappist1eDiscovery)

	// Create time references
	prehistoricTime, err := kmac.NewTimeReference("T1001", "TIMESTAMP", prehistoricDate)
	if err != nil {
		log.Fatalf("Failed to create time reference: %v", err)
	}
	disassembler.RegisterTimeReference(prehistoricTime)
	
	kepler186fDiscoveryTime, err := kmac.NewTimeReference("T1002", "TIMESTAMP", kepler186fDiscoveryDate)
	if err != nil {
		log.Fatalf("Failed to create time reference: %v", err)
	}
	disassembler.RegisterTimeReference(kepler186fDiscoveryTime)
	
	trappist1eDiscoveryTime, err := kmac.NewTimeReference("T1003", "TIMESTAMP", trappist1eDiscoveryDate)
	if err != nil {
		log.Fatalf("Failed to create time reference: %v", err)
	}
	disassembler.RegisterTimeReference(trappist1eDiscoveryTime)

	// Create temporal qualifications
	earthKnowledgeTemporal, err := kmac.NewTemporal("V1001", "POINT_IN_TIME", "#T1001")
	if err != nil {
		log.Fatalf("Failed to create temporal qualification: %v", err)
	}
	disassembler.RegisterTemporal(earthKnowledgeTemporal)
	
	kepler186fDiscoveryTemporal, err := kmac.NewTemporal("V1002", "POINT_IN_TIME", "#T1002")
	if err != nil {
		log.Fatalf("Failed to create temporal qualification: %v", err)
	}
	disassembler.RegisterTemporal(kepler186fDiscoveryTemporal)
	
	trappist1eDiscoveryTemporal, err := kmac.NewTemporal("V1003", "POINT_IN_TIME", "#T1003")
	if err != nil {
		log.Fatalf("Failed to create temporal qualification: %v", err)
	}
	disassembler.RegisterTemporal(trappist1eDiscoveryTemporal)

	// Create part-of relationships
	// The atmosphere is part of the planet
	earthAtmospherePart, err := kmac.NewPartOf("E4001", "E2001")
	if err != nil {
		log.Fatalf("Failed to create part-of relationship: %v", err)
	}
	disassembler.RegisterPartOf(earthAtmospherePart)
	
	kepler186fAtmospherePart, err := kmac.NewPartOf("E4002", "E2002")
	if err != nil {
		log.Fatalf("Failed to create part-of relationship: %v", err)
	}
	disassembler.RegisterPartOf(kepler186fAtmospherePart)
	
	trappist1eAtmospherePart, err := kmac.NewPartOf("E4003", "E2003")
	if err != nil {
		log.Fatalf("Failed to create part-of relationship: %v", err)
	}
	disassembler.RegisterPartOf(trappist1eAtmospherePart)

	// Now use the disassembler to show the knowledge in structured form
	fmt.Println("\nDISASSEMBLING KMAC KNOWLEDGE GRAPH:\n")
	
	// First show the complete knowledge graph
	disassembler.DisassembleKnowledgeGraph()
	
	// Show detailed information about Earth
	fmt.Println("\nDETAILED DISASSEMBLY OF EARTH:\n")
	disassembler.DisassembleEntity("E2001")
	
	// Show detailed information about TRAPPIST-1e
	fmt.Println("\nDETAILED DISASSEMBLY OF TRAPPIST-1e:\n")
	disassembler.DisassembleEntity("E2003")
	
	// Show the full Earth assertion about habitability
	fmt.Println("\nDETAILED DISASSEMBLY OF EARTH HABITABILITY ASSERTION:\n")
	disassembler.DisassembleAssertion("F4001")
	
	// Show the TRAPPIST-1e habitability assertion
	fmt.Println("\nDETAILED DISASSEMBLY OF TRAPPIST-1e HABITABILITY ASSERTION:\n")
	disassembler.DisassembleAssertion("F4003")
	
	// Show an entity hierarchy
	fmt.Println("\nENTITY HIERARCHY FOR EARTH:\n")
	disassembler.DisassembleEntityHierarchy("E2001")
}

// Helper variables for dates
var (
	prehistoricDate        = time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC)
	kepler186fDiscoveryDate = time.Date(2014, 4, 17, 0, 0, 0, 0, time.UTC)
	trappist1eDiscoveryDate = time.Date(2017, 2, 22, 0, 0, 0, 0, time.UTC)
)
