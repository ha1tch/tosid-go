package kmac

import (
	"errors"
	"fmt"
	"time"
)

// TemporalState represents different temporal states
type TemporalState string

const (
	PointInTime    TemporalState = "POINT_IN_TIME"
	BeganAt        TemporalState = "BEGAN_AT"
	EndedAt        TemporalState = "ENDED_AT"
	During         TemporalState = "DURING"
	Before         TemporalState = "BEFORE"
	After          TemporalState = "AFTER"
	Simultaneous   TemporalState = "SIMULTANEOUS"
)

// Temporal represents a KMAC temporal qualification
type Temporal struct {
	assertionID string
	state       TemporalState
	timestamp   string
	startTime   *time.Time
	endTime     *time.Time
	duration    *time.Duration
}

// NewTemporal creates a new KMAC temporal qualification
func NewTemporal(assertionID string, state string, timestamp string) (*Temporal, error) {
	if assertionID == "" {
		return nil, errors.New("assertion ID cannot be empty")
	}

	if !validateIdentifier(AssertionIDPrefix, assertionID) {
		return nil, fmt.Errorf("invalid assertion ID format: %s", assertionID)
	}

	// Validate temporal state
	var temporalState TemporalState
	switch state {
	case "POINT_IN_TIME":
		temporalState = PointInTime
	case "BEGAN_AT":
		temporalState = BeganAt
	case "ENDED_AT":
		temporalState = EndedAt
	case "DURING":
		temporalState = During
	case "BEFORE":
		temporalState = Before
	case "AFTER":
		temporalState = After
	case "SIMULTANEOUS":
		temporalState = Simultaneous
	default:
		return nil, fmt.Errorf("invalid temporal state: %s", state)
	}

	return &Temporal{
		assertionID: assertionID,
		state:       temporalState,
		timestamp:   timestamp,
	}, nil
}

// NewTemporalWithDuration creates a temporal qualification with duration
func NewTemporalWithDuration(assertionID string, state string, startTime, endTime time.Time) (*Temporal, error) {
	temporal, err := NewTemporal(assertionID, state, "")
	if err != nil {
		return nil, err
	}
	
	temporal.startTime = &startTime
	temporal.endTime = &endTime
	duration := endTime.Sub(startTime)
	temporal.duration = &duration
	
	return temporal, nil
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
	return string(t.state)
}

// Timestamp returns the timestamp reference
func (t *Temporal) Timestamp() string {
	return t.timestamp
}

// SetTimestamp sets the timestamp reference
func (t *Temporal) SetTimestamp(timestamp string) {
	t.timestamp = timestamp
}

// GetStartTime returns the start time if set
func (t *Temporal) GetStartTime() *time.Time {
	return t.startTime
}

// GetEndTime returns the end time if set
func (t *Temporal) GetEndTime() *time.Time {
	return t.endTime
}

// GetDuration returns the duration if set
func (t *Temporal) GetDuration() *time.Duration {
	return t.duration
}

// SetTimeRange sets a time range for this temporal qualification
func (t *Temporal) SetTimeRange(startTime, endTime time.Time) {
	t.startTime = &startTime
	t.endTime = &endTime
	duration := endTime.Sub(startTime)
	t.duration = &duration
}

// String returns a string representation of the temporal qualification in KMAC format
func (t *Temporal) String() string {
	return fmt.Sprintf("TEMPORAL #%s state=[%s] timestamp=[%s]", 
		t.assertionID, t.state, t.timestamp)
}

// StringWithDuration returns a string representation including duration info
func (t *Temporal) StringWithDuration() string {
	base := t.String()
	if t.startTime != nil && t.endTime != nil {
		base += fmt.Sprintf(" start=[%s] end=[%s] duration=[%s]",
			t.startTime.Format(time.RFC3339),
			t.endTime.Format(time.RFC3339),
			t.duration.String())
	}
	return base
}

// IsActive checks if this temporal qualification is active at a given time
func (t *Temporal) IsActive(checkTime time.Time) bool {
	if t.startTime == nil || t.endTime == nil {
		return true // No time constraints
	}
	
	return checkTime.After(*t.startTime) && checkTime.Before(*t.endTime)
}

// OverlapsWith checks if this temporal qualification overlaps with another
func (t *Temporal) OverlapsWith(other *Temporal) bool {
	if t.startTime == nil || t.endTime == nil || other.startTime == nil || other.endTime == nil {
		return false // Can't determine overlap without time ranges
	}
	
	return t.startTime.Before(*other.endTime) && t.endTime.After(*other.startTime)
}

// Causation represents a KMAC causal relationship
type Causation struct {
	sourceID string
	targetID string
	causationType string
}

// CausationType represents different types of causation
const (
	Enablement    = "ENABLEMENT"
	Prevention    = "PREVENTION"
	Triggering    = "TRIGGERING"
	Inhibition    = "INHIBITION"
	Facilitation  = "FACILITATION"
)

// NewCausation creates a new KMAC causal relationship
func NewCausation(sourceID, targetID, causationType string) (*Causation, error) {
	if sourceID == "" || targetID == "" {
		return nil, errors.New("source ID and target ID cannot be empty")
	}
	
	// Validate causation type
	validTypes := []string{Enablement, Prevention, Triggering, Inhibition, Facilitation}
	valid := false
	for _, validType := range validTypes {
		if causationType == validType {
			valid = true
			break
		}
	}
	
	if !valid {
		return nil, fmt.Errorf("invalid causation type: %s", causationType)
	}
	
	return &Causation{
		sourceID:      sourceID,
		targetID:      targetID,
		causationType: causationType,
	}, nil
}

// SourceID returns the source assertion ID
func (c *Causation) SourceID() string {
	return c.sourceID
}

// TargetID returns the target assertion ID
func (c *Causation) TargetID() string {
	return c.targetID
}

// CausationType returns the type of causation
func (c *Causation) CausationType() string {
	return c.causationType
}

// Type returns the statement type
func (c *Causation) Type() string {
	return "CAUSATION"
}

// ID returns a unique identifier for this causation
func (c *Causation) ID() string {
	return fmt.Sprintf("CAUS_%s_%s", c.sourceID, c.targetID)
}

// String returns a string representation of the causation in KMAC format
func (c *Causation) String() string {
	return fmt.Sprintf("CAUSATION source=[#%s] target=[#%s] type=[%s]", 
		c.sourceID, c.targetID, c.causationType)
}