// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"rvnx_doener_service/ent/twitchuser"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// TwitchUser is the model entity for the TwitchUser schema.
type TwitchUser struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Login holds the value of the "login" field.
	Login string `json:"login,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// OauthToken holds the value of the "oauth_token" field.
	OauthToken string `json:"oauth_token,omitempty"`
	// OauthRefreshToken holds the value of the "oauth_refresh_token" field.
	OauthRefreshToken string `json:"oauth_refresh_token,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TwitchUser) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case twitchuser.FieldID:
			values[i] = new(sql.NullInt64)
		case twitchuser.FieldLogin, twitchuser.FieldEmail, twitchuser.FieldDisplayName, twitchuser.FieldOauthToken, twitchuser.FieldOauthRefreshToken:
			values[i] = new(sql.NullString)
		case twitchuser.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TwitchUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TwitchUser fields.
func (tu *TwitchUser) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case twitchuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tu.ID = int64(value.Int64)
		case twitchuser.FieldLogin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field login", values[i])
			} else if value.Valid {
				tu.Login = value.String
			}
		case twitchuser.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				tu.Email = value.String
			}
		case twitchuser.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				tu.DisplayName = value.String
			}
		case twitchuser.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tu.CreatedAt = value.Time
			}
		case twitchuser.FieldOauthToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field oauth_token", values[i])
			} else if value.Valid {
				tu.OauthToken = value.String
			}
		case twitchuser.FieldOauthRefreshToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field oauth_refresh_token", values[i])
			} else if value.Valid {
				tu.OauthRefreshToken = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TwitchUser.
// Note that you need to call TwitchUser.Unwrap() before calling this method if this TwitchUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (tu *TwitchUser) Update() *TwitchUserUpdateOne {
	return (&TwitchUserClient{config: tu.config}).UpdateOne(tu)
}

// Unwrap unwraps the TwitchUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tu *TwitchUser) Unwrap() *TwitchUser {
	_tx, ok := tu.config.driver.(*txDriver)
	if !ok {
		panic("ent: TwitchUser is not a transactional entity")
	}
	tu.config.driver = _tx.drv
	return tu
}

// String implements the fmt.Stringer.
func (tu *TwitchUser) String() string {
	var builder strings.Builder
	builder.WriteString("TwitchUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tu.ID))
	builder.WriteString("login=")
	builder.WriteString(tu.Login)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(tu.Email)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(tu.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(tu.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("oauth_token=")
	builder.WriteString(tu.OauthToken)
	builder.WriteString(", ")
	builder.WriteString("oauth_refresh_token=")
	builder.WriteString(tu.OauthRefreshToken)
	builder.WriteByte(')')
	return builder.String()
}

// TwitchUsers is a parsable slice of TwitchUser.
type TwitchUsers []*TwitchUser

func (tu TwitchUsers) config(cfg config) {
	for _i := range tu {
		tu[_i].config = cfg
	}
}
