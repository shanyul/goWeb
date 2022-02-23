package api

import (
	"designer-api/pkg/app"
	"designer-api/pkg/e"
	"designer-api/pkg/logging"
	"designer-api/pkg/upload"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UploadApi struct{}

func (api *UploadApi) UploadImage(c *gin.Context) {
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		app.Response(c, http.StatusInternalServerError, e.ERROR_UPLOAD_FILE_NOT_INPUT, nil, "")
		return
	}

	if image == nil {
		app.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil, "")
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		app.Response(c, http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil, "")
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		logging.Warn(err)
		app.Response(c, http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil, "")
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		app.Response(c, http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil, "")
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, map[string]string{
		"url":     upload.GetImageFullUrl(imageName),
		"saveUrl": savePath + imageName,
	}, "")
}
