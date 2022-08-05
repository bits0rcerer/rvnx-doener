package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// UserOpinion holds the schema definition for the UserOpinion entity.
type UserOpinion struct {
	ent.Schema
}

// Mixin of the UserOpinion.
func (UserOpinion) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SonyflakIDMixin{},
	}
}

// Fields of an UserOpinion.
func (UserOpinion) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			Immutable(),
		field.String("opinion").MaxLen(1000),
		field.Bool("anonymous").Default(false),
	}
}

// Edges of an UserOpinion.
func (UserOpinion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("shop", KebabShop.Type).
			Ref("user_opinions").
			Unique(),
		edge.From("author", TwitchUser.Type).
			Ref("user_opinions").
			Unique(),
	}
}
