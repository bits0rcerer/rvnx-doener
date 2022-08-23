package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// TwitchUser holds the schema definition for the TwitchUser entity.
type TwitchUser struct {
	ent.Schema
}

// Fields of a TwitchUser.
func (TwitchUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("login"),
		field.String("email"),
		field.String("display_name"),
		field.Time("created_at"),
		field.String("oauth_token"),
		field.String("oauth_refresh_token"),
		field.Bool("activated").
			Default(false),
	}
}

// Edges of a TwitchUser.
func (TwitchUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("score_ratings", ScoreRating.Type),
		edge.To("user_prices", ShopPrice.Type),
		edge.To("user_opinions", UserOpinion.Type),
		edge.To("submitted", KebabShop.Type),
	}
}

// Indexes of a TwitchUser.
func (TwitchUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}
