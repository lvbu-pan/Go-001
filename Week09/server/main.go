package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

const (
	address = "127.0.0.1"
	port    = 6379
	timeout = 15
)

func setKeepAliveTime() time.Time {
	return time.Now()
}

func readMessageFunc(conn net.Conn, ch chan<- []byte) {
	keepAliveTime := setKeepAliveTime()

	reader := bufio.NewReader(conn)
	fmt.Println(conn.RemoteAddr(), keepAliveTime)
	for {

		header, err := reader.Peek(4)
		if errors.Is(err, io.EOF) || len(header) == 0 {
			close(ch)
			_ = conn.Close()
			fmt.Println("Server close connection ")
			break
		}

		headerSize := binary.BigEndian.Uint32(header)
		pkg, err := reader.Peek(int(headerSize)+4)
		if err != nil {
			keepAliveTime = setKeepAliveTime()
			ch <- []byte("Body过大")
			continue
		}
		body := pkg[4:]
		ch <- body
		keepAliveTime = setKeepAliveTime()
		reader.Reset(conn)
	}
}

func writeMessageFunc(conn net.Conn, ch <-chan []byte) {
	writer := bufio.NewWriter(conn)
	defer conn.Close()
	for msg := range ch {
		head := make([]byte, 4)
		response := append([]byte("response: "), msg...)
		binary.BigEndian.PutUint32(head[0:4], uint32(len(response)))
		_, _ = writer.Write(head)
		_, _ = writer.Write(response)
		_ = writer.Flush()
		fmt.Println("success")
	}
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		ch := make(chan []byte, 10)
		go readMessageFunc(conn, ch)
		go writeMessageFunc(conn, ch)
	}
}
