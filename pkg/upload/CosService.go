package upload

import (
	"context"
	"designer-api/pkg/setting"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
)

const CosUrl = "https://s1-1309644651.cos.ap-shanghai.myqcloud.com"

func setClient() *cos.Client {
	u, _ := url.Parse(CosUrl)
	// 用于Get Service 查询，默认全地域 service.cos.myqcloud.com
	b := &cos.BaseURL{BucketURL: u}
	// 1.永久密钥
	return cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  setting.WechatSetting.CosKey,
			SecretKey: setting.WechatSetting.CosSecret,
		},
	})
}

func UploadFile(path string, file *multipart.FileHeader) (url string, err error) {
	// 获取客户端实例
	client := setClient()
	fd, err := file.Open()
	if err != nil {
		return
	}
	defer fd.Close()
	_, err = client.Object.Put(context.Background(), path, fd, nil)
	if err != nil {
		return
	}

	return CosUrl + "/" + path, nil
}
