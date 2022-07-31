package schema

import (
	"entgo.io/ent"
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
	}
}

// Edges of a TwitchUser.
func (TwitchUser) Edges() []ent.Edge {
	return nil
}

// Indexes of a TwitchUser.
func (TwitchUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}