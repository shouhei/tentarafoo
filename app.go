package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/comail/colog"
	"github.com/erikdubbelboer/gspt"
)

func initializeLogger() {
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()
}

func changeProcessTitle(c Config) {
	log.Printf("info: change process titlte to %s", c.ProcessTitle)
	gspt.SetProcTitle(c.ProcessTitle)
}

func wait60Seconds() {
	log.Printf("info: wait for 60 seconds")
	time.Sleep(60 * time.Second)
}

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

func main() {
	initializeLogger()
	log.Printf("info: tentarafoo cause server unexpected state inflict us confuse.")
	var (
		config_path = flag.String("c", "", "configration file path")
	)
	flag.Parse()
	c, error := loadConfig(*config_path)
	if error != nil {
		log.Printf("error: failed to load config file")
		os.Exit(1)
	}
	changeProcessTitle(c)
	go openTcpPorts(c)
	go openUdpPorts(c)
	wait60Seconds()
	log.Printf("info: completed")
}
