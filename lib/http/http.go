package http

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type RequestInterface struct {
	Url     string          `json:"url"`
	Request HttpRespone     `json:"request"`
	Body    *strings.Reader `json:"body"`
}

// 设置跳过证书验证
func (s *RequestInterface) SetInsecureSkipVerify() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// 将该 Transport 赋值给 DefaultTransport
	http.DefaultTransport = tr
}
func (s *RequestInterface) SetRequestUrl(url string) {
	s.Url = url

}
func (s *RequestInterface) SetPostBody(body string) *strings.Reader {
	return strings.NewReader(body)
}
func (s *RequestInterface) HttpPostRequest(body *strings.Reader) HttpRespone {
	if s.Url == "" {
		s.Request.Error = errors.New("url为空，请检查请求是否设置url")
		return s.Request
	}
	return s.Https(Post, s.Url, body)

}
func (s *RequestInterface) HttpGetRequest() HttpRespone {
	if s.Url == "" {
		s.Request.Error = errors.New("url为空，请检查请求是否设置url")
		return s.Request
	}
	return s.Https(Get, s.Url, strings.NewReader(""))

}
func (s *RequestInterface) Https(method, url string, body *strings.Reader) HttpRespone {
	var Requset HttpRespone
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		Requset.Error = err

		return Requset
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if response != nil {
		Requset.StatusCode = response.StatusCode
	}
	if err != nil {
		Requset.Error = err
		return Requset
	}

	defer response.Body.Close()
	bodystr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Requset.Error = err
		return Requset
	}
	Requset.Body = string(bodystr)
	return Requset
}
