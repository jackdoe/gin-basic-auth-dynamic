package auth

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// Calls the callback function with (realm, user, password) if it returns true, then the user is authenticated, otherwise request is rejected with 401 header
// Example usage:
// 	r := gin.Default()
//	r.Use(auth.BasicAuth(func(context *gin.Context, realm, user, pass string) bool {
//		return user == "something" && pass == "something else"
//	}))
//
// You are also passed the context so you can do c.Set() stuff, for example:
//	r := gin.Default()
//	r.Use(auth.BasicAuth(func(context *gin.Context, realm, user, pass string) bool {
//		user := db.FindUser(user, pass)
//		if user != nil {
//			c.Set("user", user)
//			return true
//		}
//		return false
//	}))
//
func BasicAuthForRealm(callback func(*gin.Context, string, string, string) bool, realm string) gin.HandlerFunc {
	if realm == "" {
		realm = "Authorization Required"
	}
	realm = "Basic realm=" + strconv.Quote(realm)

	return func(c *gin.Context) {
		// Search user in the slice of allowed credentials

		auth := strings.SplitN(c.GetHeader("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			c.Header("WWW-Authenticate", realm)
			c.AbortWithStatus(401)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !callback(c, realm, pair[0], pair[1]) {
			c.Header("WWW-Authenticate", realm)
			c.AbortWithStatus(401)
			return
		}
	}
}

// Makes BasicAuth function for empty realm
func BasicAuth(callback func(*gin.Context, string, string, string) bool) gin.HandlerFunc {
	return BasicAuthForRealm(callback, "")
}
