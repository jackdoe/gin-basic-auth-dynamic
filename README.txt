use 'godoc cmd/github.com/jackdoe/gin-basic-auth-dynamic' for documentation on the github.com/jackdoe/gin-basic-auth-dynamic command 

PACKAGE DOCUMENTATION

package auth
    import "github.com/jackdoe/gin-basic-auth-dynamic"


FUNCTIONS

func BasicAuth(callback func(*gin.Context, string, string, string) AuthResult) gin.HandlerFunc
    Makes BasicAuth function for empty realm

func BasicAuthForRealm(callback func(*gin.Context, string, string, string) AuthResult, realm string) gin.HandlerFunc
    Calls the callback function with (realm, user, password) if it returns
    true, then the user is authenticated, otherwise request is rejected with
    401 header Example usage:

	r := gin.Default()
	r.Use(auth.BasicAuth(func(context *gin.Context, realm, user, pass string) bool {
		ok := user == "something" && pass == "something else"
		return AuthResult{Success: ok, Text: "not authorized"}
	}))

    You are also passed the context so you can do c.Set() stuff, for
    example:

	r := gin.Default()
	r.Use(auth.BasicAuth(func(context *gin.Context, realm, user, pass string) AuthResult {
		user := db.FindUser(user, pass)
		if user != nil {
			c.Set("user", user)
			return auth.AuthResult{Success:true}
		}
		return auth.AuthResult{Success:false}
	}))

TYPES

type AuthResult struct {
    Success bool
    Text    string
}


