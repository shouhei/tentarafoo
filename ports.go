package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func openTcpPorts(conf Config) {
	log.Printf("info: start to open any tcp ports")
	for _, p := range conf.TcpPorts {
		openTcpPort(p)
	}
}

func openUdpPorts(conf Config) {
	log.Printf("info: start to open any udp ports")
	for _, p := range conf.UdpPorts {
		openUdpPort(p)
	}
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
	openPort("tcp", port)
}

func openUdpPort(port int) {
	openPort("udp", port)
}
