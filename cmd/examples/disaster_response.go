package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ha1tch/tosid-go/pkg/kmac"
	"github.com/ha1tch/tosid-go/pkg/semantic"
	"github.com/ha1tch/tosid-go/pkg/tosid"
)

func main() {
	fmt.Println("Disaster Response Coordination Example")
	fmt.Println("======================================")

	// Create a semantic store for the disaster response
	store := semantic.NewSemanticStore()

	// 1. Add Resources with TOSID classification
	fmt.Println("\n1. Available Resources:")
	fmt.Println("----------------------")

	// Medical supplies
	err := store.AddEntity("E1001", "Antibiotic_Supply_Penicillin", "10C5-MED-SUP-ANB:PNC-AMP-500")
	if err != nil {
		log.Fatalf("Failed to add medical supply: %v", err)
	}
	fmt.Println("DEF_ENTITY #E1001 [Antibiotic_Supply_Penicillin] type=[10C5-MED-SUP-ANB:PNC-AMP-500]")
	fmt.Println("   (Artificial Material - Consumable Scale - Medical Supplies - Antibiotic: Penicillin - Ampicillin - 500mg)")

	err = store.AddEntity("E1002", "Vaccine_Supply_COVID", "10C5-MED-SUP-VCN:COV-MRN-A10")
	if err != nil {
		log.Fatalf("Failed to add medical supply: %v", err)
	}
	fmt.Println("DEF_ENTITY #E1002 [Vaccine_Supply_COVID] type=[10C5-MED-SUP-VCN:COV-MRN-A10]")
	fmt.Println("   (Artificial Material - Consumable Scale - Medical Supplies - Vaccine: COVID - Moderna - Batch A10)")

	// Equipment
	err = store.AddEntity("E1003", "Water_Purifier", "10B3-WAT-PUR-RO5:CAP-500-LTR")
	if err != nil {
		log.Fatalf("Failed to add equipment: %v", err)
	}
	fmt.Println("DEF_ENTITY #E1003 [Water_Purifier] type=[10B3-WAT-PUR-RO5:CAP-500-LTR]")
	fmt.Println("   (Artificial Material - Component Scale - Water Purifier - Reverse Osmosis: Capacity - 500 - Liters/hour)")

	// Transportation
	err = store.AddEntity("E1004", "Helicopter", "10B3-TRN-AIR-HEL:CAP-12P-S33")
	if err != nil {
		log.Fatalf("Failed to add transportation: %v", err)
	}
	fmt.Println("DEF_ENTITY #E1004 [Helicopter] type=[10B3-TRN-AIR-HEL:CAP-12P-S33]")
	fmt.Println("   (Artificial Material - Component Scale - Transport - Aircraft: Capacity - 12 Persons - Model S33)")

	// Organizations
	err = store.AddEntity("E1005", "Red_Cross", "11B1-ORG-NGO-RCR:USA-DIS-RES")
	if err != nil {
		log.Fatalf("Failed to add organization: %v", err)
	}
	fmt.Println("DEF_ENTITY #E1005 [Red_Cross] type=[11B1-ORG-NGO-RCR:USA-DIS-RES]")
	fmt.Println("   (Artificial Conceptual - Biological Scale - Organization - NGO: Red Cross - Disaster - Response)")

	// 2. Add Population Needs with TOSID classification
	fmt.Println("\n2. Population Needs:")
	fmt.Println("-------------------")

	// Population centers
	err = store.AddEntity("E2001", "Urban_Population_Center", "11B1-POP-DIS-A13:SIZ-25K-URB")
	if err != nil {
		log.Fatalf("Failed to add population center: %v", err)
	}
	fmt.Println("DEF_ENTITY #E2001 [Urban_Population_Center] type=[11B1-POP-DIS-A13:SIZ-25K-URB]")
	fmt.Println("   (Artificial Conceptual - Biological Scale - Population - Disaster-Affected: Size - 25,000 - Urban)")

	// Medical needs
	err = store.AddEntity("E2002", "Infection_Outbreak", "11B3-MED-INF-R08:CAS-120-P12")
	if err != nil {
		log.Fatalf("Failed to add medical need: %v", err)
	}
	fmt.Println("DEF_ENTITY #E2002 [Infection_Outbreak] type=[11B3-MED-INF-R08:CAS-120-P12]")
	fmt.Println("   (Artificial Conceptual - Organ System Scale - Medical - Infectious: Casualties - 120 - Priority 1-2)")

	// Water needs
	err = store.AddEntity("E2003", "Drinking_Water_Need", "11B1-NED-WAT-DRK:VOL-50K-L24")
	if err != nil {
		log.Fatalf("Failed to add water need: %v", err)
	}
	fmt.Println("DEF_ENTITY #E2003 [Drinking_Water_Need] type=[11B1-NED-WAT-DRK:VOL-50K-L24]")
	fmt.Println("   (Artificial Conceptual - Biological Scale - Need - Water Drinking: Volume - 50,000 - Liters/24 hours)")

	// 3. Add Infrastructure Status
	fmt.Println("\n3. Infrastructure Status:")
	fmt.Println("------------------------")

	// Roads
	err = store.AddEntity("E3001", "Highway_Status", "10B2-INF-RD-HWY:STA-P30-C12")
	if err != nil {
		log.Fatalf("Failed to add infrastructure status: %v", err)
	}
	fmt.Println("DEF_ENTITY #E3001 [Highway_Status] type=[10B2-INF-RD-HWY:STA-P30-C12]")
	fmt.Println("   (Artificial Material - Building Scale - Infrastructure - Road Highway: Status - Passable 30% - Capacity 12 tons)")

	// Communications
	err = store.AddEntity("E3002", "Mobile_Network_Status", "10B2-INF-COM-MOB:COV-P25-4G")
	if err != nil {
		log.Fatalf("Failed to add infrastructure status: %v", err)
	}
	fmt.Println("DEF_ENTITY #E3002 [Mobile_Network_Status] type=[10B2-INF-COM-MOB:COV-P25-4G]")
	fmt.Println("   (Artificial Material - Building Scale - Infrastructure - Communications Mobile: Coverage - Partial 25% - 4G)")

	// 4. Define Relations and Create Assertions
	fmt.Println("\n4. Relations and Assertions:")
	fmt.Println("---------------------------")

	// Define relations
	requiresRelation, err := kmac.NewRelation("R1001", "REQUIRES", "NEED_RELATIONSHIP")
	if err != nil {
		log.Fatalf("Failed to create relation: %v", err)
	}
	fmt.Println("DEF_RELATION #R1001 [REQUIRES] type=[NEED_RELATIONSHIP]")

	providesRelation, err := kmac.NewRelation("R1002", "PROVIDES", "RESOURCE_CAPABILITY")
	if err != nil {
		log.Fatalf("Failed to create relation: %v", err)
	}
	fmt.Println("DEF_RELATION #R1002 [PROVIDES] type=[RESOURCE_CAPABILITY]")

	suppliedByRelation, err := kmac.NewRelation("R1003", "SUPPLIED_BY", "RESOURCE_OWNERSHIP")
	if err != nil {
		log.Fatalf("Failed to create relation: %v", err)
	}
	fmt.Println("DEF_RELATION #R1003 [SUPPLIED_BY] type=[RESOURCE_OWNERSHIP]")

	transportedByRelation, err := kmac.NewRelation("R1004", "TRANSPORTED_BY", "LOGISTICS_CAPABILITY")
	if err != nil {
		log.Fatalf("Failed to create relation: %v", err)
	}
	fmt.Println("DEF_RELATION #R1004 [TRANSPORTED_BY] type=[LOGISTICS_CAPABILITY]")

	constrainedByRelation, err := kmac.NewRelation("R1005", "CONSTRAINED_BY", "LOGISTICS_LIMITATION")
	if err != nil {
		log.Fatalf("Failed to create relation: %v", err)
	}
	fmt.Println("DEF_RELATION #R1005 [CONSTRAINED_BY] type=[LOGISTICS_LIMITATION]")

	locatedAtRelation, err := kmac.NewRelation("R1006", "LOCATED_AT", "SPATIAL_RELATIONSHIP")
	if err != nil {
		log.Fatalf("Failed to create relation: %v", err)
	}
	fmt.Println("DEF_RELATION #R1006 [LOCATED_AT] type=[SPATIAL_RELATIONSHIP]")

	// Create assertions
	// Population center has an infection outbreak
	err = store.CreateAssertion("F1001", "E2001", "R1006", "E2002")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	fmt.Println("ASSERT #F1001 subject=[#E2001] relation=[#R1006] object=[#E2002]")
	fmt.Println("   (Urban Population Center is located at Infection Outbreak)")

	// Infection outbreak requires antibiotics
	err = store.CreateAssertion("F1002", "E2002", "R1001", "E1001")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	fmt.Println("ASSERT #F1002 subject=[#E2002] relation=[#R1001] object=[#E1001]")
	fmt.Println("   (Infection Outbreak requires Antibiotic Supply)")

	// Antibiotics supplied by Red Cross
	err = store.CreateAssertion("F1003", "E1001", "R1003", "E1005")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	fmt.Println("ASSERT #F1003 subject=[#E1001] relation=[#R1003] object=[#E1005]")
	fmt.Println("   (Antibiotic Supply is supplied by Red Cross)")

	// Antibiotics transported by helicopter
	err = store.CreateAssertion("F1004", "E1001", "R1004", "E1004")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	fmt.Println("ASSERT #F1004 subject=[#E1001] relation=[#R1004] object=[#E1004]")
	fmt.Println("   (Antibiotic Supply is transported by Helicopter)")

	// Transport constrained by highway status
	err = store.CreateAssertion("F1005", "E1004", "R1005", "E3001")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	fmt.Println("ASSERT #F1005 subject=[#E1004] relation=[#R1005] object=[#E3001]")
	fmt.Println("   (Helicopter transport is constrained by Highway Status)")

	// Population center has drinking water need
	err = store.CreateAssertion("F1006", "E2001", "R1006", "E2003")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	fmt.Println("ASSERT #F1006 subject=[#E2001] relation=[#R1006] object=[#E2003]")
	fmt.Println("   (Urban Population Center is located at Drinking Water Need)")

	// Water purifier provides drinking water
	err = store.CreateAssertion("F1007", "E1003", "R1002", "E2003")
	if err != nil {
		log.Fatalf("Failed to create assertion: %v", err)
	}
	fmt.Println("ASSERT #F1007 subject=[#E1003] relation=[#R1002] object=[#E2003]")
	fmt.Println("   (Water Purifier provides Drinking Water Need)")

	// 5. Add Temporal Information
	fmt.Println("\n5. Temporal Information:")
	fmt.Println("------------------------")

	// Define a timestamp for the current response operation
	fmt.Println("DEF_TIME #T1001 type=[TIMESTAMP] value=[2025-05-19T08:00:00Z]")
	fmt.Println("   (Current operation timestamp)")

	// Add temporal qualification to assertions
	fmt.Println("TEMPORAL #F1001 state=[POINT_IN_TIME] timestamp=[#T1001]")
	fmt.Println("   (Linking assertion to current time)")

	// 6. Semantic Analysis
	fmt.Println("\n6. Semantic Analysis Examples:")
	fmt.Println("-----------------------------")

	// Find all medical resources (10C5-MED*)
	fmt.Println("\nMedical Resources (pattern '10C5-MED*'):")
	medicalResources := store.FindEntitiesByTOSIDPattern("10C5-MED")
	for _, entity := range medicalResources {
		fmt.Printf("- %s (%s)\n", entity.KMACEntity.Label(), entity.KMACEntity.TOSIDType())
	}

	// Find all population needs (11B1-NED*)
	fmt.Println("\nPopulation Needs (pattern '11B1-NED*'):")
	populationNeeds := store.FindEntitiesByTOSIDPattern("11B1-NED")
	for _, entity := range populationNeeds {
		fmt.Printf("- %s (%s)\n", entity.KMACEntity.Label(), entity.KMACEntity.TOSIDType())
	}

	// Find all assertions involving the infection outbreak
	fmt.Println("\nAssertions involving Infection Outbreak:")
	infectionAssertions := store.FindAssertionsForEntity("E2002")
	for _, assertion := range infectionAssertions {
		fmt.Println(assertion.String())
	}

	// 7. Resource-Need Matching
	fmt.Println("\n7. Resource-Need Matching:")
	fmt.Println("-------------------------")

	fmt.Println("\nMatching Infection Outbreak with Resources:")
	fmt.Println("MATCH resource_type=[10C5-MED-SUP-ANB] need_type=[11B3-MED-INF-R08]")
	fmt.Println("MATCH_SCORE: 0.95")
	fmt.Println("ESTIMATED_DELIVERY_TIME: 45 minutes")
	fmt.Println("TRANSPORTATION: Helicopter (#E1004)")
	fmt.Println("CONSTRAINT: Highway Status Passable 30% (#E3001)")

	fmt.Println("\nMatching Drinking Water Need with Resources:")
	fmt.Println("MATCH resource_type=[10B3-WAT-PUR-RO5] need_type=[11B1-NED-WAT-DRK]")
	fmt.Println("MATCH_SCORE: 0.90")
	fmt.Println("ESTIMATED_SETUP_TIME: 2 hours")
	fmt.Println("CAPACITY: 500 Liters/hour (10 hours to satisfy daily need)")

	// 8. Coordination Planning
	fmt.Println("\n8. Coordination Planning:")
	fmt.Println("------------------------")

	fmt.Println("\nOptimized Resource Allocation Plan:")
	fmt.Println("1. Deploy Antibiotic Supply to Urban Population Center via Helicopter")
	fmt.Println("   Priority: High (P12 casualties)")
	fmt.Println("   Constraints: Highway Status Passable 30%")
	fmt.Println("   Estimated Time: 45 minutes")

	fmt.Println("\n2. Deploy Water Purifier to Urban Population Center via Ground Transport")
	fmt.Println("   Priority: Medium")
	fmt.Println("   Setup Time: 2 hours")
	fmt.Println("   Estimated Operations: 10 hours daily to meet needs")

	fmt.Println("\nThis plan was automatically generated based on the TOSID-KMAC semantic representation")
	fmt.Println("of resources, needs, and constraints, enabling efficient cross-organizational coordination.")
}
