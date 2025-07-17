package middleware

import (
	"strings"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/shared/datatype"
)

func extractToken(val string) (string, error) {
	token := strings.TrimPrefix(val, "Bearer ")
	
	if token == "" {
		return "", datatype.ErrUnauthorized.WithDebug("Token is required")
	}
	
	return token, nil
}

type ITokenIntrospector interface {
	Introspect(tokenStr string) (datatype.Requester, error)
}

func Auth(tokenValidator ITokenIntrospector) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractToken(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}
		
		requester, err := tokenValidator.Introspect(token)
		if err != nil {
			panic(datatype.ErrUnauthorized.WithWrap(err).WithDebug(err.Error()))
		}
		
		c.Set(datatype.KeyRequester, requester)
		c.Next()
	}
}
