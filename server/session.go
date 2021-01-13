/**
 * @Author: li.zhang
 * @Description:
 * @File:  session
 * @Version: 1.0.0
 * @Date: 2020/12/10 下午3:53
 */
package main

import (
	"fmt"
	"io"
	"net"
)

type Session struct {
	conn        net.Conn
	readMsgChan chan []byte
	sendMsgChan chan []byte
}

var connRoom = make(map[net.Conn]bool)

func NewSession(conn net.Conn) *Session {

	fmt.Println("NewSession")
	s := &Session{
		conn:        conn,
		readMsgChan: make(chan []byte),
		sendMsgChan: make(chan []byte),
	}
	connRoom[conn] = true
	go s.run()
	return s
}

func (s *Session) run() {
	go s.readMsg()
	go s.sendMsg()
}

func (s *Session) sendMsg() {
	for {
		select {
		case msg := <-s.sendMsgChan:
			for conn := range connRoom {
				_, err := conn.Write(msg)
				if err != nil {
					fmt.Println("send msg err :", err)
				}
			}

		}
	}
}

func (s *Session) readMsg() {
	// 创建消息缓冲区
	buffer := make([]byte, 1024)
	for {

		// 读取客户端发来的消息放入缓冲区
		n, err := s.conn.Read(buffer)
		if err != nil {
			logger.Println(err)
			return
		}
		// 断开连接从map中删掉
		if err == io.EOF {
			delete(connRoom, s.conn)
		}
		s.sendMsgChan <- buffer[0:n]
		// 转化为字符串输出
		clientMsg := string(buffer[0:n])
		fmt.Printf("收到消息 %v : %v", s.conn.RemoteAddr(), clientMsg)

	}

}
