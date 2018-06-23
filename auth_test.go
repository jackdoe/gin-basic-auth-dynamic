package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"

	"testing"
)

const ok = "ok"

func TestAuth(t *testing.T) {
	r := gin.Default()
	r.Use(BasicAuth(func(c *gin.Context, realm, user, pass string) bool {
		return user == "something" && pass == "something else"
	}))

	r.GET("/", func(c *gin.Context) {
		c.String(200, ok)
	})

	res1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(res1, req1)
	if res1.Body.String() == ok {
		t.Fatal("expected 401")
	}

	res2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/", nil)
	req2.SetBasicAuth("something", "something else")
	req2.Header.Set("Cookie", res1.Header().Get("Set-Cookie"))
	r.ServeHTTP(res2, req2)
	if res2.Body.String() != ok {
		t.Fatal("expected correct password")
	}

	res3 := httptest.NewRecorder()
	req3, _ := http.NewRequest("GET", "/", nil)
	req3.SetBasicAuth("something", "wrong password")
	r.ServeHTTP(res3, req3)
	if res3.Body.String() == ok {
		t.Fatal("expected auth")
	}
}
