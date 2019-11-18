package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Server partito...")
	godotenv.Load()

	listener, _ := net.Listen("tcp", os.Getenv("port"))
	conn, _ := listener.Accept()

	for {
		reader := bufio.NewReader(conn)
		mess, _ := reader.ReadString('\n')

		// fmt.Fprintf(conn, strings.ToUpper(mess))
		conn.Write([]byte(createSpaces(mess) + "\n"))
	}
}

func createSpaces(s string) string {
	s = strings.ReplaceAll(s, " ", "_")
	fmt.Println(s)
	return s
}
