package userhttpcontroller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	userservice "github.com/katatrina/go12-service/modules/user/service"
	"github.com/katatrina/go12-service/shared/datatype"
)

func (h *UserHTTPController) Register(c *gin.Context) {
	var cmd userservice.RegisterCommand
	if err := c.ShouldBind(&cmd); err != nil {
		panic(datatype.ErrBadRequest.WithWrap(err).WithDebug(err.Error()))
	}
	
	user, err := h.registerCmdHandler.Execute(c.Request.Context(), &cmd)
	if err != nil {
		panic(err)
	}
	
	c.JSON(http.StatusCreated, datatype.ResponseSuccess(user))
}
