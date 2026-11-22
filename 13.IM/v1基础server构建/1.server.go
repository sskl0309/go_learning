package main

import (
	"fmt"
	"net"
)

type Server struct {
	Port int
	Host string
}

func (s *Server) Handle(conn net.Conn) {
	//7. 处理客户端请求
	fmt.Println("处理客户端请求:", conn.RemoteAddr())
	//8. 关闭连接
	defer conn.Close()
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
	return &Server{
		Port: port,
		Host: host,
	}
}
