// Code generated by ent, DO NOT EDIT.

package event

import (
	"fmt"
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldEventType holds the string denoting the event_type field in the database.
	FieldEventType = "event_type"
	// FieldInfo holds the string denoting the info field in the database.
	FieldInfo = "info"
	// Table holds the table name of the event in the database.
	Table = "events"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
	FieldCreated,
	FieldEventType,
	FieldInfo,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "rvnx_doener_service/ent/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreated holds the default value on creation for the "created" field.
	DefaultCreated func() time.Time
)

// EventType defines the type for the "event_type" enum field.
type EventType string

// EventType values.
const (
	EventTypeKebabShopCreated        EventType = "kebab_shop.created"
	EventTypeKebabShopImported       EventType = "kebab_shop.imported"
	EventTypeKebabShopUpdatedFromOsm EventType = "kebab_shop.osm_update"
	EventTypeUserLoggedInFirstTime   EventType = "user.first_login"
	EventTypeUserLoggedIn            EventType = "user.login"
	EventTypeUserSubmittedARating    EventType = "user.submit_rating"
	EventTypeUserSubmittedAShop      EventType = "user.submit_shop"
)

func (et EventType) String() string {
	return string(et)
}

// EventTypeValidator is a validator for the "event_type" field enum values. It is called by the builders before save.
func EventTypeValidator(et EventType) error {
	switch et {
	case EventTypeKebabShopCreated, EventTypeKebabShopImported, EventTypeKebabShopUpdatedFromOsm, EventTypeUserLoggedInFirstTime, EventTypeUserLoggedIn, EventTypeUserSubmittedARating, EventTypeUserSubmittedAShop:
		return nil
	default:
		return fmt.Errorf("event: invalid enum value for event_type field: %q", et)
	}
}
