package common

import (
	"bytes"
	"crypto/tls"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func DoJsonHttp(url string, body []byte, action string) ([]byte, error) {
	payload := bytes.NewReader(body)
	req, err := http.NewRequest(action, url, payload)

	if err != nil {
		logrus.Error("HTTP 请求构造失败!：", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Error("Http 请求失败！", err)
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("读取http返回结果失败！", err)
	}
	logrus.Info("http 请求返回值:", string(data))
	return data, nil
}
