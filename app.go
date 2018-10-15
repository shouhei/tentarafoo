package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"sync"
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
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	wg.Add(1)
	go func(c Config) {
		defer wg.Done()
		if len(c.TcpPorts) == 0 && c.RandomTcpPorts {
			for i := 0; i < 10; i++ {
				c.TcpPorts = append(c.TcpPorts, (rand.Intn(48127) + 1024))
			}
		}
		openTcpPorts(c)
	}(c)
	wg.Add(1)
	go func(c Config) {
		defer wg.Done()
		if len(c.UdpPorts) == 0 && c.RandomUdpPorts {
			for i := 0; i < 10; i++ {
				c.UdpPorts = append(c.UdpPorts, (rand.Intn(48127) + 1024))
			}
		}
		openUdpPorts(c)
	}(c)
	wg.Wait()
}
