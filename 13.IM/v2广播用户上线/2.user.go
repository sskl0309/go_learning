package main

import (
	"net"
)

type User struct {
	ID   int
	ch   chan string
	Conn net.Conn
}

func NewUser(id int, conn net.Conn) *User {

	user := &User{
		ID:   id,
		ch:   make(chan string),
		Conn: conn,
	}
	go user.ListenMessage()
	return user
}


func (u *User) ListenMessage() {
	for msg := range u.ch {
		u.Conn.Write([]byte(msg + "\n"))
	}
}