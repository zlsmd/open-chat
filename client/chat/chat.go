/**
 * @Author: li.zhang
 * @Description:
 * @File:  chat
 * @Version: 1.0.0
 * @Date: 2020/12/10 下午2:19
 */
package chat

import (
	"fmt"
	"net"
	"strconv"
)

type connChat struct {
	conn        net.Conn
	ReadMsgChan chan []byte
	sendMsgChan chan []byte
}

func NewConnChat(ip string, port uint64) *connChat {
	address := ip + ":" + strconv.Itoa(int(port))
	// 拨号远程地址，建立tcp连接
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	c := &connChat{
		conn:        conn,
		ReadMsgChan: make(chan []byte),
		sendMsgChan: make(chan []byte),
	}
	go c.run()
	return c
}

func (c *connChat) run() {
	go c.readMsg()
	for {
		select {
		case msg := <-c.sendMsgChan:
			_, err := c.conn.Write(msg)
			if err != nil {
				fmt.Println("send msg err :", err)
				c.conn.Close()
			}
		}

	}
}

func (c *connChat) SendMsg(msg string) {
	c.sendMsgChan <- []byte(msg)
}

func (c *connChat) readMsg() {
	for {
		// 预先准备消息缓冲区
		buffer := make([]byte, 1024)
		n, err := c.conn.Read(buffer)
		if err != nil {
			fmt.Println("read msg err =", err)
			return
		}
		c.ReadMsgChan <- buffer[0:n]
		serverMsg := string(buffer[0:n])
		fmt.Printf("服务端msg :%v", serverMsg)
		if serverMsg == "bye" {
			break
		}
	}
}

func (c *connChat) Close() {
	err := c.conn.Close()
	if err != nil {
		fmt.Println("close conn err :", err)
	}
}
