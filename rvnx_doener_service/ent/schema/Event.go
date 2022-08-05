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

// Mixin of the Event.
func (Event) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SonyflakIDMixin{},
	}
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
			"User logged in first time", "user.first_login",
			"User logged in", "user.login",
			"User submitted a rating", "user.submit_rating",
		),
		field.JSON("info", map[string]interface{}{}),
	}
}

// Edges of an Event.
func (Event) Edges() []ent.Edge {
	return nil
}
