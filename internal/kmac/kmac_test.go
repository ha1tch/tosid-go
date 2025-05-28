package kmac

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
)

// Disassembler is a tool for displaying and analyzing KMAC statements
type Disassembler struct {
	writer        io.Writer
	indentLevel   int
	colorEnabled  bool
	entityMap     map[string]*Entity
	relationMap   map[string]*Relation
	assertionMap  map[string]*Assertion
	eventMap      map[string]*Event
	timeMap       map[string]*TimeReference
	partOfMap     map[string]*PartOf
	temporalMap   map[string]*Temporal
}

// NewDisassembler creates a new KMAC disassembler
func NewDisassembler(writer io.Writer) *Disassembler {
	if writer == nil {
		writer = os.Stdout
	}
	
	return &Disassembler{
		writer:       writer,
		indentLevel:  0,
		colorEnabled: true,
		entityMap:    make(map[string]*Entity),
		relationMap:  make(map[string]*Relation),
		assertionMap: make(map[string]*Assertion),
		eventMap:     make(map[string]*Event),
		timeMap:      make(map[string]*TimeReference),
		partOfMap:    make(map[string]*PartOf),
		temporalMap:  make(map[string]*Temporal),
	}
}

// SetColorEnabled enables or disables color output
func (d *Disassembler) SetColorEnabled(enabled bool) {
	d.colorEnabled = enabled
}

// RegisterEntity registers an entity with the disassembler
func (d *Disassembler) RegisterEntity(entity *Entity) {
	d.entityMap[entity.ID()] = entity
}

// RegisterRelation registers a relation with the disassembler
func (d *Disassembler) RegisterRelation(relation *Relation) {
	d.relationMap[relation.ID()] = relation
}

// RegisterAssertion registers an assertion with the disassembler
func (d *Disassembler) RegisterAssertion(assertion *Assertion) {
	d.assertionMap[assertion.ID()] = assertion
}

// RegisterEvent registers an event with the disassembler
func (d *Disassembler) RegisterEvent(event *Event) {
	d.eventMap[event.ID()] = event
}

// RegisterTimeReference registers a time reference with the disassembler
func (d *Disassembler) RegisterTimeReference(timeRef *TimeReference) {
	d.timeMap[timeRef.ID()] = timeRef
}

// RegisterPartOf registers a part-of relationship with the disassembler
func (d *Disassembler) RegisterPartOf(partOf *PartOf) {
	d.partOfMap[partOf.ID()] = partOf
}

// RegisterTemporal registers a temporal qualification with the disassembler
func (d *Disassembler) RegisterTemporal(temporal *Temporal) {
	d.temporalMap[temporal.AssertionID()] = temporal
}

// RegisterStatement registers any KMAC statement with the disassembler
func (d *Disassembler) RegisterStatement(stmt Statement) {
	switch s := stmt.(type) {
	case *Entity:
		d.RegisterEntity(s)
	case *Relation:
		d.RegisterRelation(s)
	case *Assertion:
		d.RegisterAssertion(s)
	case *Event:
		d.RegisterEvent(s)
	case *TimeReference:
		d.RegisterTimeReference(s)
	case *PartOf:
		d.RegisterPartOf(s)
	case *Temporal:
		d.RegisterTemporal(s)
	default:
		fmt.Fprintf(d.writer, "Unknown statement type: %T\n", s)
	}
}

// RegisterStatements registers multiple KMAC statements with the disassembler
func (d *Disassembler) RegisterStatements(statements []Statement) {
	for _, stmt := range statements {
		d.RegisterStatement(stmt)
	}
}

