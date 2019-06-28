// +build windows

package main

import (
	"github.com/microsoft/go-winio"
	"net"
)

var pipeName = `\\.\pipe\winiotestpipe`

func makeConn() (net.Conn, error) {
	c, err := winio.DialPipe(pipeName, nil)
	if err != nil {
		return nil, err
	}

	return c, nil
}
