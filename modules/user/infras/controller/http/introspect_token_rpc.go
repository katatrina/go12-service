package userhttpcontroller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	userservice "github.com/katatrina/go12-service/modules/user/service"
	"github.com/katatrina/go12-service/shared/datatype"
)

func (ctrl *UserHTTPController) IntrospectTokenRpc(c *gin.Context) {
	var bodyData struct {
		Token string `json:"token"`
	}
	
	if err := c.ShouldBindJSON(&bodyData); err != nil {
		panic(datatype.ErrBadRequest.WithWrap(err).WithDebug(err.Error()))
	}
	
	user, err := ctrl.introspectCmdHandler.Execute(c.Request.Context(), &userservice.IntrospectCommand{
		Token: bodyData.Token,
	})
	if err != nil {
		panic(err)
	}
	
	c.JSON(http.StatusOK, datatype.ResponseSuccess(user))
}
