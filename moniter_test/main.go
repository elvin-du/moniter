package main

import (
	"fmt"
	//"io"
	"bufio"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":9191")
	if nil != err {
		log.Fatal(err)
	}
	defer conn.Close()

	//_, err = conn.Write([]byte("Hey, I am client"))
	_, err = fmt.Fprintln(conn, "Hey, I am client")
	if nil != err {
		log.Println("client write failed:", err)
	}

	//bin := make([]byte, 0, 256)
	//_, err = conn.Read(bin)
	//_, err = io.ReadFull(conn, bin)
	line, err := bufio.NewReader(conn).ReadString('\n')
	if nil != err {
		log.Println("client read failed:", err)
		return
	}
	//log.Println("get from server:", bin)
	log.Println("get from server:", line)
}
