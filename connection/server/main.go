package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConnection(c net.Conn) {
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println("Received: ", ln)
	}
	return
}

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln("Error while listening: ", err)
	}

	for {
		fmt.Println("Accepting connections...")
		conn, err := ln.Accept()
		if err != nil {
			log.Panicln("Error after accepting: ", err)
		}
		fmt.Println("Conenction accepted")

		go handleConnection(conn)
	}

	fmt.Println("Shuting down server...")
}
