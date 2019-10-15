package getopt

type optionList []*Option

type Option struct {
	short    rune   // 0 means no short name
	long     string // "" means no long name
	flag     bool
	optional bool
	value    Value
	name     string
	uname    string
	help     string
}

type Value interface {
	Set(value string, opt *Option) error

	String() string
}

func (s *Set) addFlags(short rune, long string, value Value, help string, optional bool, flag bool) {
	opt := &Option{
		short:    short,
		long:     long,
		value:    value,
		help:     help,
		optional: optional,
	}
	for _, eopt := range s.options {
		if opt == eopt {
			return
		}
	}
	if opt.short != 0 {
		if _, ok := s.shortOptions[opt.short]; ok {
			return
		}
		s.shortOptions[opt.short] = opt
	}
	if opt.long != "" {
		if _, ok := s.longOptions[opt.long]; ok {
			return
		}
		s.longOptions[opt.long] = opt
	}

	s.options = append(s.options, opt)
}
