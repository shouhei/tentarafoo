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
	log_level_len := len(log_levels)
	rand.Seed(time.Now().UnixNano())
	if c.BurstStdoutFakeLog {
		for {
			log.Printf(log_levels[rand.Intn(log_level_len)] + ": " + gofakeit.HackerPhrase())
		}
	} else {
		for {
			log.Printf(log_levels[rand.Intn(log_level_len)] + ": " + gofakeit.HackerPhrase())
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		}
	}
}
