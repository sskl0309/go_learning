package main

import (
	"fmt"
	"net"
)

type User struct {
	ID     int
	ch     chan string
	Conn   net.Conn
	Server *Server
}

func NewUser(id int, conn net.Conn, s *Server) *User {

	user := &User{
		ID:     id,
		ch:     make(chan string),
		Conn:   conn,
		Server: s,
	}
	go user.ListenMessage()
	return user
}

func (u *User) ListenMessage() {
	for msg := range u.ch {
		u.Conn.Write([]byte(msg + "\n"))
	}
}

func (u *User) Online() {
	u.Server.usersLock.Lock()
	u.Server.users[u.ID] = u
	u.Server.usersLock.Unlock()
	u.Server.Broadcast(u, "已上线")
}

func (u *User) Offline() {
	u.Server.usersLock.Lock()
	delete(u.Server.users, u.ID)
	u.Server.usersLock.Unlock()
	u.Server.Broadcast(u, "已下线")
}

func (u *User) sendMessage(msg string) {
	u.ch <- msg
}

func (u *User) Send(msg string) {
	if msg == "who" {
		u.Server.usersLock.Lock()
		for _, user := range u.Server.users {
			u.sendMessage(fmt.Sprintf("[用户%d %s]:[在线]", user.ID, user.Conn.RemoteAddr().String()))
		}
		u.Server.usersLock.Unlock()
	} else {
		u.Server.Broadcast(u, msg)
	}
}
