package utils

import (
	"go.uber.org/zap"
	"io"
	"nebula.xyz/global"
	"net/http"
)

func ZLMHttpRequest(path string, body io.Reader) (string, error) {
	url := "http://" + global.MediaServer.GetAddress() + ":" + global.MediaServer.GetRestful() + "/index/api/" + path
	global.Logger.Info(url)
	var resp *http.Response
	var err error
	if body == nil {
		url += "?secret=" + global.MediaServer.GetSecret()
		resp, err = http.Get(url)
	} else {
		resp, err = http.Post(url, "application/json", body)
	}
	if err != nil {
		global.Logger.Error("向Zlm发送请求失败", zap.String("path", path))
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		global.Logger.Error("读取Zlm返回失败", zap.Error(err))
		return "", err
	}
	return string(b), nil
}