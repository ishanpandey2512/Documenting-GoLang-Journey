package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main1() {
	server, _ := net.ResolveTCPAddr("tcp", "google.com:72")
	client, _ := net.ResolveTCPAddr("tcp", ":50000")
	conn, err := net.DialTCP("tcp", client, server)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer conn.Close()
	fmt.Println(conn.LocalAddr())
	fmt.Println(conn.RemoteAddr())
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}
