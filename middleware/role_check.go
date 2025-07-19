package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/shared/datatype"
	sharedmodel "github.com/katatrina/go12-service/shared/model"
)

func CheckRoles(roles ...datatype.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(datatype.KeyRequester).(datatype.Requester)
		
		for _, role := range roles {
			if role == requester.GetRole() {
				c.Next()
				return
			}
		}
		
		panic(datatype.ErrForbidden.WithError(sharedmodel.ErrUserRoleNotAllowed.Error()))
	}
}
