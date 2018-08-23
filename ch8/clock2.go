package main

import (
	"net"
	"log"
	"io"
	"time"
)

// ./clock1 & nc localhost 8000
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn2(conn) // 一度に一つの接続を処理する
	}
}

func handleConn2(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // クライアントの接続が切れた
		}
		time.Sleep(1 * time.Second)
	}
}
