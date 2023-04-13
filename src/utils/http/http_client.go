package http

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	POST = "post"
	GET  = "get"
)

var HttpClient *httpClient

func init() {
	HttpClient = &httpClient{}
}

type httpClient struct {
}

// Get 请求表单数据
func (c *httpClient) Get(apiUrl string, params map[string]string) (string, error) {
	data := url.Values{}
	for k, v := range params {
		data.Set(k, v)
	}

	var resp *http.Response
	var err error
	// 接口调用
	resp, err = http.Get(apiUrl + "?" + data.Encode())
	if err != nil {
		log.Println("err:", err)
		return "", err
	}
	// 接口回参处理
	return respHandler(resp)
}

// Post 请求表单数据
func (c *httpClient) Post(apiUrl string, params map[string]string) (string, error) {
	data := url.Values{}
	for k, v := range params {
		data.Set(k, v)
	}
	var req *http.Request
	var resp *http.Response
	var err error

	body := strings.NewReader(data.Encode())
	req, err = http.NewRequest(POST, apiUrl, body)
	if err != nil {
		log.Println("err:", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	log.Println("post 请求参数：", body)
	// 接口调用
	cli := http.Client{}
	resp, err = cli.Do(req)
	if err != nil {
		log.Println("err:", err)
		return "", err
	}
	// 接口回参处理
	return respHandler(resp)
}

// PostJson 请求json化参数
func (c *httpClient) PostJson(apiUrl string, jsonParams string) (string, error) {
	var jsonStr = []byte(jsonParams)
	body := bytes.NewBuffer(jsonStr)
	req, err := http.NewRequest(POST, apiUrl, body)
	if err != nil {
		log.Println("err:", err)
		return "", nil
	}
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	// req.Header.Set("Connection", "false")

	//req.Close = true
	log.Println("PostJson 请求参数：", body)
	cli := http.Client{Timeout: 300 * time.Second}
	// 接口调用
	resp, err := cli.Do(req)
	if err != nil {
		log.Println("do,err:", err)
		return "", nil
	}

	// 回参处理
	return respHandler(resp)
}

// 回参处理
func respHandler(resp *http.Response) (string, error) {
	rs, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("readAll,err:", err)
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("http请求，流关闭失败")
		}
	}(resp.Body)
	return string(rs), nil
}
