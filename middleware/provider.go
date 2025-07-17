package middleware

import "github.com/gin-gonic/gin"

type MiddlewareProvider struct {
	tokenValidator ITokenIntrospector
}

func NewMiddlewareProvider(tokenValidator ITokenIntrospector) *MiddlewareProvider {
	return &MiddlewareProvider{
		tokenValidator: tokenValidator,
	}
}

func (p *MiddlewareProvider) Auth() gin.HandlerFunc {
	return Auth(p.tokenValidator)
}
