package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Panicln("Connection not established: ", err)
	}

	for {
		fmt.Print("Write a message: ")
		s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if s == "exit\n" {
			fmt.Println("End of session")
			break
		}
		// Stdout is conn
		//fmt.Fprintln(conn, "Client says: ", s)
		conn.Write([]byte(s))

	}

}
