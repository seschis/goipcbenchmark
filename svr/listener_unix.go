package main

import (
	"net"
)

var sockname = "benchunix.sock"

func getListener() (net.Listener, error) {
	l, err := net.Listen("unix", sockname)
	if err != nil {
		return nil, err
	}

	return l, nil
}

//var pipename = "benchunix.pipe"

//func getPipeListener() (net.Listener, error)
