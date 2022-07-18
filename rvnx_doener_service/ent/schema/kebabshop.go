package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// KebabShop holds the schema definition for the KebabShop entity.
type KebabShop struct {
	ent.Schema
}

// Fields of the KebabShop.
func (KebabShop) Fields() []ent.Field {
	return []ent.Field{
		field.Int("osm_id").
			Nillable().
			Optional(),
		field.String("name"),
		field.Time("created").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			Immutable(),
		field.Float("lat"),
		field.Float("lng"),
	}
}

// Edges of the KebabShop.
func (KebabShop) Edges() []ent.Edge {
	return nil
}

// Indexes of the KebabShop.
func (KebabShop) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("lat"),
		index.Fields("lng"),
		index.Fields("name"),
		index.Fields("osm_id"),
	}
}
