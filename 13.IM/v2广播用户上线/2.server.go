package main

import (
	"fmt"
	"net"
	"sync"
	"sync/atomic"
)

type Server struct {
	Port      int
	Host      string
	users     []*User
	usersLock sync.RWMutex
	userCount atomic.Int64 // 专用ID计数器，只增不减；原子操作保证并发安全
	Message   chan string
}

func (s *Server) Handle(conn net.Conn) {
	//创建user：ID由原子计数器分配，Add(1)原子自增并返回新值，
	//即使以后支持用户下线删除，ID也不会重复
	id := s.userCount.Add(1)
	user := NewUser(int(id), conn)

	s.usersLock.Lock()
	s.users = append(s.users, user)
	s.usersLock.Unlock()

	//广播用户上线
	s.Message <- fmt.Sprintf("用户 %d 上线了 %s", user.ID, user.Conn.RemoteAddr())

}

// 监听Message channel，打印用户上线信息
func (s *Server) ListenMessage() {
	for msg := range s.Message {
		//广播用户上线信息
		s.usersLock.RLock()
		for _, user := range s.users {
			user.ch <- msg
		}
		s.usersLock.RUnlock()
	}
}

func (s *Server) Start() {
	//1.创建套接字
	//2. 监听指定端口
	//5. 关闭连接
	sock, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		fmt.Println("创建套接字失败 net.Listen:", err)
		return
	}
	defer sock.Close()

	
	//3. 等待客户端连接
	//4. 处理客户端请求
	for {
		conn, err := sock.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		fmt.Println("Accept success:", conn.RemoteAddr())
		//6. 处理客户端请求
		go s.Handle(conn)
	}

}

func CreateServer(port int, host string) *Server {
	
	server := &Server{
		Port:      port,
		Host:      host,
		Message:   make(chan string),
		usersLock: sync.RWMutex{},
		userCount: atomic.Int64{},
	}
	//监听Message channel
	go server.ListenMessage()
	return server
}
