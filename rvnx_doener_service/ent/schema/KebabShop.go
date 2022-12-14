package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// KebabShop holds the schema definition for the KebabShop entity.
type KebabShop struct {
	ent.Schema
}

// Mixin of the KebabShop.
func (KebabShop) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SonyflakIDMixin{},
	}
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
		field.Bool("visible").
			Default(true),
		field.Bool("posted_anonymously").
			Optional(),
	}
}

// Edges of the KebabShop.
func (KebabShop) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_scores", ScoreRating.Type),
		edge.To("user_prices", ShopPrice.Type),
		edge.To("user_opinions", UserOpinion.Type),
		edge.From("submitted_by", TwitchUser.Type).
			Ref("submitted"),
	}
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
