package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	ProcessTitle       string `yaml:"process_title"`
	TcpPorts           []int  `yaml:"tcp_ports"`
	UdpPorts           []int  `yaml:"udp_ports"`
	RandomTcpPorts     bool   `yaml:"random_tcp_ports"`
	RandomUdpPorts     bool   `yaml:"random_udp_ports"`
	ShowStdoutFakeLogs bool   `yaml:"show_stdout_fake_logs"`
	BurstStdoutFakeLog bool   `yaml:"burst_stdout_fake_log"`
	InodeExhaustion    bool   `yaml:"inode_exhaustion"`
}

func getDefaultConfig() Config {
	return Config{
		ProcessTitle:       "tentarafoo",
		TcpPorts:           []int{},
		UdpPorts:           []int{},
		RandomTcpPorts:     false,
		RandomUdpPorts:     false,
		ShowStdoutFakeLogs: false,
		BurstStdoutFakeLog: false,
		InodeExhaustion:    false,
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