// DisassembleAssertion disassembles a single assertion, resolving references
func (d *Disassembler) DisassembleAssertion(assertionID string) {
	assertion, ok := d.assertionMap[assertionID]
	if !ok {
		fmt.Fprintf(d.writer, "Assertion %s not found\n", assertionID)
		return
	}
	
	// Get subject entity
	subject, subjectOk := d.entityMap[assertion.Subject()]
	if !subjectOk {
		subject, subjectOk = d.eventMap[assertion.Subject()]
	}
	
	// Get relation
	relation, relationOk := d.relationMap[assertion.Relation()]
	
	// Get object entity
	object, objectOk := d.entityMap[assertion.Object()]
	if !objectOk {
		object, objectOk = d.eventMap[assertion.Object()]
	}
	
	// Get confidence
	confidence, confidenceSource := assertion.GetConfidence()
	
	// Get temporal information
	temporal, temporalOk := d.temporalMap[assertion.ID()]
	
	// Print assertion header
	fmt.Fprintf(d.writer, "ASSERTION #%s:\n", assertion.ID())
	
	// Print subject
	fmt.Fprintf(d.writer, "  SUBJECT: ")
	if subjectOk {
		if subject.Type() == "DEF_ENTITY" {
			fmt.Fprintf(d.writer, "#%s [%s] (Entity)\n", subject.ID(), subject.(*Entity).Label())
		} else {
			fmt.Fprintf(d.writer, "#%s [%s] (Event)\n", subject.ID(), subject.(*Event).Label())
		}
	} else {
		fmt.Fprintf(d.writer, "#%s (Unknown)\n", assertion.Subject())
	}
	
	// Print relation
	fmt.Fprintf(d.writer, "  RELATION: ")
	if relationOk {
		fmt.Fprintf(d.writer, "#%s [%s] type=[%s]\n", relation.ID(), relation.Label(), relation.RelationType())
	} else if assertion.Relation() == "AGENT" {
		fmt.Fprintf(d.writer, "AGENT (Built-in relation)\n")
	} else if assertion.Relation() == "LOCATION" {
		fmt.Fprintf(d.writer, "LOCATION (Built-in relation)\n")
	} else if assertion.Relation() == "OCCURRED_AT" {
		fmt.Fprintf(d.writer, "OCCURRED_AT (Built-in relation)\n")
	} else if assertion.Relation() == "INSTANCE_OF" {
		fmt.Fprintf(d.writer, "INSTANCE_OF (Built-in relation)\n")
	} else {
		fmt.Fprintf(d.writer, "#%s (Unknown)\n", assertion.Relation())
	}
	
	// Print object
	fmt.Fprintf(d.writer, "  OBJECT: ")
	if objectOk {
		if object.Type() == "DEF_ENTITY" {
			fmt.Fprintf(d.writer, "#%s [%s] (Entity)\n", object.ID(), object.(*Entity).Label())
		} else {
			fmt.Fprintf(d.writer, "#%s [%s] (Event)\n", object.ID(), object.(*Event).Label())
		}
	} else if strings.HasPrefix(assertion.Object(), "E") || strings.HasPrefix(assertion.Object(), "V") {
		fmt.Fprintf(d.writer, "#%s (Unknown reference)\n", assertion.Object())
	} else {
		fmt.Fprintf(d.writer, "%s (Literal value)\n", assertion.Object())
	}
	
	// Print confidence if available
	if confidence > 0 {
		fmt.Fprintf(d.writer, "  CONFIDENCE: %.4f from [%s]\n", confidence, confidenceSource)
	}
	
	// Print temporal information if available
	if temporalOk {
		fmt.Fprintf(d.writer, "  TEMPORAL: %s timestamp=[%s]\n", temporal.State(), temporal.Timestamp())
	}
	
	fmt.Fprintln(d.writer)
}

