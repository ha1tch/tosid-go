# Semantic Computing with TOSID and KMAC

This project provides a Go implementation of the Taxonomic Ontological Semantic IDentification System (TOSID) and Knowledge Machine Assembler Code (KMAC), along with integration between the two systems for semantic computing applications.

## Overview

TOSID and KMAC are complementary semantic systems designed to address coordination challenges in complex environments:

- **TOSID** is a universal taxonomic framework for uniquely identifying and classifying entities using structured alphanumeric codes that embed hierarchical information.
- **KMAC** is a semantic encoding system for representing facts, relationships, and assertions with machine-level precision, transforming knowledge from a storage problem to a computational problem.

Together, they form a comprehensive semantic infrastructure where TOSID provides entity classification (what things *are*) and KMAC provides relationship encoding (what things *do* and how they *relate*).

## Project Structure

```
semantic-computing/
├── cmd/
│   ├── examples/            # Examples demonstrating TOSID and KMAC functionality
├── internal/
│   ├── kmac/                # Core KMAC implementation
│   ├── tosid/               # Core TOSID implementation
│   └── integration/         # Integration between TOSID and KMAC
├── pkg/
│   ├── kmac/                # Public KMAC API
│   ├── tosid/               # Public TOSID API
│   └── semantic/            # Combined semantic computing API
```

## Core Features

### TOSID

- Parse and create TOSID codes with format `TTN-XXX-XXX-XXX:XXX-XXX-XXX-XXX`
- Extract semantic information embedded in TOSID codes
- Match TOSID codes against patterns for classification
- Check compatibility between TOSID codes

### KMAC

- Create various KMAC statement types:
  - Entities (DEF_ENTITY)
  - Events (DEF_EVENT)
  - Relations (DEF_RELATION)
  - Assertions (ASSERT)
  - Temporal qualifications (TEMPORAL)
  - Part-whole relationships (PART_OF)
- Add confidence levels to assertions
- Build complex knowledge representations

### Integration

- Convert between TOSID and KMAC representations
- Generate KMAC statements from TOSID hierarchies
- Find entities by TOSID patterns
- Extract semantic information from entity classifications

## Getting Started

### Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/semantic-computing.git
cd semantic-computing
```

Build the examples:

```bash
go build -o bin/examples cmd/examples/main.go
```

### Basic Usage

Run the examples:

```bash
./bin/examples
```

This will demonstrate:
1. Basic TOSID functionality
2. Basic KMAC functionality
3. Integration between TOSID and KMAC

## Usage Examples

### Working with TOSID

```go
// Parse a TOSID code
sunTosid, err := tosid.Parse("00B2-SOL-STR-SUN:000-000-000-001")
if err != nil {
    log.Fatalf("Failed to parse TOSID: %v", err)
}

// Get human-readable classification
fmt.Printf("Classification: %s\n", sunTosid.ClassificationDescription())

// Create a new TOSID code
earthTosid, err := tosid.Create("00", "B", "SOL-SYS-ERT:000-000-000-001")
if err != nil {
    log.Fatalf("Failed to create TOSID: %v", err)
}

// Check compatibility
fmt.Printf("Sun and Earth are compatible: %v\n", sunTosid.IsCompatibleWith(earthTosid))
```

### Working with KMAC

```go
// Create entities
nasa, err := kmac.NewEntity("E1001", "NASA", "10C1-ORG-GOV-USA:NASA")
apollo11, err := kmac.NewEntity("E1002", "APOLLO_11", "10B2-SPC-MSN-APL:11")

// Create a relation
operates, err := kmac.NewRelation("R1001", "OPERATES", "AGENT_OPERATION")

// Create an assertion
assertion, err := kmac.NewAssertion("F1001", "E1001", "R1001", "E1002")

// Set confidence
assertion.SetConfidence(0.9999, "HISTORICAL_RECORD")

// Output KMAC statements
fmt.Println(nasa.String())          // DEF_ENTITY #E1001 [NASA] type=[10C1-ORG-GOV-USA:NASA]
fmt.Println(assertion.String())     // ASSERT #F1001 subject=[#E1001] relation=[#R1001] object=[#E1002]
fmt.Println(assertion.ConfidenceString()) // CONFIDENCE #F1001 level=[0.9999] source=[HISTORICAL_RECORD]
```

### Integration Example

```go
// Create a semantic store
store := semantic.NewSemanticStore()

// Add entities with TOSID classification
err = store.AddEntity("E1001", "NASA", "10C1-ORG-GOV-USA:NASA")
err = store.AddEntity("E1002", "APOLLO_11", "10B2-SPC-MSN-APL:11")
err = store.AddEntity("E1003", "LUNAR_SURFACE", "00B2-CEL-MON-SFC:000-000-000-001")

// Create relation between entities
err = store.CreateAssertion("F1001", "E1001", "R1001", "E1002")

// Find entities by TOSID pattern
celestialBodies := store.FindEntitiesByTOSIDPattern("00B")
for _, entity := range celestialBodies {
    fmt.Printf("- %s (%s)\n", entity.KMACEntity.Label(), entity.KMACEntity.TOSIDType())
}

// Find all assertions for an entity
nasaAssertions := store.FindAssertionsForEntity("E1001")
for _, assertion := range nasaAssertions {
    fmt.Println(assertion.String())
}
```

## Real-World Applications

This implementation provides a foundation for several practical applications:

1. **Disaster Response Coordination**: Enabling precise resource-need matching across organizational boundaries
2. **Space Program Management**: Supporting multi-decade projects across organizations
3. **Exoplanetary Classification**: Providing semantically rich identification for astronomical discoveries
4. **Universal Commerce**: Creating a semantic backbone for supply chain management
5. **Knowledge Integration**: Connecting information across different domains with semantic precision

## Extending the Framework

The framework is designed to be extended in several ways:

1. **Add Domain-Specific TOSID Types**: Create specialized TOSID patterns for your domain
2. **Add Custom KMAC Statement Types**: Implement new statement types for specific relationship patterns
3. **Create Domain-Specific Semantic Processors**: Build on the semantic store for specific applications
4. **Integrate with Database Systems**: Add persistence layers for TOSID and KMAC

## Conceptual Background

The TOSID-KMAC relationship is based on the concept that:

- **TOSID** provides the "noun vocabulary" - identifying and classifying entities with embedded semantic information
- **KMAC** provides the "operational semantics" - encoding precisely how these entities relate, interact, transform, and exist in time

This approach enables both semantic precision and computational efficiency, transforming knowledge from passive storage to active computation.

## License

This project is licensed under the Apache 2.0 license - see the LICENSE file for details.
