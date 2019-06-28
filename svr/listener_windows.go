// +build windows

package main

import (
	"github.com/microsoft/go-winio"
	"net"
)

var pipeName = `\\.\pipe\winiotestpipe`

func getListener() (net.Listener, error) {
	l, err := winio.ListenPipe(pipeName, nil)
	if err != nil {
		return nil, err
	}

	return l, nil
}
