package api

import (
	"designer-api/pkg/app"
	"designer-api/pkg/e"
	"designer-api/pkg/logging"
	"designer-api/pkg/upload"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

func (api *UploadApi) Upload(c *gin.Context) {
	file, err := c.FormFile("filename")
	if err != nil {
		logging.Warn(err)
		app.Response(c, http.StatusInternalServerError, e.ERROR_UPLOAD_FILE_NOT_INPUT, nil, "")
		return
	}
	// 限制 20M
	if file.Size > 20*1024*1024 {
		logging.Warn(err)
		app.Response(c, http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_SIZE_FAIL, nil, "")
		return
	}

	filename := upload.GetImageName(file.Filename)
	path := fmt.Sprintf("aseert/%d%d/%s", time.Now().Year(), time.Now().Month(), filename)

	url, err := upload.UploadFile(path, file)
	if err != nil {
		logging.Warn(err)
		app.Response(c, http.StatusInternalServerError, e.ERROR_UPLOAD_FAIL, nil, "")
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, map[string]string{
		"name": file.Filename,
		"url":  url,
	}, "")
}
