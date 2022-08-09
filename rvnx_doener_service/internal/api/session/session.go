package session

import (
	"crypto/sha256"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

const (
	ginSessionSecretKey           = "SESSION_SECRET"
	ginSessionEncryptionSecretKey = "SESSION_ENCRYPTION_SECRET"

	SessionCookieName = "rvxn-doener-session"
)

func InitSessions(g *gin.RouterGroup) {
	secret := os.Getenv(ginSessionSecretKey)
	if secret == "" {
		log.Panic("$" + ginSessionSecretKey + " is not set")
	}

	sha2 := sha256.New()
	sha2.Write([]byte(secret))
	authSecret := sha2.Sum(nil)
	sha2.Reset()

	encSecretStr := os.Getenv(ginSessionEncryptionSecretKey)
	if encSecretStr == "" {
		log.Panic("$" + ginSessionEncryptionSecretKey + " is not set")
	}

	sha2.Write([]byte(encSecretStr))
	encSecret := sha2.Sum(nil)
	sha2.Reset()

	sessionStore := cookie.NewStore(authSecret, encSecret)
	sessionStore.Options(sessions.Options{
		Path:     "/",
		SameSite: http.SameSiteDefaultMode,
		MaxAge:   3600 * 12, // 12hrs,
		Secure:   true,
		HttpOnly: true,
	})
	g.Use(sessions.Sessions(SessionCookieName, sessionStore))
}
