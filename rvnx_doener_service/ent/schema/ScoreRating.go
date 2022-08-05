package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"errors"
	"time"
)

const (
	maxRatingScore = 5
	minRatingScore = 0
)

// ScoreRating holds the schema definition for the ScoreRating entity.
type ScoreRating struct {
	ent.Schema
}

// Mixin of the ScoreRating.
func (ScoreRating) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SonyflakIDMixin{},
	}
}

// Fields of an ScoreRating.
func (ScoreRating) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			Immutable(),
		field.Float("score").
			Validate(func(f float64) error {
				if f > maxRatingScore {
					return errors.New("given score is too big")
				}
				if f < minRatingScore {
					return errors.New("negative score is not allowed")
				}
				return nil
			}),
		field.Bool("anonymous").Default(false),
	}
}

// Edges of an ScoreRating.
func (ScoreRating) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("shop", KebabShop.Type).
			Ref("user_scores").
			Unique(),
		edge.From("author", TwitchUser.Type).
			Ref("score_ratings").
			Unique(),
	}
}
