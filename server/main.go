/**
 * @Author: li.zhang
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2020/12/8 上午11:39
 */
package main

import (
	"flag"
	"fmt"
	"github.com/zlsmd/zchat/server/handler"
	"github.com/zlsmd/zchat/server/model"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

var (
	help        bool
	port        uint64
	localIpAddr string
	logger      = log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)
)

func init() {
	flag.BoolVar(&help, "h", false, "this help")
	flag.Uint64Var(&port, "p", 8080, "tcp server port")

	localAddrs, err := net.InterfaceAddrs()
	if err != nil {
		logger.Panicf("get service ip address err:%v", err)
	}
	for _, address := range localAddrs {
		// 检查ip地址判断是否回环地址
		if inet, ok := address.(*net.IPNet); ok && !inet.IP.IsLoopback() {
			if inet.IP.To4() != nil {
				localIpAddr = inet.IP.To4().String()
				break
			}
		}
	}
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	model.InitOrm()

	go listenTcp()

	//1.注册一个处理器函数
	http.HandleFunc("/login", handler.Login)

	//2.设置监听的TCP地址并启动服务
	//参数1:TCP地址(IP+Port)
	//参数2:handler handler参数一般会设为nil，此时会使用DefaultServeMux。
	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if err != nil {
		fmt.Printf("http.ListenAndServe()函数执行错误,错误为:%v\n", err)
		return
	}
}

func listenTcp() {
	address := localIpAddr + ":" + strconv.Itoa(int(port))
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	logger.Println("tcp server start listen on :", address)

	for {
		//循环接入所有客户端得到专线连接
		conn, err := listener.Accept()
		if err != nil {
			logger.Println(err)
			continue
		}
		// 创建会话
		NewSession(conn)
	}
}
