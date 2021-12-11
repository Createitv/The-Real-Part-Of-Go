package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int
	//	在线用户map
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//	消息广播的channel
	Message chan string
}

// 创建server接口的构造函数
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

//
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message

		//	讲msg通知所有在线User
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 广播消息
func (this *Server) Broadcast(user *User, msg string) {
	sendMsg := "[" + user.Address + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

func (this *Server) Handle(conn net.Conn) {
	//	 当前业务连接
	user := NewUser(conn)

	fmt.Println("Listen connect success")
	// 用户上线,加入到map[string]string
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()

	//	广播当前用户上下信息
	this.Broadcast(user, "is online")

	//	当前handle阻塞
	select {}

}

// 启动服务器
func (this *Server) Start() {
	//	socket listen on
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.listen error:", err)
		return
	}
	fmt.Println("服务器等待客户端建立连接...")
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("net.close error:", err)
		}
	}(listener)
	// 启动Message的goroutine
	go this.ListenMessage()

	for {

		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listen.Accept error:", err)
			continue
		}

		// do handshake
		go this.Handle(conn)

		// close listen
	}
}
