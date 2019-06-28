package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var number = flag.Int("n", 10000, "number of calls to make")
	var ping_sz = flag.Int("s", 1024, "size of ping msg")
	var concurrent = flag.Int("c", 1, "number of concurrent pings")

	flag.Parse()

	fmt.Printf("%v concurrent connections\n", *concurrent)
	var ping = make([]byte, *ping_sz)

	binary.LittleEndian.PutUint32(ping, uint32(*ping_sz-4))

	for i := 4; i < cap(ping); i++ {
		ping[i] = byte(i)
	}

	var totalstart = time.Now()

	for i := 0; i < *concurrent; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("thread[%v] started\n", id)

			conn, err := makeConn()
			if err != nil {
				panic("error making conn")
			}
			defer conn.Close()

			var count = int(*number / (*concurrent))
			var start = time.Now()
			for j := 0; j < count; j++ {
				sendPing(conn, ping)
			}
			elapsed := time.Since(start)
			fmt.Printf("thread[%v] %v ops. elapsed %s per op\n", id, count, elapsed/time.Duration(count))
		}(i)
	}
	wg.Wait()
	var totalelapsed = time.Since(totalstart)
	fmt.Printf("total %v ops. total time %s, total elapsed %s per op", *number, totalelapsed, totalelapsed/time.Duration(*number))
}

func sendPing(conn net.Conn, ping []byte) {
	//ping := make([]byte, ping_sz)

	//binary.LittleEndian.PutUint32(ping, uint32(ping_sz-4))

	//for i := 4; i < cap(ping); i++ {
	//	ping[i] = byte(i)
	//}

	//fmt.Printf("sending ping [%v]\n", ping)
	var total_outn int
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
