package main

import "os"

type config struct {
	addr string
}

func loadConfig() config {
	c := config{}

	c.addr = os.Getenv("ADDR")
	if c.addr == "" {
		c.addr = ":8010"
	}

	return c
}
