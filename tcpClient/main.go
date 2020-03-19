package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// Client conntects to server
func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	for {
		reader := bufio.NewReader(os.Stdin)
		testo, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, testo+"\n")

		messaggio, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("Messaggio ricevuto da server:\n%s\n", messaggio)
		fmt.Println("Messaggio ricevuto da server:\n%s", messaggio)
	}
}
