package main

import (
	"github.com/jaymzee/img/term"
	"os"
	"strconv"
	"strings"
)

var cfg Config

type Config struct {
	eFlag    bool
	fFlag    bool
	lFlag    bool
	pFlag    bool
	rFlag    bool
	sFlag    float64
	nFlag    string
	terminal string
	termXres uint16
	termYres uint16
	plot     bool
}

func (c *Config) ProcessFlags() {
	if c.pFlag || c.rFlag || c.sFlag < 0.0 || c.fFlag {
		c.plot = true
	}

	if strings.Contains(os.Getenv("TERM"), "kitty") && term.Isatty() {
		c.terminal = "kitty"
	}

	env := strings.ToLower(os.Getenv("WAVDUMP"))
	for _, expr := range strings.Split(env, ":") {
		s := strings.Split(expr, "=")
		switch s[0] {
		case "term":
			c.terminal = s[1]
		case "xres":
			val, _ := strconv.ParseUint(s[1], 10, 16)
			c.termXres = uint16(val)
		case "yres":
			val, _ := strconv.ParseUint(s[1], 10, 16)
			c.termYres = uint16(val)
		case "nogfx":
			c.terminal = ""
		}
	}
}
