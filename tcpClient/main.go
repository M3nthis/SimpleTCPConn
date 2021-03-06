package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	for {
		reader := bufio.NewReader(os.Stdin)
		testo, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, testo+"\n")

		// NewReader ritorna un nuovo Reader il cui buffer il cui buffer ha dimensione di default
		messaggio, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("Messaggio ricevuto da server:\n%s\n", messaggio)
		fmt.Printf("\n")
	}
}
