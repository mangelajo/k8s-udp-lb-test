package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	go udpListenOnPort(3000)
	udpListenOnPort(3001)
}

func udpListenOnPort(port int) {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("0.0.0.0"),
	})
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	fmt.Printf("server listening %s\n", conn.LocalAddr().String())

	for {
		message := make([]byte, 4096)
		rlen, remote, err := conn.ReadFromUDP(message[:])
		if err != nil {
			panic(err)
		}

		data := strings.TrimSpace(string(message[:rlen]))
		fmt.Printf("received: %s from %s\n", data, remote)
		conn.WriteTo(message[:rlen], remote)
	}
}
