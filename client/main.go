/**
 * @Author: li.zhang
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2020/12/8 下午3:23
 */
package main

import (
	"flag"
	_ "github.com/andlabs/ui/winmanifest"
	"github.com/zlsmd/zchat/client/ui"
	"log"
	"os"
)

var (
	help   bool
	port   uint64
	ip     string
	logger = log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)
)

func init() {
	flag.BoolVar(&help, "h", false, "this help")
	flag.Uint64Var(&port, "p", 8080, "tcp server port")
	flag.StringVar(&ip, "a", "192.168.2.213", "tcp server ip")
}
func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	// 登录界面
	ok := ui.LoginWindow()

	if ok {
		// 聊天界面
		ui.ChatWindow(ip, port)
	}

}
