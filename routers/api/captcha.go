package api

import (
	"bytes"
	"designer-api/internal/service"
	"designer-api/pkg/app"
	"designer-api/pkg/e"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"path"
	"time"
)

type CaptchaApi struct {
	captchaService service.CaptchaService
}

func (api *CaptchaApi) Get(c *gin.Context) {
	captchaData := api.captchaService.Get()
	if captchaData.CaptchaId == "" {
		app.Response(c, http.StatusOK, e.ERROR_GET_CAPTCHA_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, captchaData)
}

func (api *CaptchaApi) Show(c *gin.Context) {
	w := com.StrTo(c.DefaultQuery("width", "0")).MustInt()
	h := com.StrTo(c.DefaultQuery("height", "0")).MustInt()
	d := com.StrTo(c.DefaultQuery("download", "0")).MustInt()
	if w == 0 {
		w = captcha.StdWidth
	}
	if h == 0 {
		h = captcha.StdHeight
	}
	download := false
	if d == 1 {
		download = true
	}
	image := c.Param("image")
	ext := path.Ext(image)
	captchaId := image[:len(image)-len(ext)]
	_ = api.server(c.Writer, c.Request, captchaId, ext, "zh", download, w, h)
}

func (api *CaptchaApi) server(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	if id == "" {
		return captcha.ErrNotFound
	}

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))

	return nil
}
