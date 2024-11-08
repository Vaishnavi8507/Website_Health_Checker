package main

import (
	"fmt"
	"net"
	"time"
)

func Check(domain, port string) string {
	address := domain + ":" + port // Use `domain` instead of `destination`
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)

	var status string
	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable, \nError: %v", domain, err)
	} else {
		// Convert net.Addr to string
		status = fmt.Sprintf("[UP] %v is reachable,\n From: %v\n To: %v", domain, conn.LocalAddr(), conn.RemoteAddr())
		conn.Close() // Don't forget to close the connection after using it
	}

	return status
}


