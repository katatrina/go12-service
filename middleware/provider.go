package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/shared/datatype"
)

type MiddlewareProvider struct {
	tokenIntrospector ITokenIntrospector
}

func NewMiddlewareProvider(tokenValidator ITokenIntrospector) *MiddlewareProvider {
	return &MiddlewareProvider{
		tokenIntrospector: tokenValidator,
	}
}

func (p *MiddlewareProvider) Auth() gin.HandlerFunc {
	return Auth(p.tokenIntrospector)
}

func (p *MiddlewareProvider) CheckRoles(roles ...datatype.UserRole) gin.HandlerFunc {
	return CheckRoles(roles...)
}
