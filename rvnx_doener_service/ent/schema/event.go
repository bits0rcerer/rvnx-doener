package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"time"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of an Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			Immutable(),
		field.Enum("event_type").NamedValues(
			"KebabShop created", "kebab_shop.created",
			"KebabShop imported", "kebab_shop.imported",
			"KebabShop updated from osm", "kebab_shop.osm_update",
		),
		field.JSON("info", map[string]interface{}{}),
	}
}

// Edges of an Event.
func (Event) Edges() []ent.Edge {
	return nil
}
