package util

import (
	"designer-api/pkg/logging"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Get 请求
func Get(requestUrl string, timeout time.Duration) (string, error) {
	// 超时时间：5秒
	if timeout == 0 {
		timeout = 5
	}
	client := &http.Client{Timeout: timeout * time.Second}
	resp, err := client.Get(requestUrl)
	if err != nil {
		logging.Info(requestUrl, err)
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.Info(requestUrl, err)
		return "", err
	}
	return string(body), nil
}

// Post 请求
func Post(requestUrl string, data map[string]string, timeout time.Duration) (string, error) {
	// 超时时间：5秒
	if timeout == 0 {
		timeout = 5
	}
	client := &http.Client{Timeout: timeout * time.Second}
	var postData url.Values
	for k, v := range data {
		postData.Add(k, v)
	}
	resp, err := client.PostForm(requestUrl, postData)
	if err != nil {
		logging.Info(requestUrl, data, err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.Info(requestUrl, data, err)
		return "", err
	}
	return string(body), nil
}
