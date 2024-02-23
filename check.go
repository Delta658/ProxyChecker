package main

import (
	"log"
	"strings"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

// 返回值1为HTTPS,2为Socks4,3为Socks5,0为不可用
func CheckProxy(i int) int {
	var proxy string = proxies[i]
	var isProxy = strings.Contains(HttpsGet(proxy), "google")
	if isProxy {
		return 1
	}
	isProxy = strings.Contains(Socks4Get(proxy), "google")
	if isProxy {
		return 2
	}
	isProxy = strings.Contains(Socks5Get(proxy), "google")
	if isProxy {
		return 3
	} else {
		return 0
	}
}

func Socks4Get(proxy string) string {
	client := &fasthttp.Client{}
	client.Dial = fasthttpproxy.FasthttpSocksDialer("socks4://" + proxy)
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("https://www.google.com")
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()
	if err := client.Do(req, resp); err != nil {
		log.Fatal(err)
	}
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	return string(resp.Body())
}

func Socks5Get(proxy string) string {
	client := &fasthttp.Client{}
	client.Dial = fasthttpproxy.FasthttpSocksDialer("socks5://" + proxy)
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("https://www.google.com")
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()
	if err := client.Do(req, resp); err != nil {
		log.Fatal(err)
	}
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	return string(resp.Body())
}

func HttpsGet(proxy string) string {
	client := &fasthttp.Client{}
	client.Dial = fasthttpproxy.FasthttpSocksDialer("https://" + proxy)
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("https://www.google.com")
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()
	if err := client.Do(req, resp); err != nil {
		log.Fatal(err)
	}
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	return string(resp.Body())
}
