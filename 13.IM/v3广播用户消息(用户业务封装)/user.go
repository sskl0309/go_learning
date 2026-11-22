package main

import (
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
	u.Server.users = append(u.Server.users, u)
	u.Server.usersLock.Unlock()
	u.Server.Broadcast(u, "已上线")
}

func (u *User) Offline() {
	u.Server.Broadcast(u, "已下线")
	u.Server.usersLock.Lock()
	for i := range u.Server.users {
		if u.Server.users[i] == u {
			u.Server.users = append(u.Server.users[:i], u.Server.users[i+1:]...)
			break
		}
	}
	u.Server.usersLock.Unlock()
}

func (u *User) Send(msg string) {
	u.Server.Broadcast(u, msg)
}
