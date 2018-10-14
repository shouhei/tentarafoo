package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func openTcpPorts(conf Config) {
	log.Printf("info: start to open any tcp ports")
	var wg sync.WaitGroup
	for _, p := range conf.TcpPorts {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			openTcpPort(p)
		}(p)
	}
	wg.Wait()
}

func openUdpPorts(conf Config) {
	log.Printf("info: start to open any udp ports")
	var wg sync.WaitGroup
	for _, p := range conf.UdpPorts {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			openUdpPort(p)
		}(p)
	}
	wg.Wait()
}

func openPort(protocol string, port int) {
	log.Printf("info: try to open %s %d port", protocol, port)
	_, err := net.Listen(protocol, "127.0.0.1:"+fmt.Sprint(port))
	if err != nil {
		log.Printf("error: failed to open %s %d port", protocol, port)
		log.Fatal(err)
	}
	log.Printf("info: success to open %s %d port", protocol, port)
	for {
		time.Sleep(1 * time.Second)
	}
}

func openTcpPort(port int) {
	log.Printf("info: try to open tcp %d port", port)
	_, err := net.Listen("tcp", "127.0.0.1:"+fmt.Sprint(port))
	if err != nil {
		log.Printf("error: failed to open tcp %d port", port)
		log.Fatal(err)
	}
	log.Printf("info: success to open tcp %d port", port)
	for {
		time.Sleep(1 * time.Second)
	}
}

func openUdpPort(port int) {
	log.Printf("info: try to open udp %d port", port)
	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: port,
	}
	_, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Printf("error: failed to open udp %d port", port)
		log.Fatal(err)
	}
	log.Printf("info: success to open udp %d port", port)
	for {
		time.Sleep(1 * time.Second)
	}
}
