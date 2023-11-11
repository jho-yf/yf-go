package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户列表
	OnlineUserMap map[string]*User
	mapLock       sync.RWMutex

	// 消息广播的channel
	BroadcastChannel chan string
}

func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:               ip,
		Port:             port,
		OnlineUserMap:    make(map[string]*User),
		BroadcastChannel: make(chan string),
	}
}

func (server *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	// close listen socket
	defer listener.Close()

	// 监听广播消息
	go server.ListenBroadcastMessage()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}

		// do on connected
		go server.OnConnected(conn)
	}
}

func (server *Server) OnConnected(conn net.Conn) {
	user := NewUser(conn, server.OnUserOnline, server.OnUserOffline)

	// 用户上线
	fmt.Printf("User '%s' connected.", user.Name)
	user.Online()

	go server.HandleMsg(user)

	for {
		// 阻塞
		select {
		case <-user.IsLive:
			// 当前用户为活跃状态，重置定时器，激活select更新以下定时器
		case <-time.After(time.Second * 120):
			server.SendMsg(user, "You are offline by the system.")
			user.conn.Close()
			// 退出协程
			return
		}
	}

}

// 接收user对应客户端的消息
func (server *Server) HandleMsg(user *User) {
	buf := make([]byte, 4096)
	for {
		n, err := user.conn.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("Connection[", user.Addr, "] Read Err:", err)
			user.Offline()
			return
		}

		if n == 0 {
			user.Offline()
			return
		}

		// 转换用户消息，去掉\n
		msg := string(buf[:n-1])

		if msg == "online list" {
			server.mapLock.Lock()
			server.SendMsg(user, "[Online List]\n")
			for _, u := range server.OnlineUserMap {
				onlineMsg := "[" + u.Addr + "]" + u.Name + "(online)\n"
				server.SendMsg(user, onlineMsg)
			}
			server.mapLock.Unlock()
		} else if len(msg) > 7 && msg[:7] == "rename|" {
			newName := strings.Split(msg, "|")[1]
			_, ok := server.OnlineUserMap[newName]
			if ok {
				server.SendMsg(user, "Name '"+newName+"' already exist!")
			} else {
				server.mapLock.Lock()
				delete(server.OnlineUserMap, user.Name)
				user.Name = newName
				server.OnlineUserMap[user.Name] = user
				server.mapLock.Unlock()
				server.SendMsg(user, "Rename '"+newName+"' success!")
			}
		} else {
			server.Broadcast(user, msg)
		}

		user.IsLive <- true
	}
}

func (server *Server) OnUserOnline(user *User, msg string) {
	// 将用户添加至onlineMap中
	server.mapLock.Lock()
	server.OnlineUserMap[user.Name] = user
	server.mapLock.Unlock()

	// 广播当前用户上线
	server.Broadcast(user, msg)
}

func (server *Server) OnUserOffline(user *User, msg string) {
	// 清除缓存
	server.mapLock.Lock()
	delete(server.OnlineUserMap, user.Name)
	server.mapLock.Unlock()

	// 广播下线消息
	server.Broadcast(user, msg)
}

func (server *Server) Broadcast(user *User, msg string) {
	server.BroadcastChannel <- "[" + user.Addr + "]" + user.Name + ": " + msg
}

func (server *Server) SendMsg(user *User, msg string) {
	user.conn.Write([]byte(msg))
}

func (server *Server) ListenBroadcastMessage() {
	for {
		msg := <-server.BroadcastChannel

		// 广播消息
		server.mapLock.Lock()
		for _, user := range server.OnlineUserMap {
			user.Channel <- msg
		}
		server.mapLock.Unlock()
	}
}
