# Kmacfile Format Specification

## Overview

A Kmacfile is a text file containing a sequence of instructions to build a semantic knowledge base using the TOSID-KMAC framework. Similar to how a Dockerfile builds container images, a Kmacfile builds knowledge graphs through imperative, step-by-step construction.

## File Naming Convention

- **Primary**: `Kmacfile` (no extension, capital K)
- **Alternative**: `*.kmac` for named knowledge bases (e.g., `apollo11.kmac`)
- **Development**: `Kmacfile.dev`, `Kmacfile.test` for different environments

## Basic Syntax

### Structure
```
# Comment
INSTRUCTION arguments
INSTRUCTION arguments \
    --option=value \
    --flag

# Multi-line instruction with heredoc
INSTRUCTION <<DELIMITER
content
DELIMITER
```

### Formatting Rules
- One instruction per line (continuation with `\`)
- Case-sensitive instruction names (ALL CAPS)
- Arguments separated by spaces
- Options use `--key=value` or `--flag` syntax
- Comments start with `#`
- Heredoc syntax for multi-line content

## Core Instructions

### FROM - Knowledge Base Inheritance

```
FROM base_knowledge_base [AS alias]
FROM semantic:foundation
FROM disaster-response:latest AS emergency
FROM scratch  # Empty knowledge base
```

Establishes the base knowledge base to build upon. Must be the first non-comment instruction.

**Arguments:**
- `base_knowledge_base`: Reference to existing knowledge base
- `AS alias`: Optional alias for referencing in multi-stage builds

### LABEL - Metadata Declaration

```
LABEL key=value [key2=value2 ...]
LABEL version="1.0"
LABEL domain="space_exploration" 
LABEL author="NASA Knowledge Engineering"
LABEL description="Apollo 11 mission semantic model"
```

Adds metadata to the knowledge base.

### ENTITY - Entity Definition

```
ENTITY id label tosid_type [options]
ENTITY E1001 NASA "10C1-ORG-GOV-USA:NASA"
ENTITY E1002 "Apollo 11" "10B2-SPC-MSN-APL:11" \
    --founding-year=1969 \
    --status=completed
```

Creates a new entity in the knowledge base.

**Arguments:**
- `id`: Unique entity identifier (E-prefix convention)
- `label`: Human-readable entity name (quotes if spaces)
- `tosid_type`: TOSID classification string

**Options:**
- `--{property}={value}`: Set entity properties
- `--inherit=parent_id`: Inherit properties from parent entity

### RELATION - Relation Definition

```
RELATION id label relation_type [options]
RELATION R1001 OPERATES "AGENT_OPERATION"
RELATION R1002 "LOCATED_AT" "SPATIAL_RELATIONSHIP" \
    --bidirectional=false \
    --transitive=true
```

Defines a relationship type for use in assertions.

**Arguments:**
- `id`: Unique relation identifier (R-prefix convention)
- `label`: Human-readable relation name
- `relation_type`: Semantic relation category

**Options:**
- `--bidirectional={true|false}`: Whether relation works both ways
- `--transitive={true|false}`: Whether relation is transitive
- `--symmetric={true|false}`: Whether relation is symmetric

### ASSERT - Assertion Creation

```
ASSERT id subject relation object [options]
ASSERT F1001 E1001 R1001 E1002
ASSERT F1002 E1002 AGENT V1001 \
    --confidence=0.999 \
    --source="HISTORICAL_RECORD" \
    --temporal=T1001
```

Creates a factual assertion linking entities through relations.

**Arguments:**
- `id`: Unique assertion identifier (F-prefix convention)
- `subject`: Subject entity ID or reference
- `relation`: Relation ID or built-in relation name
- `object`: Object entity ID, event ID, or literal value

**Options:**
- `--confidence=float`: Confidence level (0.0-1.0)
- `--source=string`: Evidence source description
- `--temporal=time_ref`: Link to temporal reference
- `--context=string`: Additional context information

**Built-in Relations:**
- `AGENT`: Who/what performed an action
- `LOCATION`: Where something exists/occurred
- `OCCURRED_AT`: When something happened
- `INSTANCE_OF`: Classification relationship
- `PART_OF`: Component relationship

### EVENT - Event Definition

```
EVENT id label tosid_type [options]
EVENT V1001 "Moon Landing" "11B3-EVT-HST-FST:000-000-000-001"
EVENT V1002 "Apollo 11 Launch" "11B3-EVT-TRV-LNC:000-000-000-001" \
    --duration="8 days 3 hours 18 minutes"
```

Defines an event entity with temporal characteristics.

**Arguments:**
- `id`: Unique event identifier (V-prefix convention)
- `label`: Human-readable event name
- `tosid_type`: TOSID classification for event

**Options:**
- `--duration=string`: Event duration
- `--location=entity_id`: Where event occurred
- `--{property}={value}`: Custom event properties

### TIME - Temporal Reference

```
TIME id type value [options]
TIME T1001 TIMESTAMP "1969-07-20T20:17:40Z"
TIME T1002 DURATION "8 days 3 hours"
TIME T1003 INTERVAL "1969-07-16T13:32:00Z/1969-07-24T16:50:35Z"
```

Creates temporal references for time-sensitive assertions.

**Arguments:**
- `id`: Unique time reference (T-prefix convention)
- `type`: Temporal type (TIMESTAMP, DURATION, INTERVAL, PERIODIC)
- `value`: Time value in appropriate format

**Options:**
- `--timezone=string`: Timezone specification
- `--precision=string`: Temporal precision level
- `--calendar=string`: Calendar system used

### TEMPORAL - Temporal Qualification

```
TEMPORAL assertion_id state time_reference [options]
TEMPORAL F1001 POINT_IN_TIME T1001
TEMPORAL F1002 DURING T1003
TEMPORAL F1003 BEFORE T1001 \
    --certainty=0.95
```

Adds temporal context to existing assertions.

**Arguments:**
- `assertion_id`: Assertion to qualify
- `state`: Temporal relationship (POINT_IN_TIME, DURING, BEFORE, AFTER, OVERLAPS)
- `time_reference`: Time reference ID

**Options:**
- `--certainty=float`: Temporal certainty level

### HIERARCHY - Part-Whole Relationships

```
HIERARCHY part_id whole_id [options]
HIERARCHY E4001 E2001  # Atmosphere part of Earth
HIERARCHY E1003 E1005 \
    --composition=0.25 \
    --essential=true
```

Establishes part-whole relationships between entities.

**Arguments:**
- `part_id`: Entity that is part of whole
- `whole_id`: Entity that contains part

**Options:**
- `--composition=float`: Percentage composition (0.0-1.0)
- `--essential={true|false}`: Whether part is essential to whole
- `--relationship_type=string`: Type of part-whole relationship

## Advanced Instructions

### IMPORT - External Knowledge

```
IMPORT source [AS alias] [options]
IMPORT "organizations.kmac" AS orgs
IMPORT "https://semantic.nasa.gov/space-missions.kmac" \
    --verify-checksum=sha256:abc123... \
    --trust-level=high
```

Imports entities and relations from external knowledge bases.

**Options:**
- `AS alias`: Namespace for imported entities
- `--verify-checksum=hash`: Cryptographic verification
- `--trust-level={low|medium|high}`: Trust level for imported data
- `--filter=pattern`: Import only matching entities

### INHERIT - Property Inheritance

```
INHERIT entity_id FROM parent_id [options]
INHERIT E1002 FROM missions.lunar_mission_template
INHERIT E1003 FROM space.celestial_body \
    --override=mass,diameter \
    --merge=composition
```

Applies property inheritance from parent entities.

**Options:**
- `--override=properties`: Properties to override rather than inherit
- `--merge=properties`: Properties to merge rather than replace
- `--exclude=properties`: Properties to exclude from inheritance

### VALIDATE - Consistency Checking

```
VALIDATE [scope] [options]
VALIDATE entities --strict
VALIDATE assertions --confidence-threshold=0.8
VALIDATE temporal --check-causality
VALIDATE ALL \
    --fail-on-warning=false \
    --report-format=json
```

Performs validation checks on knowledge base components.

**Scopes:**
- `entities`: Validate entity definitions and TOSID compliance
- `assertions`: Check assertion consistency and references
- `temporal`: Validate temporal relationships and causality
- `hierarchy`: Check part-whole relationship consistency
- `ALL`: Comprehensive validation

**Options:**
- `--strict`: Enable strict validation mode
- `--confidence-threshold=float`: Minimum confidence for assertions
- `--fail-on-warning={true|false}`: Whether warnings cause failure
- `--report-format={text|json|xml}`: Validation report format

### RUN - Semantic Operations

```
RUN command [arguments] [options]
RUN semantic-infer --method=forward-chaining
RUN optimize-graph --algorithm=community-detection
RUN export-summary --format=turtle --output=summary.ttl
```

Executes semantic processing operations on the knowledge base.

**Common Commands:**
- `semantic-infer`: Perform inference operations
- `optimize-graph`: Optimize knowledge graph structure
- `export-summary`: Generate knowledge base summaries
- `check-completeness`: Analyze knowledge completeness
- `find-conflicts`: Detect conflicting assertions

### COPY - Knowledge Transfer

```
COPY source dest [options]
COPY entities.json /kb/entities/
COPY --from=nasa-kb E1001,E1002 ./
```

Copies entities, relations, or assertions between knowledge bases.

**Options:**
- `--from=source_kb`: Copy from specific knowledge base
- `--transform=mapping`: Apply transformation during copy
- `--filter=criteria`: Filter copied elements

### EXPOSE - Public Interface

```
EXPOSE entity_list [AS interface_name]
EXPOSE E1001,R1001,F1001 AS public_api
EXPOSE space_missions AS mission_catalog
```

Defines the public interface of the knowledge base for external consumption.

## Multi-Stage Builds

```
# Stage 1: Base knowledge
FROM semantic:foundation AS base
ENTITY E1001 NASA "10C1-ORG-GOV-USA:NASA"
ENTITY E1002 "Space Program" "10C1-PRG-SPC-USA:NASA"

# Stage 2: Mission-specific knowledge  
FROM base AS missions
IMPORT "apollo-missions.kmac"
ASSERT F1001 E1001 OPERATES apollo.E2001

# Stage 3: Production knowledge base
FROM missions AS production
VALIDATE ALL --strict
EXPOSE space_missions AS public_interface
```

## Configuration and Environment

### Build Arguments

```
ARG CONFIDENCE_THRESHOLD=0.8
ARG VALIDATION_LEVEL=strict
ARG DOMAIN_CONTEXT=space_exploration

ENTITY E1001 NASA "10C1-ORG-GOV-USA:NASA" \
    --confidence-threshold=${CONFIDENCE_THRESHOLD}
```

### Environment Variables

```
ENV SEMANTIC_VERSION=2.1
ENV DEFAULT_NAMESPACE=nasa.gov
ENV INFERENCE_ENGINE=forward_chaining

ASSERT F1001 E1001 R1001 E1002 \
    --namespace=${DEFAULT_NAMESPACE}
```

## Error Handling and Debugging

### Conditional Instructions

```
# Only execute if condition is met
RUN [ "${VALIDATION_LEVEL}" = "strict" ] && validate-strict.sh

# Skip on specific conditions  
ASSERT F1001 E1001 R1001 E1002 || echo "Assertion failed, continuing..."
```

### Health Checks

```
HEALTHCHECK --interval=30m --timeout=10s \
    CMD semantic-validate --quick || exit 1
```

## Best Practices

### Instruction Ordering

1. `FROM` - Must be first
2. `ARG` and `ENV` - Early for configuration
3. `LABEL` - For metadata
4. `IMPORT` - External dependencies
5. `ENTITY`, `RELATION`, `EVENT` - Core definitions
6. `ASSERT`, `TEMPORAL`, `HIERARCHY` - Relationships
7. `VALIDATE` - Consistency checks
8. `EXPOSE` - Public interface

### Optimization Guidelines

- **Layer caching**: Group related instructions to maximize cache hits
- **Minimize layers**: Combine related operations when possible
- **Validate incrementally**: Use intermediate validation for large builds
- **Use multi-stage builds**: Separate build-time and runtime knowledge

### Security Considerations

- **Verify imports**: Always verify checksums for external knowledge
- **Limit scope**: Only expose necessary entities in public interface
- **Confidence tracking**: Maintain provenance for all assertions
- **Access control**: Use namespace isolation for sensitive knowledge

## Example Kmacfile

```dockerfile
# Apollo 11 Mission Knowledge Base
FROM semantic:space-exploration AS base

LABEL version="1.0"
LABEL domain="space_missions" 
LABEL mission="apollo_11"
LABEL classification="historical"

# Build arguments
ARG CONFIDENCE_LEVEL=0.999
ARG HISTORICAL_SOURCE="NASA Archives"

# Import foundational knowledge
IMPORT "organizations.kmac" AS orgs
IMPORT "celestial-bodies.kmac" AS space

# Define mission entities
ENTITY E1001 NASA "10C1-ORG-GOV-USA:NASA"
ENTITY E1002 "Apollo 11" "10B2-SPC-MSN-APL:11" \
    --mission-number=11 \
    --program=apollo

ENTITY E1003 "Lunar Surface" "00B2-CEL-MON-SFC:000-000-000-001"
ENTITY E1004 "Command Module" "10B3-VEH-SPC-CM:APL-11-COL"

# Define events
EVENT V1001 "Moon Landing" "11B3-EVT-HST-FST:000-000-000-001"
EVENT V1002 "Apollo 11 Launch" "11B3-EVT-TRV-LNC:000-000-000-001"

# Define relations
RELATION R1001 OPERATES "AGENT_OPERATION"
RELATION R1002 DESTINATION "TRAVEL_ENDPOINT"
RELATION R1003 "PART_OF" "COMPOSITION_RELATIONSHIP"

# Create temporal references
TIME T1001 TIMESTAMP "1969-07-20T20:17:40Z"
TIME T1002 TIMESTAMP "1969-07-16T13:32:00Z"

# Build assertions
ASSERT F1001 E1001 R1001 E1002 \
    --confidence=${CONFIDENCE_LEVEL} \
    --source="${HISTORICAL_SOURCE}"

ASSERT F1002 V1001 AGENT E1002 \
    --confidence=1.0 \
    --source="DIRECT_OBSERVATION"

ASSERT F1003 V1001 R1002 E1003 \
    --confidence=1.0 \
    --source="MISSION_RECORDS"

# Add temporal context
TEMPORAL F1002 POINT_IN_TIME T1001
TEMPORAL F1001 DURING T1002

# Create hierarchies
HIERARCHY E1004 E1002

# Validate construction
VALIDATE assertions --confidence-threshold=0.95
VALIDATE temporal --check-causality
VALIDATE ALL --strict

# Export public interface
EXPOSE E1001,E1002,V1001,V1002 AS apollo_11_public
EXPOSE historical_missions AS mission_catalog

# Health check
HEALTHCHECK --interval=1h \
    CMD semantic-validate --entity-count=4 --assertion-count=3
```

## File Extensions and Conventions

- **Kmacfile**: Main knowledge base definition
- **Kmacfile.{env}**: Environment-specific variants
- **.kmacignore**: Files to ignore during build context
- **.kmac**: Named knowledge base files
- **.semantic**: Alternative extension for semantic knowledge files

## Integration with Build Systems

### Command Line Interface

```bash
# Build knowledge base
kmac build .
kmac build -f apollo11.kmac -t nasa/apollo11:v1.0

# Run knowledge base
kmac run nasa/apollo11:v1.0

# Validate knowledge base
kmac validate apollo11.kmac

# Export knowledge base
kmac export nasa/apollo11:v1.0 --format=turtle > apollo11.ttl
```

### Build Context

The build context includes all files in the directory containing the Kmacfile, excluding those listed in `.kmacignore`.

## Version Compatibility

This specification targets Kmacfile format version 1.0, compatible with:
- TOSID specification v2.0+
- KMAC specification v1.5+
- Semantic reasoning engines supporting RDF/OWL