// DisassembleEntity disassembles a single entity, showing related assertions
func (d *Disassembler) DisassembleEntity(entityID string) {
	entity, ok := d.entityMap[entityID]
	if !ok {
		fmt.Fprintf(d.writer, "Entity %s not found\n", entityID)
		return
	}
	
	fmt.Fprintf(d.writer, "ENTITY #%s [%s]\n", entity.ID(), entity.Label())
	fmt.Fprintf(d.writer, "  TYPE: %s\n", entity.TOSIDType())
	
	// Find all assertions where this entity is the subject
	fmt.Fprintf(d.writer, "  SUBJECT OF ASSERTIONS:\n")
	foundSubject := false
	for _, assertion := range d.assertionMap {
		if assertion.Subject() == entityID {
			foundSubject = true
			relation, relationOk := d.relationMap[assertion.Relation()]
			relationName := assertion.Relation()
			if relationOk {
				relationName = relation.Label()
			} else if assertion.Relation() == "AGENT" || assertion.Relation() == "LOCATION" || 
				  assertion.Relation() == "OCCURRED_AT" || assertion.Relation() == "INSTANCE_OF" {
				relationName = assertion.Relation()
			}
			
			objectName := assertion.Object()
			object, objectOk := d.entityMap[assertion.Object()]
			if objectOk {
				objectName = object.Label()
			}
			
			fmt.Fprintf(d.writer, "    #%s: %s -> %s\n", assertion.ID(), relationName, objectName)
		}
	}
	if !foundSubject {
		fmt.Fprintf(d.writer, "    None\n")
	}
	
	// Find all assertions where this entity is the object
	fmt.Fprintf(d.writer, "  OBJECT OF ASSERTIONS:\n")
	foundObject := false
	for _, assertion := range d.assertionMap {
		if assertion.Object() == entityID {
			foundObject = true
			relation, relationOk := d.relationMap[assertion.Relation()]
			relationName := assertion.Relation()
			if relationOk {
				relationName = relation.Label()
			} else if assertion.Relation() == "AGENT" || assertion.Relation() == "LOCATION" ||
				  assertion.Relation() == "OCCURRED_AT" || assertion.Relation() == "INSTANCE_OF" {
				relationName = assertion.Relation()
			}
			
			subjectName := assertion.Subject()
			subject, subjectOk := d.entityMap[assertion.Subject()]
			if subjectOk {
				subjectName = subject.Label()
			}
			
			fmt.Fprintf(d.writer, "    #%s: %s <- %s\n", assertion.ID(), relationName, subjectName)
		}
	}
	if !foundObject {
		fmt.Fprintf(d.writer, "    None\n")
	}
	
	// Find part-of relationships
	fmt.Fprintf(d.writer, "  PART-OF RELATIONSHIPS:\n")
	foundPartOf := false
	for _, partOf := range d.partOfMap {
		if partOf.PartID() == entityID {
			foundPartOf = true
			wholeEntity, wholeOk := d.entityMap[partOf.WholeID()]
			wholeName := partOf.WholeID()
			if wholeOk {
				wholeName = wholeEntity.Label()
			}
			fmt.Fprintf(d.writer, "    Part of #%s [%s]\n", partOf.WholeID(), wholeName)
		}
		if partOf.WholeID() == entityID {
			foundPartOf = true
			partEntity, partOk := d.entityMap[partOf.PartID()]
			partName := partOf.PartID()
			if partOk {
				partName = partEntity.Label()
			}
			fmt.Fprintf(d.writer, "    Contains part #%s [%s]\n", partOf.PartID(), partName)
		}
	}
	if !foundPartOf {
		fmt.Fprintf(d.writer, "    None\n")
	}
	
	// Print properties
	fmt.Fprintf(d.writer, "  PROPERTIES:\n")
	foundProps := false
	for key, _ := range entity.properties {
		foundProps = true
		value, _ := entity.GetProperty(key)
		fmt.Fprintf(d.writer, "    %s: %s\n", key, value)
	}
	if !foundProps {
		fmt.Fprintf(d.writer, "    None\n")
	}
	
	fmt.Fprintln(d.writer)
}

// DisassembleEntityHierarchy displays the part-of hierarchy for a given entity
func (d *Disassembler) DisassembleEntityHierarchy(rootID string) {
	entity, ok := d.entityMap[rootID]
	if !ok {
		fmt.Fprintf(d.writer, "Entity %s not found\n", rootID)
		return
	}
	
	fmt.Fprintf(d.writer, "ENTITY HIERARCHY ROOTED AT #%s [%s]:\n", entity.ID(), entity.Label())
	d.disassembleEntityHierarchyRecursive(rootID, 1)
	fmt.Fprintln(d.writer)
}

// disassembleEntityHierarchyRecursive recursively displays the part-of hierarchy
func (d *Disassembler) disassembleEntityHierarchyRecursive(entityID string, depth int) {
	entity, ok := d.entityMap[entityID]
	if !ok {
		indent := strings.Repeat("  ", depth)
		fmt.Fprintf(d.writer, "%s#%s (Unknown)\n", indent, entityID)
		return
	}
	
	indent := strings.Repeat("  ", depth)
	fmt.Fprintf(d.writer, "%s#%s [%s] type=[%s]\n", indent, entity.ID(), entity.Label(), entity.TOSIDType())
	
	// Find parts of this entity
	for _, partOf := range d.partOfMap {
		if partOf.WholeID() == entityID {
			d.disassembleEntityHierarchyRecursive(partOf.PartID(), depth+1)
		}
	}
}

