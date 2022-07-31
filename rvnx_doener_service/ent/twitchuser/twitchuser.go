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
	// Table holds the table name of the twitchuser in the database.
	Table = "twitch_users"
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
