package auth

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type AuthResult struct {
	Success bool
	Text    string
}

// Calls the callback function with (realm, user, password) if it returns true, then the user is authenticated, otherwise request is rejected with 401 header
// Example usage:
// 	r := gin.Default()
//	r.Use(auth.BasicAuth(func(context *gin.Context, realm, user, pass string) bool {
//		ok := user == "something" && pass == "something else"
//		return AuthResult{Success: ok, Text: "not authorized"}
//	}))
//
// You are also passed the context so you can do c.Set() stuff, for example:
//	r := gin.Default()
//	r.Use(auth.BasicAuth(func(context *gin.Context, realm, user, pass string) AuthResult {
//		user := db.FindUser(user, pass)
//		if user != nil {
//			c.Set("user", user)
//			return auth.AuthResult{Success:true}
//		}
//		return auth.AuthResult{Success:false}
//	}))
//
func BasicAuthForRealm(callback func(*gin.Context, string, string, string) AuthResult, realm string) gin.HandlerFunc {
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

		if len(pair) != 2 {
			c.Header("WWW-Authenticate", realm)
			c.AbortWithStatus(401)
			return
		}
		res := callback(c, realm, pair[0], pair[1])
		if !res.Success {
			c.Header("WWW-Authenticate", realm)
			if res.Text != "" {
				c.String(401, res.Text)
			} else {
				c.Status(401)
				c.Writer.WriteHeaderNow()
			}

			c.Abort()
			return
		}
	}
}

// Makes BasicAuth function for empty realm
func BasicAuth(callback func(*gin.Context, string, string, string) AuthResult) gin.HandlerFunc {
	return BasicAuthForRealm(callback, "")
}
