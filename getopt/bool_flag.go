package getopt

import (
	"fmt"
	"strings"
)

type boolValue bool

func (b *boolValue) Set(value string, opt *Option) error {
	switch strings.ToLower(value) {
	case "", "1", "true", "on", "t":
		*b = true
	case "0", "false", "off", "f":
		*b = false
	default:
		return fmt.Errorf("invalid value for bool %s: %q", opt.name, value)
	}
	return nil
}

func (b *boolValue) String() string {
	if *b {
		return "true"
	}
	return "false"
}

func (s *Set) AddBool(short rune, long string, value boolValue, help string, optional bool) {
	s.addFlags(short, long, &value, help, optional, true)
}
