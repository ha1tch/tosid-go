
## The KMAC Disassembler

The `disassembler.go` file implements a full KMAC disassembler with the following capabilities:

### Core Registration Functions
- `RegisterEntity`, `RegisterRelation`, `RegisterAssertion`, etc. - Register KMAC statements for analysis
- `RegisterStatement` - Generic function to register any KMAC statement
- `RegisterStatements` - Register multiple statements at once

### Disassembly Functions
- `DisassembleAssertion` - Analyze a specific assertion with all its references
- `DisassembleEntity` - Show an entity with all its relationships and properties
- `DisassembleEntityHierarchy` - Display the hierarchical structure of part-whole relationships
- `DisassembleKnowledgeGraph` - Show the complete knowledge graph in tabular format
- `DisassembleAll` - Comprehensive analysis of all registered statements

## Using the Disassembler

The `kmac_disassembler_example.go` file shows how to use the disassembler with exoplanetary data:

```go
// Create the disassembler
disassembler := kmac.NewDisassembler(nil) // nil means use stdout

// Register KMAC entities, relations, assertions, etc.
disassembler.RegisterEntity(earth)
disassembler.RegisterRelation(orbitsRelation)
disassembler.RegisterAssertion(earthOrbitsAssertion)
// ...

// Disassemble the knowledge graph in various ways
disassembler.DisassembleKnowledgeGraph()
disassembler.DisassembleEntity("E2001")  // Earth
disassembler.DisassembleAssertion("F4001")  // Earth habitability
disassembler.DisassembleEntityHierarchy("E2001")  // Earth hierarchy
```

## Sample Output

When you run the disassembler, you get structured output like:

### Knowledge Graph Overview
```
KMAC KNOWLEDGE GRAPH
==================

ENTITIES:
ID      LABEL           TOSID TYPE
--      -----           ---------
#E1001  Sol             00B2-SOL-STR-SGL:SPT-G2V-001
#E1002  Kepler-186      00B2-SOL-STR-SGL:SPT-M1V-186
#E1003  TRAPPIST-1      00B2-SOL-STR-SGL:SPT-M8V-T01
#E2001  Earth           00B3-EXO-TE-P01:RAD-1.0E-M1
#E2002  Kepler-186f     00B3-EXO-TE-P02:RAD-1.17E-M1.4
#E2003  TRAPPIST-1e     00B3-EXO-TE-P03:RAD-0.91E-M0.77
...
```

### Detailed Entity Analysis
```
ENTITY #E2001 [Earth]
  TYPE: 00B3-EXO-TE-P01:RAD-1.0E-M1
  SUBJECT OF ASSERTIONS:
    #F1001: ORBITS -> Sol
    #F2001: HAS_ORBIT -> Earth_Orbit
    #F3001: HAS_ATMOSPHERE -> Earth_Atmosphere
    #F4001: HABITABILITY_POTENTIAL -> HIGH
  OBJECT OF ASSERTIONS:
    None
  PART-OF RELATIONSHIPS:
    Contains part #E4001 [Earth_Atmosphere]
  PROPERTIES:
    diameter: 12,742 km
    mass: 5.97×10^24 kg
    has_liquid_water: true
```

### Assertion Analysis
```
ASSERTION #F4001:
  SUBJECT: #E2001 [Earth] (Entity)
  RELATION: #R1004 [HABITABILITY_POTENTIAL] type=[ASTROBIOLOGICAL_ASSESSMENT]
  OBJECT: HIGH (Literal value)
  CONFIDENCE: 1.0000 from [DIRECT_OBSERVATION]
```

### Entity Hierarchy
```
ENTITY HIERARCHY ROOTED AT #E2001 [Earth]:
  #E2001 [Earth] type=[00B3-EXO-TE-P01:RAD-1.0E-M1]
    #E4001 [Earth_Atmosphere] type=[00B3-EXO-AT-P01:ATM-N2O-H67]
```

## Benefits of the Disassembler

This approach offers several significant advantages over the previous implementation:

1. **Actual Computation**: It performs real analysis of the relationships between entities, rather than just printing predefined text.

2. **Semantic Navigation**: It allows navigation through the knowledge graph by following relationships between entities.

3. **In-Memory Representation**: It maintains a coherent in-memory representation of the knowledge that can be queried and analyzed.

4. **Multiple Views**: It provides different perspectives on the same knowledge (tables, hierarchies, detailed views).

5. **True "Disassembly"**: It actually "disassembles" KMAC statements by showing their component parts and relationships to other statements.

This implementation is much closer to what KMAC is meant to be - a semantic infrastructure for knowledge computation rather than just a notation for knowledge representation.
