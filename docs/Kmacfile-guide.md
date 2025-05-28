# Kmacfile Starter Guide

*Build semantic knowledge bases like code*

## What Are Kmacfiles?

Kmacfiles let you build knowledge bases the same way you build software - with version control, collaboration, and reproducible builds. Think "Dockerfile for knowledge."

Instead of manually creating databases and hoping everyone stays synchronized, you write a simple text file that describes your knowledge, then build consistent knowledge bases from it.

## Your First Kmacfile

Let's build a simple space mission knowledge base. Create a file called `Kmacfile`:

```dockerfile
# Apollo 11 Mission Knowledge Base
FROM semantic:foundation

# Define the key players
ENTITY E1001 NASA "10C1-ORG-GOV-USA:NASA"
ENTITY E1002 "Apollo 11" "10B2-SPC-MSN-APL:11"
ENTITY E1003 "Moon" "00B2-CEL-MON-LUN:000-000-000-001"

# Define what they do
RELATION R1001 "OPERATES" "AGENT_OPERATION"
RELATION R1002 "DESTINATION" "TRAVEL_TARGET"

# State the facts
ASSERT F1001 E1001 R1001 E1002
ASSERT F1002 E1002 R1002 E1003

# Make sure everything is consistent
VALIDATE assertions
```

That's it! You've just created a semantic knowledge base that captures:
- NASA operates Apollo 11
- Apollo 11's destination was the Moon

## The Six Essential Commands

### 1. FROM - Start with existing knowledge

```dockerfile
FROM semantic:foundation
```

Every Kmacfile starts by inheriting from a base knowledge set. This gives you common concepts and prevents everyone from starting from scratch.

### 2. ENTITY - Define things in your domain

```dockerfile
ENTITY E1001 NASA "10C1-ORG-GOV-USA:NASA"
```

**Format**: `ENTITY id label tosid_classification`
- `id`: Unique identifier (start with E for entities)
- `label`: Human-readable name
- `tosid_classification`: Semantic type (more on this later)

### 3. RELATION - Define relationships

```dockerfile
RELATION R1001 "OPERATES" "AGENT_OPERATION"
```

**Format**: `RELATION id label relationship_type`
- `id`: Unique identifier (start with R for relations)
- `label`: What the relationship means
- `relationship_type`: Category of relationship

### 4. ASSERT - State facts

```dockerfile
ASSERT F1001 E1001 R1001 E1002
```

**Format**: `ASSERT id subject relation object`
- `id`: Unique identifier (start with F for facts)
- `subject`: Who/what is doing something
- `relation`: What they're doing
- `object`: Who/what they're doing it to

This creates: "NASA (E1001) operates (R1001) Apollo 11 (E1002)"

### 5. IMPORT - Include external knowledge

```dockerfile
IMPORT "space-organizations.kmac"
```

Pull in knowledge bases created by others. Perfect for collaboration.

### 6. VALIDATE - Check your work

```dockerfile
VALIDATE assertions
```

Make sure all your facts reference entities and relations that actually exist.

## Complete Example: Building a Research Lab Knowledge Base

Let's build something more realistic:

```dockerfile
# Research Lab Knowledge Base
FROM semantic:foundation

# Import existing knowledge about universities
IMPORT "universities.kmac" AS unis

# Define our lab and researchers
ENTITY E1001 "AI Research Lab" "10C1-ORG-RES-AI:LAB-001"
ENTITY E1002 "Dr. Sarah Chen" "11B1-PER-RES-AI:PHD-CS"
ENTITY E1003 "Dr. Mike Rodriguez" "11B1-PER-RES-AI:PHD-EE"
ENTITY E1004 "Neural Network Project" "11B3-PRJ-RES-NN:CNN-VIS"

# Define relationships
RELATION R1001 "EMPLOYS" "EMPLOYMENT_RELATIONSHIP"
RELATION R1002 "LEADS" "PROJECT_LEADERSHIP"
RELATION R1003 "COLLABORATES_WITH" "PEER_RELATIONSHIP"

# State the facts
ASSERT F1001 E1001 R1001 E1002
ASSERT F1002 E1001 R1001 E1003
ASSERT F1003 E1002 R1002 E1004
ASSERT F1004 E1002 R1003 E1003

# Validate everything is consistent
VALIDATE assertions --strict
```

## Building Your Knowledge Base

Save your Kmacfile and run:

```bash
kmac build .
```

This creates a queryable knowledge base you can search:

```bash
# Find all researchers
kmac query "ENTITY * * 11B1-PER-RES-*"

# Find who leads what
kmac query "ASSERT * * LEADS *"

# Find all collaborations
kmac query "ASSERT * * COLLABORATES_WITH *"
```

## Why This Approach Works

**Version Control**: Track changes to your knowledge over time
```bash
git add Kmacfile
git commit -m "Added new research collaborations"
```

**Collaboration**: Multiple people can contribute to the same knowledge base
```bash
git merge feature/new-researchers
```

**Reproducibility**: Anyone can rebuild the exact same knowledge base
```bash
kmac build --tag=v1.2.3
```

**Validation**: Catch errors before they become problems
```bash
kmac validate --confidence-threshold=0.8
```

## Common Patterns

### Adding Confidence Levels

```dockerfile
ASSERT F1001 E1001 R1001 E1002 \
    --confidence=0.95 \
    --source="HR_DATABASE"
```

### Time-Sensitive Facts

```dockerfile
ASSERT F1001 E1002 "LOCATION" "Building_A" \
    --valid-from="2024-01-01" \
    --valid-until="2024-12-31"
```

### Conditional Knowledge

```dockerfile
# Only include if we're building the public version
IF public_build THEN
    ENTITY E1005 "Public Research Project" "11B3-PRJ-RES-PUB:ML-001"
ENDIF
```

## What's Next?

Once you're comfortable with these basics, you can explore:

- **Multi-stage builds** for complex knowledge bases
- **Cryptographic verification** for trusted knowledge
- **Authority management** for organizational governance
- **Advanced validation** with custom rules
- **Integration** with databases and APIs

Check out the **Kmacfile User Manual** for the complete reference.

## Quick Reference

```dockerfile
# Essential Kmacfile structure
FROM base_knowledge

ENTITY id "label" "classification"
RELATION id "label" "type"  
ASSERT id subject relation object

IMPORT "external.kmac"
VALIDATE assertions
```

## Real-World Example: Emergency Response

```dockerfile
FROM emergency:foundation

# Resources
ENTITY E1001 "Medical Supply Cache" "10C5-MED-SUP-EMR:LOC-A"
ENTITY E1002 "Hurricane Maria" "01B2-WTH-HUR-ATL:2017-09"

# Capabilities
RELATION R1001 "CAN_RESPOND_TO" "EMERGENCY_CAPABILITY"

# Facts
ASSERT F1001 E1001 R1001 E1002 \
    --confidence=0.9 \
    --response-time="2-hours"

VALIDATE emergency-response --check-coverage
```

This creates a knowledge base that emergency coordinators can query to find resources during disasters.

## Getting Help

- **Documentation**: Full specification in the User Manual
- **Examples**: See the `examples/` directory
- **Community**: Join the semantic infrastructure community
- **Issues**: Report bugs and request features

Start simple, then grow into the complexity as you need it. That's the power of progressive discovery.

---

*Ready to build your first knowledge base? Create a Kmacfile and start experimenting!*