package main

import (
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf) // reads the buffer
	if err != nil {
		log.Fatal(err)
	}

	// do something
	time.Sleep(8 * time.Second)

	// send a http response to the client
	conn.Write([]byte("HTTP/1.0 200 OK\r\n\r\nHello, Piu!\r\n\r\n"))

	conn.Close()
}

func main() {
	// our server listening at port 1729
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("Wiating for a client to connect...")
		conn, err := listener.Accept() // We got a client!
		if err != nil {
			log.Fatal(err)
		}
		go do(conn) // this processing with now be executed in a separate goroutine to leverage multi-threading
	}
}
