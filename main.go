package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	log.Println("Moniter Begin")
	ln, err := net.Listen("tcp", ":9191")
	if nil != err {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if nil != err {
			log.Println("accept failed, but continue:", err)
			continue
		}

		go Service(conn)
	}
	log.Println("Moniter End")
}

func Service(conn net.Conn) {
	defer conn.Close()
	//_, err := conn.Write([]byte("I am server"))
	_, err := fmt.Fprintln(conn, "I am server")
	if nil != err {
		log.Println("server write failed:", err)
		return
	}

	bin := make([]byte, 0, 256)
	//_, err = conn.Read(bin)
	_, err = io.ReadFull(conn, bin)
	if nil != err {
		log.Println("server read failed:", err)
		return
	}
	log.Println("get from client:", bin)
}
