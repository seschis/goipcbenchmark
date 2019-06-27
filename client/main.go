package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	fmt.Println("I'm the client")

	var number = flag.Int("n", 10000, "number of calls to make")
	var ping_sz = flag.Int("s", 1024, "size of ping msg")

	flag.Parse()

	conn, err := makeConn()
	if err != nil {
		panic("error making conn")
	}
	defer conn.Close()

	count := *number
	start := time.Now()
	for i := 0; i < count; i++ {
		sendPing(conn, *ping_sz)
	}
	elapsed := time.Since(start)
	fmt.Printf("%v ops. elapsed %s per op\n", count, elapsed/time.Duration(count))
}

func sendPing(conn net.Conn, ping_sz int) {
	ping := make([]byte, ping_sz)

	binary.LittleEndian.PutUint32(ping, 4)

	for i := 4; i < cap(ping); i++ {
		ping[i] = byte(i)
	}

	//fmt.Printf("sending ping [%v]\n", ping)
	total_outn := 0
	for total_outn < len(ping) {
		n, err := conn.Write(ping)
		if err != nil {
			log.Fatalf("ping write err: %v", err)
		}
		total_outn += n
	}

	//fmt.Printf("waiting for pong\n")

	inbuf := make([]byte, 4096)
	total_in := 0
	for total_in < 4 {
		n, err := conn.Read(inbuf)
		if err != nil {
			log.Fatalf("err reading pong: %v", err)
		}
		//fmt.Printf("got pong bytes (%v) %v\n", n, inbuf[0:n])
		total_in += n
	}

	//fmt.Printf("done pong\n")
}
