package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	ProcessTitle string `yaml:"process_title"`
	TcpPorts     []int  `yaml:"tcp_ports"`
	UdpPorts     []int  `yaml:"udp_ports"`
}

func getDefaultConfig() Config {
	return Config{
		ProcessTitle: "tentarafoo",
		TcpPorts:     []int{},
		UdpPorts:     []int{},
	}
}

func mergeDefaultAndLoaded(def, loaded Config) Config {
	log.Printf("info: merge default config and load config")
	if loaded.ProcessTitle == "" {
		loaded.ProcessTitle = def.ProcessTitle
	}
	return loaded
}

func loadConfig(filePath string) (Config, error) {
	c := getDefaultConfig()
	if filePath == "" {
		log.Printf("info: use default config")
		return c, nil
	} else {
		log.Printf("info: config file path is " + filePath)
		buf, err := ioutil.ReadFile(filePath)
		if err != nil {
			return Config{}, err
		}
		var loadConf Config
		err = yaml.Unmarshal(buf, &loadConf)
		if err != nil {
			return Config{}, err
		}
		loaded := mergeDefaultAndLoaded(c, loadConf)
		return loaded, nil
	}
}
