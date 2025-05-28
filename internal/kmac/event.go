package kmac

import (
	"errors"
	"fmt"
	"time"
)

// Event represents a KMAC event definition
type Event struct {
	id       string
	label    string
	tosidType string
	properties map[string]string
}

// NewEvent creates a new KMAC event
func NewEvent(id string, label string, tosidType string) (*Event, error) {
	if id == "" {
		return nil, errors.New("event ID cannot be empty")
	}

	if !validateIdentifier(EventIDPrefix, id) {
		return nil, fmt.Errorf("invalid event ID format: %s", id)
	}

	return &Event{
		id:       id,
		label:    label,
		tosidType: tosidType,
		properties: make(map[string]string),
	}, nil
}

// ID returns the event's identifier
func (e *Event) ID() string {
	return e.id
}

// Type returns the statement type
func (e *Event) Type() string {
	return "DEF_EVENT"
}

// Label returns the event's label
func (e *Event) Label() string {
	return e.label
}

// TOSIDType returns the event's TOSID type
func (e *Event) TOSIDType() string {
	return e.tosidType
}

// SetProperty sets a property on the event
func (e *Event) SetProperty(key, value string) {
	e.properties[key] = value
}

// GetProperty retrieves a property from the event
func (e *Event) GetProperty(key string) (string, bool) {
	val, ok := e.properties[key]
	return val, ok
}

// String returns a string representation of the event in KMAC format
func (e *Event) String() string {
	return fmt.Sprintf("DEF_EVENT #%s [%s] type=[%s]", e.id, e.label, e.tosidType)
}

// TimeReference represents a KMAC time definition
type TimeReference struct {
	id       string
	timeType string
	value    time.Time
}

// NewTimeReference creates a new KMAC time reference
func NewTimeReference(id string, timeType string, value time.Time) (*TimeReference, error) {
	if id == "" {
		return nil, errors.New("time reference ID cannot be empty")
	}

	if !validateIdentifier(TimeIDPrefix, id) {
		return nil, fmt.Errorf("invalid time reference ID format: %s", id)
	}

	return &TimeReference{
		id:       id,
		timeType: timeType,
		value:    value,
	}, nil
}

// ID returns the time reference's identifier
func (t *TimeReference) ID() string {
	return t.id
}

// Type returns the statement type
func (t *TimeReference) Type() string {
	return "DEF_TIME"
}

// TimeType returns the time reference's type
func (t *TimeReference) TimeType() string {
	return t.timeType
}

// Value returns the time reference's value
func (t *TimeReference) Value() time.Time {
	return t.value
}

// String returns a string representation of the time reference in KMAC format
func (t *TimeReference) String() string {
	return fmt.Sprintf("DEF_TIME #%s type=[%s] value=[%s]", 
		t.id, t.timeType, t.value.Format(time.RFC3339))
}

// Temporal represents a KMAC temporal qualification
type Temporal struct {
	assertionID string
	state       string
	timestamp   string
}

// NewTemporal creates a new KMAC temporal qualification
func NewTemporal(assertionID string, state string, timestamp string) (*Temporal, error) {
	if assertionID == "" {
		return nil, errors.New("assertion ID cannot be empty")
	}

	if !validateIdentifier(AssertionIDPrefix, assertionID) {
		return nil, fmt.Errorf("invalid assertion ID format: %s", assertionID)
	}

	return &Temporal{
		assertionID: assertionID,
		state:       state,
		timestamp:   timestamp,
	}, nil
}

// AssertionID returns the associated assertion's identifier
func (t *Temporal) AssertionID() string {
	return t.assertionID
}

// Type returns the statement type
func (t *Temporal) Type() string {
	return "TEMPORAL"
}

// State returns the temporal state
func (t *Temporal) State() string {
	return t.state
}

// Timestamp returns the timestamp reference
func (t *Temporal) Timestamp() string {
	return t.timestamp
}

// String returns a string representation of the temporal qualification in KMAC format
func (t *Temporal) String() string {
	return fmt.Sprintf("TEMPORAL #%s state=[%s] timestamp=[%s]", 
		t.assertionID, t.state, t.timestamp)
}

// PartOf represents a KMAC part-whole relationship
type PartOf struct {
	partID  string
	wholeID string
}

// NewPartOf creates a new KMAC part-whole relationship
func NewPartOf(partID string, wholeID string) (*PartOf, error) {
	if partID == "" || wholeID == "" {
		return nil, errors.New("part ID and whole ID cannot be empty")
	}

	return &PartOf{
		partID:  partID,
		wholeID: wholeID,
	}, nil
}

// PartID returns the part entity's identifier
func (p *PartOf) PartID() string {
	return p.partID
}

// WholeID returns the whole entity's identifier
func (p *PartOf) WholeID() string {
	return p.wholeID
}

// Type returns the statement type
func (p *PartOf) Type() string {
	return "PART_OF"
}

// ID returns an identifier for the part-of relationship
func (p *PartOf) ID() string {
	return fmt.Sprintf("PO_%s_%s", p.partID, p.wholeID)
}

// String returns a string representation of the part-whole relationship in KMAC format
func (p *PartOf) String() string {
	return fmt.Sprintf("PART_OF #%s whole=[#%s]", p.partID, p.wholeID)
}
