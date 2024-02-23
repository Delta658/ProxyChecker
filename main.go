package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

var proxies []string //定义一个全局存放每一条proxy的变量

var HttpsCount, Socks4Count, Socks5Count int //用于计数

var console sync.Mutex //控制台输出互斥锁

const (
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Reset   = "\033[0m"
)

// 这个工程是练手并且当一个多线程访问网页的例子，会尽量使其改几个东西就可以换一个项目（
func main() {
	fmt.Println("代理IP检查器\n")
	fmt.Println("用途:多线程检查代理ip是否可用和代理ip的类型\n")
	fmt.Println("检查器自动从程序根目录的proxies.txt文件中获取代理ip进行验证。\n")
	fmt.Println("请输入您需要的线程数：")
	var ThreadCount int
	_, err := fmt.Scan(&ThreadCount)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Fatalln("创建线程池")
	CreateFool(ThreadCount)
	fmt.Println("正在读取代理IP...\n")
	ReadProxies()
	var count = len(proxies)
	fmt.Printf("成功读取%d条代理IP\n", &count)

	StartTask(count)

	fmt.Println("全部测试完毕！")
	fmt.Scanln()
}

func Work(i interface{}) {
	n := i.(int)
	var res int = CheckProxy(n)
	if res == 0 { //无法访问
		console.Lock()
		fmt.Println(Red + "[Failed]" + proxies[n] + Reset)
		console.Unlock()
	} else if res == 1 {
		console.Lock()
		fmt.Println(Green + "[HTTPS]" + proxies[n] + Reset)
		console.Unlock()
		WriteHttps(proxies[n])
	} else if res == 2 {
		console.Lock()
		fmt.Println(Green + "[Socks4]" + proxies[n] + Reset)
		console.Unlock()
		WriteSocks4(proxies[n])
	} else if res == 3 {
		console.Lock()
		fmt.Println(Green + "[Socks5]" + proxies[n] + Reset)
		console.Unlock()
		WriteSocks5(proxies[n])
	}

}

func ReadProxies() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// 循环扫描每一行
	for scanner.Scan() {
		// 获取每一行的文本
		line := scanner.Text()
		// 将每一行的文本追加到字符串切片中
		proxies = append(proxies, line)
	}

	// 检查是否有扫描错误
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
}
