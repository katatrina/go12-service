package userhttpcontroller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/shared/datatype"
)

func (ctrl *UserHTTPController) GetProfile(c *gin.Context) {
	requester := c.MustGet(datatype.KeyRequester).(datatype.Requester)
	
	c.JSON(http.StatusOK, datatype.ResponseSuccess(requester))
}
