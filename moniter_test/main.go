package main

import (
	"fmt"
	//"io"
	//"bufio"
	"log"
	"net"
	//"time"
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

	bin := make([]byte, 256)
	//conn.SetDeadline(time.Now().Add(time.Minute))
	_, err = conn.Read(bin)
	//_, err = io.ReadFull(conn, bin)
	//line, err := bufio.NewReader(conn).ReadString('\n')
	if nil != err {
		log.Println("client read failed:", err)
		return
	}
	log.Println("get from server:", string(bin))
	//log.Println("get from server:", line)
}
