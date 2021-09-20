package common

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

func DoJsonHttp(url string, body []byte, action string) ([]byte, error) {
	payload := bytes.NewReader(body)
	req, err := http.NewRequest(action, url, payload)
	defer req.Body.Close()

	if err != nil {
		Logger().Error("HTTP 请求构造失败!：", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	client := *http.DefaultClient
	client.Timeout = 10 * time.Minute
	resp, err := client.Do(req)
	if err != nil {
		Logger().Error("Http 请求失败！", err)
		return nil, err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		Logger().Error("读取http返回结果失败！", err)
	}
	Logger().Info("http 请求返回值:", string(data))
	return data, nil
}
