package main

import (
	"flag"
	"log"
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
