// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "event_type", Type: field.TypeEnum, Enums: []string{"kebab_shop.created", "kebab_shop.imported", "kebab_shop.osm_update", "user.first_login", "user.login", "user.submit_rating", "user.submit_shop"}},
		{Name: "info", Type: field.TypeJSON},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
	}
	// KebabShopsColumns holds the columns for the "kebab_shops" table.
	KebabShopsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "osm_id", Type: field.TypeInt, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "lat", Type: field.TypeFloat64},
		{Name: "lng", Type: field.TypeFloat64},
		{Name: "visible", Type: field.TypeBool, Default: true},
		{Name: "posted_anonymously", Type: field.TypeBool, Nullable: true},
	}
	// KebabShopsTable holds the schema information for the "kebab_shops" table.
	KebabShopsTable = &schema.Table{
		Name:       "kebab_shops",
		Columns:    KebabShopsColumns,
		PrimaryKey: []*schema.Column{KebabShopsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "kebabshop_lat",
				Unique:  false,
				Columns: []*schema.Column{KebabShopsColumns[4]},
			},
			{
				Name:    "kebabshop_lng",
				Unique:  false,
				Columns: []*schema.Column{KebabShopsColumns[5]},
			},
			{
				Name:    "kebabshop_name",
				Unique:  false,
				Columns: []*schema.Column{KebabShopsColumns[2]},
			},
			{
				Name:    "kebabshop_osm_id",
				Unique:  false,
				Columns: []*schema.Column{KebabShopsColumns[1]},
			},
		},
	}
	// ScoreRatingsColumns holds the columns for the "score_ratings" table.
	ScoreRatingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "score", Type: field.TypeFloat64},
		{Name: "anonymous", Type: field.TypeBool, Default: false},
		{Name: "kebab_shop_user_scores", Type: field.TypeUint64, Nullable: true},
		{Name: "twitch_user_score_ratings", Type: field.TypeInt64, Nullable: true},
	}
	// ScoreRatingsTable holds the schema information for the "score_ratings" table.
	ScoreRatingsTable = &schema.Table{
		Name:       "score_ratings",
		Columns:    ScoreRatingsColumns,
		PrimaryKey: []*schema.Column{ScoreRatingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "score_ratings_kebab_shops_user_scores",
				Columns:    []*schema.Column{ScoreRatingsColumns[4]},
				RefColumns: []*schema.Column{KebabShopsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "score_ratings_twitch_users_score_ratings",
				Columns:    []*schema.Column{ScoreRatingsColumns[5]},
				RefColumns: []*schema.Column{TwitchUsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "scorerating_kebab_shop_user_scores_twitch_user_score_ratings",
				Unique:  true,
				Columns: []*schema.Column{ScoreRatingsColumns[4], ScoreRatingsColumns[5]},
			},
		},
	}
	// ShopPricesColumns holds the columns for the "shop_prices" table.
	ShopPricesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "price", Type: field.TypeOther, SchemaType: map[string]string{"postgres": "numeric"}},
		{Name: "currency", Type: field.TypeEnum, Enums: []string{"EUR", "CHF", "JPY", "SEK", "DDK", "USD", "GBP"}},
		{Name: "price_type", Type: field.TypeEnum, Enums: []string{"normalKebab", "vegiKebab", "normalYufka", "vegiYufka", "doenerBox"}},
		{Name: "anonymous", Type: field.TypeBool, Default: false},
		{Name: "kebab_shop_user_prices", Type: field.TypeUint64, Nullable: true},
		{Name: "twitch_user_user_prices", Type: field.TypeInt64, Nullable: true},
	}
	// ShopPricesTable holds the schema information for the "shop_prices" table.
	ShopPricesTable = &schema.Table{
		Name:       "shop_prices",
		Columns:    ShopPricesColumns,
		PrimaryKey: []*schema.Column{ShopPricesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "shop_prices_kebab_shops_user_prices",
				Columns:    []*schema.Column{ShopPricesColumns[6]},
				RefColumns: []*schema.Column{KebabShopsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "shop_prices_twitch_users_user_prices",
				Columns:    []*schema.Column{ShopPricesColumns[7]},
				RefColumns: []*schema.Column{TwitchUsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "shopprice_price_type",
				Unique:  false,
				Columns: []*schema.Column{ShopPricesColumns[4]},
			},
		},
	}
	// TwitchUsersColumns holds the columns for the "twitch_users" table.
	TwitchUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "login", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "display_name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "oauth_token", Type: field.TypeString},
		{Name: "oauth_refresh_token", Type: field.TypeString},
		{Name: "activated", Type: field.TypeBool, Default: false},
		{Name: "kebab_shop_submitted_by", Type: field.TypeUint64, Nullable: true},
	}
	// TwitchUsersTable holds the schema information for the "twitch_users" table.
	TwitchUsersTable = &schema.Table{
		Name:       "twitch_users",
		Columns:    TwitchUsersColumns,
		PrimaryKey: []*schema.Column{TwitchUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "twitch_users_kebab_shops_submitted_by",
				Columns:    []*schema.Column{TwitchUsersColumns[8]},
				RefColumns: []*schema.Column{KebabShopsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "twitchuser_id",
				Unique:  false,
				Columns: []*schema.Column{TwitchUsersColumns[0]},
			},
		},
	}
	// UserOpinionsColumns holds the columns for the "user_opinions" table.
	UserOpinionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "opinion", Type: field.TypeString, Size: 1000},
		{Name: "anonymous", Type: field.TypeBool, Default: false},
		{Name: "kebab_shop_user_opinions", Type: field.TypeUint64, Nullable: true},
		{Name: "twitch_user_user_opinions", Type: field.TypeInt64, Nullable: true},
	}
	// UserOpinionsTable holds the schema information for the "user_opinions" table.
	UserOpinionsTable = &schema.Table{
		Name:       "user_opinions",
		Columns:    UserOpinionsColumns,
		PrimaryKey: []*schema.Column{UserOpinionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_opinions_kebab_shops_user_opinions",
				Columns:    []*schema.Column{UserOpinionsColumns[4]},
				RefColumns: []*schema.Column{KebabShopsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "user_opinions_twitch_users_user_opinions",
				Columns:    []*schema.Column{UserOpinionsColumns[5]},
				RefColumns: []*schema.Column{TwitchUsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EventsTable,
		KebabShopsTable,
		ScoreRatingsTable,
		ShopPricesTable,
		TwitchUsersTable,
		UserOpinionsTable,
	}
)

func init() {
	ScoreRatingsTable.ForeignKeys[0].RefTable = KebabShopsTable
	ScoreRatingsTable.ForeignKeys[1].RefTable = TwitchUsersTable
	ShopPricesTable.ForeignKeys[0].RefTable = KebabShopsTable
	ShopPricesTable.ForeignKeys[1].RefTable = TwitchUsersTable
	TwitchUsersTable.ForeignKeys[0].RefTable = KebabShopsTable
	UserOpinionsTable.ForeignKeys[0].RefTable = KebabShopsTable
	UserOpinionsTable.ForeignKeys[1].RefTable = TwitchUsersTable
}
