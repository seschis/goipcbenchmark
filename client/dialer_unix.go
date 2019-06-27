package main

import (
	"net"
)

var sockname = "benchunix.sock"

func makeConn() (net.Conn, error) {
	c, err := net.Dial("unix", sockname)
	if err != nil {
		return nil, err
	}

	return c, nil
}

