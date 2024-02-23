package main

import (
	"os"
	"sync"
)

var _Socks4Files, _Socks5Files, _HttpsFiles *os.File
var Socks4Mu, Socks5Mu, HttpsMu sync.Mutex //创建三个互斥锁，保证写入是线程安全的

func FileInit() bool {
	Socks4Files, err1 := os.OpenFile("Socks4.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Socks5Files, err2 := os.OpenFile("Socks5.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	HttpsFiles, err3 := os.OpenFile("HTTPS.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//我这里是想让这些当全局变量的，但是我不知道咋创建，就用这看起来就很蠢的办法算了
	_Socks4Files = Socks4Files
	_Socks5Files = Socks5Files
	_HttpsFiles = HttpsFiles
	if err1 != nil || err2 != nil || err3 != nil {
		return false
	} else {
		return true
	}
}

func WriteSocks4(proxy string) bool {
	Socks4Mu.Lock()
	_, err := _Socks4Files.WriteString(proxy + "\n")
	Socks4Mu.Unlock()
	if err != nil {
		return false
	} else {
		return true
	}
}

func WriteSocks5(proxy string) bool {
	Socks5Mu.Lock()
	_, err := _Socks5Files.WriteString(proxy + "\n")
	Socks5Mu.Unlock()
	if err != nil {
		return false
	} else {
		return true
	}
}

func WriteHttps(proxy string) bool {
	HttpsMu.Lock()
	_, err := _HttpsFiles.WriteString(proxy + "\n")
	HttpsMu.Unlock()
	if err != nil {
		return false
	} else {
		return true
	}
}
