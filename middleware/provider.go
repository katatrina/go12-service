package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/shared/datatype"
)

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

func (p *MiddlewareProvider) CheckRoles(roles ...datatype.UserRole) gin.HandlerFunc {
	return CheckRoles(roles...)
}
