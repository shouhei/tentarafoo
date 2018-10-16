package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit"
)

func showStdoutFakeLogs(c Config) {
	log_levels := []string{
		"trace",
		"debug",
		"info",
		"warn",
		"error",
		"alert",
	}
	rand.Seed(time.Now().UnixNano())
	if c.BurstStdoutFakeLog {
		for {
			log.Printf(log_levels[rand.Intn(6)] + ": " + gofakeit.HackerPhrase())
		}
	} else {
		for {
			log.Printf(log_levels[rand.Intn(6)] + ": " + gofakeit.HackerPhrase())
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		}
	}
}
