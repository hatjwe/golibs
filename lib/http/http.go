package http

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	urls "net/url"
	"strings"
	"time"
)

type RequestInterface struct {
	Url       string              `json:"url"`
	Request   HttpRespone         `json:"request"`
	Body      *strings.Reader     `json:"body"`
	Headers   map[string]string   `json:"headers"` // 存储请求头
	Timeout   time.Duration       // 全局超时时间
	Proxy     string              // 上游代理地址
	ReqHeader map[string][]string //请求的header头

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
func (s *RequestInterface) getAllHeader() map[string][]string {
	return s.ReqHeader

}

// 设置全局超时时间
func (s *RequestInterface) SetTimeout(timeout time.Duration) {
	s.Timeout = timeout
}

// 设置上游代理
func (s *RequestInterface) SetProxy(proxy string) {
	s.Proxy = proxy
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

// 设置请求头
func (s *RequestInterface) SetHeaders(headers map[string]string) {
	if s.Headers == nil {
		s.Headers = make(map[string]string)
	}
	for key, value := range headers {
		s.Headers[key] = value
	}
}
func (s *RequestInterface) Https(method, url string, body *strings.Reader) HttpRespone {
	var Requset HttpRespone
	client := &http.Client{}
	// 设置超时时间（如果有）
	if s.Timeout > 0 {
		client.Timeout = s.Timeout
	}

	// 设置代理
	if s.Proxy != "" {
		proxyURL, err := urls.Parse(s.Proxy)
		if err != nil {
			Requset.Error = err
			return Requset
		}
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
		client.Transport = transport
	}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		Requset.Error = err

		return Requset
	}
	// 添加自定义请求头
	for key, value := range s.Headers {
		request.Header.Set(key, value)
	}

	// 获取完整的请求信息

	response, err := client.Do(request)
	if response != nil {
		Requset.StatusCode = response.StatusCode
	}
	if err != nil {
		Requset.Error = err
		return Requset
	}

	defer response.Body.Close()
	// 获取完整的响应信息

	bodystr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Requset.Error = err
		return Requset
	}
	Requset.Body = string(bodystr)

	s.ReqHeader = make(map[string][]string) // 初始化 map
	for key, values := range request.Header {
		s.ReqHeader[key] = values
	}
	return Requset
}