// DisassembleKnowledgeGraph disassembles knowledge statements as a graph
func (d *Disassembler) DisassembleKnowledgeGraph() {
	fmt.Fprintln(d.writer, "KMAC KNOWLEDGE GRAPH")
	fmt.Fprintln(d.writer, "==================")
	
	// Create table writer for formatted output
	w := tabwriter.NewWriter(d.writer, 0, 0, 2, ' ', 0)
	
	// List all entities
	fmt.Fprintln(w, "\nENTITIES:")
	fmt.Fprintln(w, "ID\tLABEL\tTOSID TYPE")
	fmt.Fprintln(w, "--\t-----\t---------")
	var entityIDs []string
	for id := range d.entityMap {
		entityIDs = append(entityIDs, id)
	}
	sort.Strings(entityIDs)
	for _, id := range entityIDs {
		entity := d.entityMap[id]
		fmt.Fprintf(w, "#%s\t%s\t%s\n", entity.ID(), entity.Label(), entity.TOSIDType())
	}
	
	// List all events
	fmt.Fprintln(w, "\nEVENTS:")
	fmt.Fprintln(w, "ID\tLABEL\tTOSID TYPE")
	fmt.Fprintln(w, "--\t-----\t---------")
	var eventIDs []string
	for id := range d.eventMap {
		eventIDs = append(eventIDs, id)
	}
	sort.Strings(eventIDs)
	for _, id := range eventIDs {
		event := d.eventMap[id]
		fmt.Fprintf(w, "#%s\t%s\t%s\n", event.ID(), event.Label(), event.TOSIDType())
	}
	
	// List all relations
	fmt.Fprintln(w, "\nRELATIONS:")
	fmt.Fprintln(w, "ID\tLABEL\tRELATION TYPE")
	fmt.Fprintln(w, "--\t-----\t-------------")
	var relationIDs []string
	for id := range d.relationMap {
		relationIDs = append(relationIDs, id)
	}
	sort.Strings(relationIDs)
	for _, id := range relationIDs {
		relation := d.relationMap[id]
		fmt.Fprintf(w, "#%s\t%s\t%s\n", relation.ID(), relation.Label(), relation.RelationType())
	}
	
	// List all assertions
	fmt.Fprintln(w, "\nASSERTIONS:")
	fmt.Fprintln(w, "ID\tSUBJECT\tRELATION\tOBJECT\tCONFIDENCE")
	fmt.Fprintln(w, "--\t-------\t--------\t------\t----------")
	var assertionIDs []string
	for id := range d.assertionMap {
		assertionIDs = append(assertionIDs, id)
	}
	sort.Strings(assertionIDs)
	for _, id := range assertionIDs {
		assertion := d.assertionMap[id]
		
		subjectLabel := assertion.Subject()
		if subject, ok := d.entityMap[assertion.Subject()]; ok {
			subjectLabel = subject.Label()
		} else if subject, ok := d.eventMap[assertion.Subject()]; ok {
			subjectLabel = subject.Label()
		}
		
		relationLabel := assertion.Relation()
		if relation, ok := d.relationMap[assertion.Relation()]; ok {
			relationLabel = relation.Label()
		}
		
		objectLabel := assertion.Object()
		if object, ok := d.entityMap[assertion.Object()]; ok {
			objectLabel = object.Label()
		} else if object, ok := d.eventMap[assertion.Object()]; ok {
			objectLabel = object.Label()
		}
		
		confidence, source := assertion.GetConfidence()
		confidenceStr := "-"
		if confidence > 0 {
			confidenceStr = fmt.Sprintf("%.4f (%s)", confidence, source)
		}
		
		fmt.Fprintf(w, "#%s\t%s\t%s\t%s\t%s\n", 
			assertion.ID(), subjectLabel, relationLabel, objectLabel, confidenceStr)
	}
	
	// List all part-of relationships
	fmt.Fprintln(w, "\nPART-WHOLE RELATIONSHIPS:")
	fmt.Fprintln(w, "PART\tWHOLE")
	fmt.Fprintln(w, "----\t-----")
	for _, partOf := range d.partOfMap {
		partLabel := partOf.PartID()
		if part, ok := d.entityMap[partOf.PartID()]; ok {
			partLabel = part.Label()
		}
		
		wholeLabel := partOf.WholeID()
		if whole, ok := d.entityMap[partOf.WholeID()]; ok {
			wholeLabel = whole.Label()
		}
		
		fmt.Fprintf(w, "%s\t%s\n", partLabel, wholeLabel)
	}
	
	w.Flush()
}

// DisassembleAll disassembles all registered KMAC statements
func (d *Disassembler) DisassembleAll() {
	// First show the knowledge graph overview
	d.DisassembleKnowledgeGraph()
	
	// Then show detailed disassembly of each assertion
	fmt.Fprintln(d.writer, "\nDETAILED ASSERTION DISASSEMBLY")
	fmt.Fprintln(d.writer, "=============================")
	
	var assertionIDs []string
	for id := range d.assertionMap {
		assertionIDs = append(assertionIDs, id)
	}
	sort.Strings(assertionIDs)
	
	for _, id := range assertionIDs {
		d.DisassembleAssertion(id)
	}
	
	// Then show detailed disassembly of each entity
	fmt.Fprintln(d.writer, "DETAILED ENTITY DISASSEMBLY")
	fmt.Fprintln(d.writer, "==========================")
	
	var entityIDs []string
	for id := range d.entityMap {
		entityIDs = append(entityIDs, id)
	}
	sort.Strings(entityIDs)
	
	for _, id := range entityIDs {
		d.DisassembleEntity(id)
	}
}
