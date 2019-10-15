package getopt

import (
	"errors"
	"strings"
)

type Set struct {
	value        string
	shortOptions map[rune]*Option
	longOptions  map[string]*Option
	options      optionList
}

func NewSet() *Set {
	return &Set{
		shortOptions: make(map[rune]*Option),
		longOptions:  make(map[string]*Option),
	}
}
func (s *Set) Lookup(name interface{}) *Option {
	switch v := name.(type) {
	case rune:
		return s.shortOptions[v]
	case int:
		return s.shortOptions[rune(v)]
	case string:
		return s.longOptions[v]
	}
	return nil
}

func (s *Set) IsSet(name interface{}) bool {
	if opt := s.Lookup(name); opt != nil {
		return true
	}
	return false
}

func (s *Set) GetValue(name interface{}) string {
	if opt := s.Lookup(name); opt != nil {
		return opt.value.String()
	}
	return ""
}

func (s *Set) GetOption() string {
	return s.value
}

func (s *Set) ParseOpt(toBeParsed string) error {
	args := strings.Split(toBeParsed, " ")

	for _, arg := range args {
		if len(arg) >= 2 && arg[0] == '-' {
			if arg[1] != '-' {
				shargs := arg[1:]
				for shargsIdx, c := range shargs {

					opt := s.shortOptions[c]
					if opt == nil {
						return errors.New("Option is unknown")
					}

					if len(shargs) < shargsIdx+1 && shargs[shargsIdx+1] != 0 {
						if err := opt.value.Set("true", opt); err != nil {
							return err
						}
						continue
					}
					var value string

					if len(args) > 1 {
						if args[1][0] != '-' {
							value = args[1]
							args = args[2:]
						}
					}
					if err := opt.value.Set(value, opt); err != nil {
						return err
					}
				}
			} else {
				option := arg[2:]
				opt := s.longOptions[option]
				if opt == nil {
					return errors.New("Option is unknown")
				}
				var value string
				if len(args) > 1 {
					if args[1][0] != '-' {
						value = args[1]
						args = args[2:]
					}
				}
				if err := opt.value.Set(value, opt); err != nil {
					return err
				}
			}
		}
		s.value = arg
	}
	return nil
}
