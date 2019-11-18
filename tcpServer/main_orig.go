package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func mainOrig() {
	fmt.Println("Server partito...")

	listener, _ := net.Listen("tcp", ":8080")

	conn, _ := listener.Accept()

	for {
		mex, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("Messaggio ricevuto: %s", mex)

		nuovoMessaggio := strings.ToUpper(mex)
		conn.Write([]byte(nuovoMessaggio + "\n"))
	}
}
