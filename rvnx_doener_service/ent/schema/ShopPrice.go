package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/jackc/pgtype"
	"time"
)

// ShopPrice holds the schema definition for the ShopPrice entity.
type ShopPrice struct {
	ent.Schema
}

// Mixin of the ShopPrice.
func (ShopPrice) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SonyflakIDMixin{},
	}
}

// Fields of an ShopPrice.
func (ShopPrice) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			Immutable(),
		field.Other("price", &pgtype.Numeric{}).
			SchemaType(map[string]string{
				dialect.Postgres: "numeric",
			}),
		field.Enum("currency").NamedValues(
			"Euro", "EUR",
			"Swiss franc", "CHF",
			"Japanese yen", "JPY",
			"Swedish krona", "SEK",
			"Danish krone", "DDK",
			"United States dollar", "USD",
			"Great British Pound", "GBP",
		),
		field.Enum("price_type").NamedValues(
			"normal kebab", "normalKebab",
			"vegetarian kebab", "vegiKebab",
		),
		field.Bool("anonymous").Default(false),
	}
}

// Edges of an ShopPrice.
func (ShopPrice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("shop", KebabShop.Type).
			Ref("user_prices").
			Unique(),
		edge.From("author", TwitchUser.Type).
			Ref("user_prices").
			Unique(),
	}
}

// Indexes of the ShopPrice.
func (ShopPrice) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("price_type"),
	}
}
