// Code generated by ent, DO NOT EDIT.

package twitchuser

const (
	// Label holds the string label denoting the twitchuser type in the database.
	Label = "twitch_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLogin holds the string denoting the login field in the database.
	FieldLogin = "login"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldOauthToken holds the string denoting the oauth_token field in the database.
	FieldOauthToken = "oauth_token"
	// FieldOauthRefreshToken holds the string denoting the oauth_refresh_token field in the database.
	FieldOauthRefreshToken = "oauth_refresh_token"
	// FieldActivated holds the string denoting the activated field in the database.
	FieldActivated = "activated"
	// EdgeScoreRatings holds the string denoting the score_ratings edge name in mutations.
	EdgeScoreRatings = "score_ratings"
	// EdgeUserPrices holds the string denoting the user_prices edge name in mutations.
	EdgeUserPrices = "user_prices"
	// EdgeUserOpinions holds the string denoting the user_opinions edge name in mutations.
	EdgeUserOpinions = "user_opinions"
	// Table holds the table name of the twitchuser in the database.
	Table = "twitch_users"
	// ScoreRatingsTable is the table that holds the score_ratings relation/edge.
	ScoreRatingsTable = "score_ratings"
	// ScoreRatingsInverseTable is the table name for the ScoreRating entity.
	// It exists in this package in order to avoid circular dependency with the "scorerating" package.
	ScoreRatingsInverseTable = "score_ratings"
	// ScoreRatingsColumn is the table column denoting the score_ratings relation/edge.
	ScoreRatingsColumn = "twitch_user_score_ratings"
	// UserPricesTable is the table that holds the user_prices relation/edge.
	UserPricesTable = "shop_prices"
	// UserPricesInverseTable is the table name for the ShopPrice entity.
	// It exists in this package in order to avoid circular dependency with the "shopprice" package.
	UserPricesInverseTable = "shop_prices"
	// UserPricesColumn is the table column denoting the user_prices relation/edge.
	UserPricesColumn = "twitch_user_user_prices"
	// UserOpinionsTable is the table that holds the user_opinions relation/edge.
	UserOpinionsTable = "user_opinions"
	// UserOpinionsInverseTable is the table name for the UserOpinion entity.
	// It exists in this package in order to avoid circular dependency with the "useropinion" package.
	UserOpinionsInverseTable = "user_opinions"
	// UserOpinionsColumn is the table column denoting the user_opinions relation/edge.
	UserOpinionsColumn = "twitch_user_user_opinions"
)

// Columns holds all SQL columns for twitchuser fields.
var Columns = []string{
	FieldID,
	FieldLogin,
	FieldEmail,
	FieldDisplayName,
	FieldCreatedAt,
	FieldOauthToken,
	FieldOauthRefreshToken,
	FieldActivated,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultActivated holds the default value on creation for the "activated" field.
	DefaultActivated bool
)
