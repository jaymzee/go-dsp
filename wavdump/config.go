package main

import (
	"github.com/jaymzee/img/term"
	"os"
	"strings"
)

type Config struct {
	eFlag    bool
	fFlag    bool
	lFlag    bool
	pFlag    bool
	rFlag    bool
	sFlag    float64
	nFlag    string
	terminal string
	plot     bool
}

var cfg Config

func (c *Config) ProcessFlags() {
	if c.pFlag || c.rFlag || c.sFlag < 0.0 || c.fFlag {
		c.plot = true
	}
	if !strings.Contains(os.Getenv("WAVDUMP"), "nogfx") {
		if strings.Contains(os.Getenv("TERM"), "kitty") && term.Isatty() {
			c.terminal = "kitty"
		}
		if strings.Contains(os.Getenv("WAVDUMP"), "iTerm") {
			c.terminal = "iTerm"
		}
	}
}
