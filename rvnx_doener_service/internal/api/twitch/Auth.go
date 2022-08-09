package twitch

import (
	"crypto/rand"
	"encoding/base32"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/nicklaw5/helix/v2"
	"log"
	"net/http"
	"rvnx_doener_service/internal/services"
	"strings"
)

const (
	stateLength     = 32
	stateSessionKey = "AUTH_STATE"

	UserIDSessionKey        = "USER_ID"
	UserDisplaySessionKey   = "USER_DISPLAY"
	UserActivatedSessionKey = "USER_ACTIVATED"

	authCallbackPath = "/api/twitch/auth-callback"
)

func RouteTwitchAuth(r *gin.RouterGroup, env *services.ServiceEnvironment) {
	r.Any("/logout", logOutHandler())
	r.GET("/login", logInHandler(env.TwitchUserService))
	r.GET("/auth-callback", authCallbackHandler(env.TwitchUserService))
	r.GET("/me", meHandler(env.TwitchUserService))
}

func logOutHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		err := session.Save()
		if err != nil {
			log.Panic(err)
		}

		c.Redirect(http.StatusTemporaryRedirect, getProtocol(c.Request.Host)+c.Request.Host)
	}
}

func logInHandler(service *services.TwitchUserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		state := make([]byte, stateLength)
		_, err := rand.Read(state)
		if err != nil {
			log.Panic(err)
		}
		state = []byte(base32.StdEncoding.EncodeToString(state))

		session := sessions.Default(c)
		session.Set(stateSessionKey, string(state))
		err = session.Save()
		if err != nil {
			log.Panic(err)
		}

		helixClient, err := helix.NewClient(&helix.Options{
			ClientID:     service.GetClientID(),
			ClientSecret: service.GetClientSecret(),
			RedirectURI:  getProtocol(c.Request.Host) + strings.TrimRight(c.Request.Host, "/") + authCallbackPath,
		})
		if err != nil {
			log.Panic(err)
		}

		c.Redirect(http.StatusTemporaryRedirect,
			helixClient.GetAuthorizationURL(&helix.AuthorizationURLParams{
				ResponseType: "code",
				Scopes: []string{
					"user:read:follows",
					"user:read:subscriptions",
					"user:read:email",
				},
				State:       string(state),
				ForceVerify: false,
			}),
		)
	}
}

type authCallbackPayload struct {
	Code             string `form:"code" query:"code"`
	State            string `form:"state" query:"state"`
	Scope            string `form:"scope" query:"scope"`
	Error            string `form:"error" query:"error"`
	ErrorDescription string `form:"error_description" query:"error_description"`
}

func (p authCallbackPayload) IsFailed() bool {
	return p.Error != ""
}

func authCallbackHandler(service *services.TwitchUserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var payload authCallbackPayload
		err := c.BindQuery(&payload)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		session := sessions.Default(c)
		authState := session.Get(stateSessionKey)

		if authState != payload.State {
			c.String(http.StatusBadRequest, "state invalid")
			return
		}

		if payload.IsFailed() {
			c.String(http.StatusForbidden, "Login failed - %s (%s)", payload.Error, payload.ErrorDescription)
			return
		}

		user, err := service.FinalizeUserLogin(payload.Code,
			getProtocol(c.Request.Host)+strings.TrimRight(c.Request.Host, "/")+authCallbackPath)
		if err != nil {
			log.Panic(err)
		}

		session.Clear()
		session.Set(UserIDSessionKey, user.ID)
		session.Set(UserDisplaySessionKey, user.DisplayName)
		session.Set(UserActivatedSessionKey, user.Activated)

		err = session.Save()
		if err != nil {
			log.Panic(err)
		}

		c.Redirect(http.StatusTemporaryRedirect, getProtocol(c.Request.Host)+c.Request.Host)
	}
}

func getProtocol(host string) string {
	if strings.HasPrefix(host, "localhost") {
		return "http://"
	}
	return "https://"
}

func meHandler(service *services.TwitchUserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if id, ok := session.Get(UserIDSessionKey).(int64); ok {
			userData, exists, err := service.GetTwitchUserData(id)
			if err != nil {
				log.Panic(err)
			}
			if !exists {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"user": gin.H{
					"id":                userData.ID,
					"name":              userData.DisplayName,
					"profile_image_url": userData.ProfileImageURL,
					"activated":         session.Get(UserActivatedSessionKey),
				},
			})
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
