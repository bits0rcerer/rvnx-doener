package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"

	"github.com/jackc/pgtype"
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
		field.Other("point", &pgtype.Point{}).
			SchemaType(map[string]string{
				dialect.Postgres: "POINT",
			}),
	}
}

// Edges of the KebabShop.
func (KebabShop) Edges() []ent.Edge {
	return nil
}

// Indexes of the KebabShop.
func (KebabShop) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("point").
			Annotations(entsql.IndexTypes(map[string]string{
				dialect.Postgres: "GIST",
			})),
		index.Fields("name"),
		index.Fields("osm_id"),
	}
}
