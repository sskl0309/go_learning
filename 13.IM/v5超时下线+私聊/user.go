package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
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

func (u *User) DoMessage(msg string) {
	if msg == "who" {
		u.Server.usersLock.Lock()
		for _, user := range u.Server.users {
			u.sendMessage(fmt.Sprintf("[用户%d %s]:[在线]", user.ID, user.Conn.RemoteAddr().String()))
		}
		u.Server.usersLock.Unlock()
	} else if len(msg) > 4 && msg[:3] == "to|" {
		parts := strings.Split(msg, "|")
		if len(parts) < 3 {
			u.sendMessage("格式错误，应为 to|用户ID|内容")
			return
		}
		// string转int：strconv.Atoi，成功返回数值，失败返回error
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			u.sendMessage("用户ID必须是数字")
			return
		}

		u.Server.usersLock.Lock()
		user, ok := u.Server.users[id]
		u.Server.usersLock.Unlock() // 查完立刻解锁，避免分支中return导致死锁
		if !ok {
			u.sendMessage("用户不存在")
			return
		}
		user.sendMessage(fmt.Sprintf("用户%d对您说:%s", u.ID, parts[2]))
	} else {
		u.Server.Broadcast(u, msg)
	}
}
