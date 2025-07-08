package middleware

import (
	"fmt"
	"log"
	"os"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/shared/datatype"
)

type CanGetStatusCode interface {
	StatusCode() int
}

func RecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				isProduction := os.Getenv("ENV") == "prod" || os.Getenv("GIN_MODE") == "release"
				
				if appErr, ok := r.(CanGetStatusCode); ok {
					c.JSON(appErr.StatusCode(), appErr)
					
					if !isProduction {
						log.Printf("Error: %+v\n", appErr)
						panic(r)
					}
					
					return
				}
				
				appErr := datatype.ErrInternalServerError
				
				if isProduction {
					c.JSON(appErr.StatusCode(), appErr.WithDebug(""))
				} else {
					c.JSON(appErr.StatusCode(), appErr.WithDebug(fmt.Sprintf("%s", r)))
					panic(r)
				}
			}
		}()
		
		c.Next()
	}
}
