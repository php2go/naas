package middleware

import (
	"time"

	"github.com/nilorg/naas/internal/server/auth"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/spf13/viper"
)

// NewJwtMiddleware 创建jwt授权中间件
func NewJwtMiddleware() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "naas",
		Key:             []byte(viper.GetString("jwt.secret")),
		Timeout:         viper.GetDuration("jwt.timeout") * time.Minute,
		MaxRefresh:      viper.GetDuration("jwt.max_refresh") * time.Minute,
		IdentityKey:     jwt.IdentityKey,
		PayloadFunc:     auth.PayloadFunc,
		IdentityHandler: auth.IdentityHandler,
		Authenticator:   auth.Authenticator,
		Authorizator:    auth.Authorizator,
		Unauthorized:    auth.Unauthorized,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
}
