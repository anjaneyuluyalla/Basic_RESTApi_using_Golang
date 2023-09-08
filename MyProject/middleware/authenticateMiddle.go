package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

const(
	user="name"
	passw="password"
)

func Authenticate(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authParts := strings.SplitN(authHeader, " ", 2)
	if len(authParts) != 2 || authParts[0] != "Basic" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	decodedCreds, err := base64.StdEncoding.DecodeString(authParts[1])
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	creds := string(decodedCreds)
	credentials := strings.SplitN(creds, ":", 2)
	if len(credentials) != 2 || credentials[0] != user || credentials[1] != passw {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}
