package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":9000") // Listen on any IP Address
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			buf := make([]byte, 1024)
			_, err := c.Read(buf)
			if err != nil {
				log.Fatal()
			}
			log.Print(string(buf))
			conn.Write([]byte("Test TCP server"))
			c.Close()
		}(conn)
	}
}
