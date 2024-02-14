package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	data := []byte("Client Data Received")

	for i := 0; i < 10; i++ {
		_, err := conn.Write(data)
		if err != nil {
			log.Fatal(err)
		}

		log.Print("Client Data Sent", i+1)
	}
}
