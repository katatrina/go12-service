package middleware

import (
	"strings"
	
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/katatrina/go12-service/shared/datatype"
)

func extractToken(val string) (string, error) {
	token := strings.TrimPrefix(val, "Bearer ")
	
	if token == "" {
		return "", datatype.ErrUnauthorized.WithDebug("Token is required")
	}
	
	return token, nil
}

type ITokenValidator interface {
	Validate(tokenStr string) (*jwt.RegisteredClaims, error)
}

func AuthMiddleware(tokenValidator ITokenValidator) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractToken(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}
		
		claims, err := tokenValidator.Validate(token)
		if err != nil {
			panic(datatype.ErrUnauthorized.WithWrap(err).WithDebug(err.Error()))
		}
		
		c.Set(datatype.KeyRequester, datatype.NewRequester(claims.Subject))
		c.Next()
	}
}
