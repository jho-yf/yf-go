package main

import (
	"net"
)

type User struct {
	conn          net.Conn
	Name          string
	Addr          string
	Channel       chan string
	IsLive        chan bool
	OnlineAction  func(*User, string)
	OfflineAction func(*User, string)
}

func NewUser(conn net.Conn, onlineAction func(*User, string), offlineAction func(*User, string)) *User {
	userAddr := conn.RemoteAddr().String()

	return &User{
		conn:          conn,
		Name:          userAddr,
		Addr:          userAddr,
		Channel:       make(chan string),
		IsLive:        make(chan bool),
		OnlineAction:  onlineAction,
		OfflineAction: offlineAction,
	}
}

// 用户上线
func (user *User) Online() {
	user.OnlineAction(user, "i am online.")
	go user.ListenMessage()
}

// 用户下线
func (user *User) Offline() {
	user.OfflineAction(user, "i am offline.")
	close(user.IsLive)
	close(user.Channel)
	user.conn.Close()
}

// 监听User channel，一旦有消息就发送给客户端
func (user *User) ListenMessage() {
	for {
		msg := <-user.Channel
		user.conn.Write([]byte(msg + "\n"))
	}
}
