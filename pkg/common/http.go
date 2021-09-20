package common

import (
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
)

func DoJsonHttp(url string, body []byte, action string) ([]byte, error) {
	payload := bytes.NewReader(body)
	req, err := http.NewRequest(action, url, payload)

	if err != nil {
		Logger().Error("HTTP 请求构造失败!：", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		Logger().Error("Http 请求失败！", err)
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		Logger().Error("读取http返回结果失败！", err)
	}
	Logger().Info("http 请求返回值:", string(data))
	return data, nil
}
