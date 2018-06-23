use 'godoc cmd/github.com/jackdoe/gin-basic-auth-dynamic' for documentation on the github.com/jackdoe/gin-basic-auth-dynamic command 

PACKAGE DOCUMENTATION

package auth
    import "github.com/jackdoe/gin-basic-auth-dynamic"


FUNCTIONS

func BasicAuth(callback func(*gin.Context, string, string, string) bool) gin.HandlerFunc
    Makes BasicAuth function for empty realm

func BasicAuthForRealm(callback func(*gin.Context, string, string, string) bool, realm string) gin.HandlerFunc
    Calls the callback function with (realm, user, password) if it returns
    true, then the user is authenticated, otherwise request is rejected with
    401 header Example usage:

	r := gin.Default()
	r.Use(auth.BasicAuth(func(context *gin.Context, realm, user, pass string) bool {
		return user == "something" && pass == "something else"
	}))

    You are also passed the context so you can do c.Set() stuff, for
    example:

	r := gin.Default()
	r.Use(auth.BasicAuth(func(context *gin.Context, realm, user, pass string) bool {
		user := db.FindUser(user, pass)
		if user != nil {
			c.Set("user", user)
			return true
		}
		return false
	}))


