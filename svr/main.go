package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	fmt.Println("I'm the svr\n")

	l, err := getListener()
	if err != nil {
		panic("err making listener")
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			panic("got error in accept")
		}
		go handleConn(conn)
	}
}

var pong = [...]byte{0, 1, 2, 3}

func handleConn(conn net.Conn) {
	b := make([]byte, 32768)

	//fmt.Printf("got a connection!\n")

	for {

		n, err := io.ReadAtLeast(conn, b, 4)
		if err != nil {
			//log.Fatalf("err reading hdr: %v", err)
			return
		}
		size := int(binary.LittleEndian.Uint32(b[0:4]))
		//fmt.Printf("got hdr (%v) buffer has %v bytes\n", size, n)
		//fmt.Printf("buf: %v\n", b[0:n])

		for (n - 4) < size {
			//fmt.Printf("reading payload\n")
			in_sz, err := conn.Read(b)
			if err != nil {
				log.Fatalf("bad read: %v", err)
			}
			n += in_sz
			//fmt.Printf("read %v bytes (%v of %v read)\n", in_sz, n, size)
			// XXX: ignore b since I only care about the transfer happending

		}

		// After recieving the ping data, we have to respond with pong.
		//fmt.Printf("sending pong\n")
		total_outn := 0
		for total_outn < len(pong) {
			outn, err := conn.Write(pong[total_outn:])
			if err != nil {
				log.Fatalf("pong failed: %v", err)
			}
			//fmt.Printf("send %v bytes of pong\n", outn)
			total_outn += outn
		}
	}
	return
}
