package main

func main() {
	server := CreateServer(9000, "127.0.0.1")
	server.Start()
}
