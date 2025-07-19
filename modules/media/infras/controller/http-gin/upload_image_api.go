package mediahttpgin

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"
	
	media "github.com/katatrina/go12-service/modules/media/model"
	mediaservice "github.com/katatrina/go12-service/modules/media/service"
	"github.com/katatrina/go12-service/shared/datatype"
	
	"github.com/gin-gonic/gin"
)

func (ctrl *MediaHTTPController) UploadImageAPI(c *gin.Context) {
	folder := c.DefaultPostForm("folder", "uploads")
	fileHeader, err := c.FormFile("file")
	
	if err != nil {
		panic(datatype.ErrBadRequest.WithError(err.Error()))
	}
	
	filename := fmt.Sprintf("%d_%s", time.Now().UTC().UnixNano(), fileHeader.Filename)
	dst := fmt.Sprintf("%s/%s", folder, filename)
	
	// if err := c.SaveUploadedFile(fileHeader, dst); err != nil {
	// 	panic(datatype.ErrBadRequest.WithError(err.Error()))
	// }
	
	file, err := fileHeader.Open()
	
	if err != nil {
		panic(datatype.ErrBadRequest.WithError(err.Error()))
	}
	
	defer file.Close()
	
	contentType := fileHeader.Header.Get("Content-Type")
	
	ctrl.uploader.SaveFileUpload(c.Request.Context(), file, dst, contentType, fileHeader.Size)
	
	mediaCreate := media.MediaCreateDTO{
		Filename:  dst,
		CloudName: "aws-s3",
		Size:      fileHeader.Size,
		Ext:       strings.ReplaceAll(filepath.Ext(fileHeader.Filename), ".", ""),
	}
	
	// Create command and call handler
	cmd := mediaservice.CreateCommand{MediaCreate: mediaCreate}
	
	id, err := ctrl.createHandler.Execute(c.Request.Context(), &cmd)
	
	if err != nil {
		panic(err)
	}
	
	// url := fmt.Sprintf("http://localhost:3000%s", strings.Replace(dst, ".", "", 1))
	
	mediaCreate.Fulfill(ctrl.uploader.GetDomain())
	
	c.JSON(http.StatusCreated, gin.H{"data": id, "media": mediaCreate})
}
