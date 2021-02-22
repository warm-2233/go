package main

import (
	"net"
	"net/http"
	// "path/filepath"
)
import "os"

import "fmt"

// func index(writer http.ResponseWriter, request *http.Request) {
// 	fmt.Fprintf(writer, "index")
// }package
var path, _ = os.Getwd()

func main() {
	getIp()

	// path = filepath.Join(path, "files")
	fmt.Println(path)
	fmt.Println("host:127.0.0.1")
	fmt.Println("port:8080")
	fmt.Println("url of http://127.0.0.1:8080")
	//创建一个Http多路复用器
	mux := http.NewServeMux()

	//定义一个Http文件服务器,本机的绝对路径（大家可以试试自己机器上）
	files := http.FileServer(http.Dir(path))

	//去掉URL路径前缀，返回指定文件
	mux.Handle("/", http.StripPrefix("/", files))

	//接收到URL为 ‘/’ 交给 index 函数处理
	// mux.HandleFunc("/", index)

	//指定端口，传递多路复用器
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()

}

func getIp() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}
