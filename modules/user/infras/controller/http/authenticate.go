package userhttpcontroller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	userservice "github.com/katatrina/go12-service/modules/user/service"
	"github.com/katatrina/go12-service/shared/datatype"
)

func (ctrl *UserHTTPController) Authenticate(c *gin.Context) {
	var cmd userservice.AuthenticateCommand
	if err := c.ShouldBind(&cmd); err != nil {
		panic(datatype.ErrBadRequest.WithError(err.Error()))
	}
	
	resp, err := ctrl.authenticateCmdHandler.Execute(c.Request.Context(), &cmd)
	if err != nil {
		panic(err)
	}
	
	c.JSON(http.StatusOK, datatype.ResponseSuccess(resp))
}
