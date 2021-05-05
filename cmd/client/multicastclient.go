package main

import (
	"encoding/hex"
	"log"
	"net"
	"time"
)

const (
	srvAddr         = "[ff15:1234:0:0:0:0:0:0]:3000"
	maxDatagramSize = 8192
)

func main() {
//	go ping(srvAddr)
//	serveMulticastUDP(srvAddr, msgHandler)
	for {
		ping(srvAddr)
	}
}

func ping(address string) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatal(err)
	}
	c, err := net.DialUDP("udp", nil, addr)
	for {
		c.Write([]byte("hello, world\n"))
		time.Sleep(1 * time.Second)
	}
}

func msgHandler(src *net.UDPAddr, bytesRead int, buffer []byte) {
	log.Println(bytesRead, "bytes read from", src)
	log.Println(hex.Dump(buffer[:bytesRead]))
}

func serveMulticastUDP(address string, servicehandler func(*net.UDPAddr, int, []byte)) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatal(err)
	}
	listen, err := net.ListenMulticastUDP("udp", nil, addr)
	listen.SetReadBuffer(maxDatagramSize)
	for {
		buffer := make([]byte, maxDatagramSize)
		bytesRead, src, err := listen.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}
		servicehandler(src, bytesRead, buffer)
	}
}
