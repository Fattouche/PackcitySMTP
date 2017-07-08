package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	// Connect to the remote SMTP server.
	conn, err := net.Dial("tcp", "127.0.0.1:2525")
	if err != nil {
		log.Fatal(err)
	}

	in := bufio.NewReader(conn)

	str, err := in.ReadString('\n')
	//	fmt.Println(str)
	fmt.Fprint(conn, "Testing Message\r\n")
	str, err = in.ReadString('\n')
	//	fmt.Println(str)
	fmt.Fprint(conn, "MAIL FROM:<test@example.com>r\r\n")
	str, err = in.ReadString('\n')
	//	fmt.Println(str)
	fmt.Fprint(conn, "RCPT TO:test@grr.la\r\n")
	str, err = in.ReadString('\n')
	//	fmt.Println(str)
	fmt.Fprint(conn, "DATA\r\n")
	str, err = in.ReadString('\n')
	//	fmt.Println(str)
	fmt.Fprint(conn, "Subject: Test subject\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, "A an email body\r\n")
	fmt.Fprint(conn, ".\r\n")
	str, err = in.ReadString('\n')
	//	fmt.Println(str)
	_ = str
}
