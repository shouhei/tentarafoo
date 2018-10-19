package main

import (
	"io/ioutil"
	"log"
	"runtime"
	"sync"
)

func inodeExhaustion(c Config) {
	var wg sync.WaitGroup
	for i := 0; i < (runtime.NumCPU() * 4); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				tmpFile, err := ioutil.TempFile("", "")
				log.Printf("info: %s created", tmpFile.Name())
				if err != nil {
					log.Printf("error: create file failed")
					break
				}
			}
		}()
	}
	wg.Wait()
}
