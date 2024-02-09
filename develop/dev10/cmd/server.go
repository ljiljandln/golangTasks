package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func msgHandle(conn net.Conn) {
	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		switch msg {
		case "time is up":
			return
		default:
			fmt.Print("Message from client: ", msg)
			_, err := conn.Write([]byte(time.Now().String() + " " + msg + "\n"))
			if err != nil {
				log.Print(err)
			}
		}
	}
}

func main() {
	ln, _ := net.Listen("tcp", "localhost:3000")
	conn, _ := ln.Accept()

	msgHandle(conn)

	err := ln.Close()
	if err != nil {
		log.Print(err)
	}
}
