package kmac

import (
	internal_kmac "github.com/ha1tch/tosid-go/internal/kmac"
)

// Re-export types from internal package
type Statement = internal_kmac.Statement
type Entity = internal_kmac.Entity
type Relation = internal_kmac.Relation
type Assertion = internal_kmac.Assertion
type Property = internal_kmac.Property
type Event = internal_kmac.Event
type TimeReference = internal_kmac.TimeReference
type Temporal = internal_kmac.Temporal
type PartOf = internal_kmac.PartOf
type Causation = internal_kmac.Causation

// Re-export constructor functions
var (
	NewEntity        = internal_kmac.NewEntity
	NewRelation      = internal_kmac.NewRelation
	NewAssertion     = internal_kmac.NewAssertion
	NewProperty      = internal_kmac.NewProperty
	NewEvent         = internal_kmac.NewEvent
	NewTimeReference = internal_kmac.NewTimeReference
	NewTemporal      = internal_kmac.NewTemporal
	NewPartOf        = internal_kmac.NewPartOf
	NewCausation     = internal_kmac.NewCausation
)

// Re-export constants
const (
	EntityIDPrefix    = internal_kmac.EntityIDPrefix
	EventIDPrefix     = internal_kmac.EventIDPrefix
	RelationIDPrefix  = internal_kmac.RelationIDPrefix
	PropertyIDPrefix  = internal_kmac.PropertyIDPrefix
	TimeIDPrefix      = internal_kmac.TimeIDPrefix
	AssertionIDPrefix = internal_kmac.AssertionIDPrefix
)