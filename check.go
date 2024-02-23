package main

import (
	"strings"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

// 返回值1为HTTPS,2为Socks4,3为Socks5,0为不可用
func CheckProxy(i int) int {
	var proxy string = proxies[i]
	var isProxy = strings.Contains(HttpsGet(proxy), "百度")
	if isProxy {
		return 1
	}
	isProxy = strings.Contains(Socks4Get(proxy), "百度")
	if isProxy {
		return 2
	}
	isProxy = strings.Contains(Socks5Get(proxy), "百度")
	if isProxy {
		return 3
	} else {
		return 0
	}
}

func Socks4Get(proxy string) string {
	client := &fasthttp.Client{}
	client.Dial = fasthttpproxy.FasthttpSocksDialer(proxy)
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("https://www.baidu.com")
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	if err != nil {
		//fmt.Println(err)
		return ""
	}
	return string(resp.Body())
}

func Socks5Get(proxy string) string {
	client := &fasthttp.Client{}
	client.Dial = fasthttpproxy.FasthttpSocksDialer("socks5://" + proxy)
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("https://www.baidu.com")
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	if err != nil {
		//fmt.Println(err)
		return ""
	}
	return string(resp.Body())
}

func HttpsGet(proxy string) string {
	client := &fasthttp.Client{}
	client.Dial = fasthttpproxy.FasthttpHTTPDialer(proxy)
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("https://www.baidu.com")
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	if err != nil {
		//fmt.Println(err)
		return ""
	}
	return string(resp.Body())
}
