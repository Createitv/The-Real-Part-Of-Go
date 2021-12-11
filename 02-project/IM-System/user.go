package main

import "net"

type User struct {
	Name    string
	Address string
	C       chan string
	conn    net.Conn
}

// 创建一个用户的构造函数
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	// 启动一个goroutine

	user := &User{
		Name:    userAddr,
		Address: userAddr,
		C:       make(chan string),
		conn:    conn,
	}
	go user.ListenMessage()
	return user
}

// 箭头当前User channel的方法
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
