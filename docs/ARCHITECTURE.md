# TOSID and KMAC Architecture

This document describes the architecture of the TOSID and KMAC semantic computing system.

## Overview

The system is designed as a semantic infrastructure that combines:

- **TOSID**: Universal entity classification system
- **KMAC**: Knowledge representation and assertion system
- **Integration Layer**: Bridges between TOSID and KMAC for unified semantic processing

## Core Components

### TOSID (Taxonomic Ontological Semantic IDentification System)

TOSID provides universal entity classification with embedded semantic information.

#### Structure
```
TTN-XXX-XXX-XXX:XXX-XXX-XXX-XXX
└─┬─┘ └─────────┬─────────────────┘
  │             └─ Identifier (24 chars, 124 bits)
  └─ Taxonomy Code (2 digits) + Netmask (1 letter)
```

#### Components

**Parser** (`internal/tosid/parser.go`)
- Validates and parses TOSID strings
- Extracts semantic components
- Batch processing capabilities

**Taxonomy** (`internal/tosid/taxonomy.go`)
- Defines classification hierarchies
- Maps codes to human-readable descriptions
- Supports biological and other specialized hierarchies

**Validator** (`internal/tosid/validator.go`)
- Ensures syntactic correctness
- Validates semantic consistency
- Provides warnings for potential issues

### KMAC (Knowledge Machine Assembler Code)

KMAC provides precise representation of facts, relationships, and assertions.

#### Statement Types

**Entities** (`internal/kmac/entity.go`)
- Physical or conceptual entities with TOSID classification
- Properties and metadata support

**Relations** (`internal/kmac/relation.go`)
- Relationship types between entities
- Domain and range constraints
- Mathematical properties (symmetric, transitive, etc.)

**Assertions** (`internal/kmac/assertion.go`)
- Subject-relation-object statements
- Confidence levels and sources
- Negation support

**Events** (`internal/kmac/event.go`)
- Temporal events and processes
- Time references and timestamps

**Temporal** (`internal/kmac/temporal.go`)
- Temporal qualifications for assertions
- Duration and time range support
- Causation relationships

#### Disassembler

The KMAC disassembler (`internal/kmac/disassembler.go`) provides analysis capabilities:

- Knowledge graph visualization
- Relationship traversal
- Entity hierarchy exploration
- Assertion validation

### Integration Layer

The integration layer (`internal/integration/integration.go`) provides:

- TOSID-KMAC conversion utilities
- Semantic pattern matching
- Hierarchical relationship generation
- Cross-domain query support

### Semantic Store

The semantic store (`pkg/semantic/semantic.go`) combines TOSID and KMAC:

- Entity management with embedded classification
- Relationship tracking
- Pattern-based querying
- Validation and consistency checking

## Data Flow

```
Input Data → TOSID Classification → KMAC Representation → Semantic Store
                ↓                        ↓                     ↓
        Entity Classification    Relationship Encoding    Query Processing
```

### Classification Process

1. **Entity Identification**: Assign TOSID codes based on entity characteristics
2. **Relationship Definition**: Define KMAC relations between entities  
3. **Assertion Creation**: Create KMAC assertions with confidence levels
4. **Semantic Integration**: Store in unified semantic store for querying

### Query Processing

1. **Pattern Matching**: Use TOSID patterns to filter entities
2. **Relationship Traversal**: Follow KMAC assertions between entities
3. **Semantic Inference**: Apply reasoning based on relationship types
4. **Result Integration**: Combine results across classification boundaries

## Scalability Considerations

### Hierarchical Sharding

The system uses hierarchical sharding based on TOSID structure:

```
Shard 0: 00* (Natural entities)
Shard 1: 10* (Artificial material entities)  
Shard 2: 01* (Natural conceptual entities)
Shard 3: 11* (Artificial conceptual entities)
```

### Indexing Strategy

- **Primary Index**: TOSID pattern matching
- **Secondary Index**: KMAC relationship types
- **Tertiary Index**: Confidence levels and sources
- **Temporal Index**: Time-based queries

### Caching Architecture

- **L1 Cache**: Frequently accessed entities and relationships
- **L2 Cache**: TOSID pattern query results
- **L3 Cache**: Computed relationship paths
- **Distributed Cache**: Cross-node consistency

## Performance Characteristics

### Complexity Analysis

| Operation | TOSID | KMAC | Integrated |
|-----------|-------|------|------------|
| Entity Lookup | O(1) | O(1) | O(1) |
| Pattern Match | O(n) | N/A | O(n) |
| Relationship Query | N/A | O(k) | O(k) |
| Cross-Domain Query | N/A | N/A | O(n log n) |

Where:
- n = number of entities
- k = number of relationships per entity

### Memory Usage

- **TOSID Code**: 30 bytes (string representation)
- **KMAC Entity**: ~200 bytes (with properties)
- **KMAC Assertion**: ~100 bytes (with confidence)
- **Semantic Reference**: ~50 bytes (pointer overhead)

## Error Handling

### Validation Levels

1. **Syntactic Validation**: Format correctness
2. **Semantic Validation**: Consistency checking  
3. **Pragmatic Validation**: Real-world constraints
4. **Cross-Reference Validation**: Relationship integrity

### Error Recovery

- **Graceful Degradation**: System continues with warnings
- **Partial Results**: Return available data with error indicators
- **Automatic Correction**: Fix minor formatting issues
- **User Guidance**: Provide specific error messages and suggestions

## Extension Points

### Custom Taxonomy Extensions

TOSID supports custom taxonomy branches for domain-specific classification:

```go
// Example: Medical device taxonomy
customTaxonomy := map[string]string{
    "12": "Medical/Biomedical",
}

// Add to existing taxonomy
tosid.RegisterCustomTaxonomy(customTaxonomy)
```

### Custom KMAC Statement Types

KMAC can be extended with new statement types:

```go
type CustomStatement struct {
    id string
    customField string
}

func (cs *CustomStatement) ID() string { return cs.id }
func (cs *CustomStatement) Type() string { return "CUSTOM" }
func (cs *CustomStatement) String() string { return fmt.Sprintf("CUSTOM #%s", cs.id) }
```

### Plugin Architecture

The system supports plugins for:

- Custom parsers for domain-specific formats
- Specialized validators for industry requirements  
- Custom disassemblers for specific visualization needs
- Integration adapters for external systems

## Security Considerations

### Input Validation

- All inputs are validated before processing
- Pattern matching is bounded to prevent ReDoS attacks
- Memory usage is limited to prevent resource exhaustion

### Access Control

- Future versions will support role-based access control
- Read/write permissions per taxonomy branch
- Audit logging for all modifications

### Data Integrity

- Checksums for stored TOSID and KMAC data
- Version control for taxonomy changes
- Backup and recovery procedures

This architecture provides a foundation for semantic computing that scales from individual applications to enterprise-wide semantic infrastructure.