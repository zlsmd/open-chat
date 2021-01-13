/**
 * @Author: li.zhang
 * @Description:
 * @File:  view
 * @Version: 1.0.0
 * @Date: 2020/12/10 下午2:09
 */
package ui

import (
	"fmt"
	aui "github.com/andlabs/ui"
	"github.com/zlsmd/zchat/client/chat"
)

func ChatWindow(ip string, port uint64) {
	connChat := chat.NewConnChat(ip, port)
	err := aui.Main(func() {
		// 生成：文本框
		name := aui.NewEntry()

		// 生成：按钮
		button := aui.NewButton(`发送`)
		// 设置：按钮点击事件
		button.OnClicked(func(*aui.Button) {
			connChat.SendMsg(name.Text())
		})
		// 生成：标签
		greeting := aui.NewLabel(``)
		go func() {
			for {

				select {
				case msg := <-connChat.ReadMsgChan:
					fmt.Println("read server msg :", string(msg))
					greeting.SetText(string(msg) + `！`)
				}
			}
		}()

		// 生成：垂直容器
		box := aui.NewVerticalBox()

		// 往 垂直容器 中添加 控件
		box.Append(aui.NewLabel(`请输入内容：`), false)
		box.Append(name, false)
		box.Append(button, false)
		box.Append(greeting, false)

		// 生成：窗口（标题，宽度，高度，是否有 菜单 控件）
		window := aui.NewWindow(`ChatRoom`, 500, 300, false)

		// 窗口容器绑定
		window.SetChild(box)

		// 设置：窗口关闭时
		window.OnClosing(func(*aui.Window) bool {
			// 窗体关闭
			aui.Quit()
			// 关闭会话
			connChat.Close()
			return true
		})

		// 窗体显示
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}

func LoginWindow() bool {
	var ok bool
	err := aui.Main(func() {
		// 生成：文本框
		name := aui.NewEntry()
		// 密码框
		pass := aui.NewPasswordEntry()

		// 生成：标签
		greeting := aui.NewLabel(``)

		// 生成：按钮
		button := aui.NewButton(`登录`)

		// 生成：垂直容器
		box := aui.NewVerticalBox()

		// 往 垂直容器 中添加 控件
		box.Append(aui.NewLabel(`请输入用户名：`), false)
		box.Append(name, false)
		box.Append(aui.NewLabel(`请输入密码：`), false)
		box.Append(pass, false)
		box.Append(button, true)
		box.Append(greeting, false)

		// 生成：窗口（标题，宽度，高度，是否有 菜单 控件）
		window := aui.NewWindow(`ChatRoom`, 500, 300, false)

		// 设置：按钮点击事件
		button.OnClicked(func(*aui.Button) {
			username := name.Text()
			if username == "" {
				greeting.SetText("请填写用户名!")
			}
			password := pass.Text()
			if password == "" {
				greeting.SetText("请填写密码!")
			}
			if username == "zhangli" && password == "123456" {
				ok = true
				aui.Quit()
				window.Destroy()
				return
			} else {
				greeting.SetText("用户名或密码错误!")
			}

		})
		// 窗口容器绑定
		window.SetChild(box)

		// 设置：窗口关闭时
		window.OnClosing(func(*aui.Window) bool {
			// 窗体关闭
			aui.Quit()
			return true
		})

		// 窗体显示
		window.Show()
	})
	if err != nil {
		panic(err)
	}
	return ok
}
