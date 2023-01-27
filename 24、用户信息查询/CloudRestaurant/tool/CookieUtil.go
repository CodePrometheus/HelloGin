package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const CookieName = "cookie_user"
const CookieTimeLength = 10 * 60 //10分钟

func CookieAuth(context *gin.Context) (*http.Cookie, error) {
	cookie, err := context.Request.Cookie(CookieName)
	if err == nil {
		context.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		return cookie, nil
	} else {
		return nil, err
	}
}
