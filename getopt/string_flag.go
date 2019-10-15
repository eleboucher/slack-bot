package getopt

type stringValue string

func (s *stringValue) Set(value string, opt *Option) error {
	*s = stringValue(value)
	return nil
}

func (s *stringValue) String() string {
	return string(*s)
}

func (s *Set) AddString(short rune, long string, value stringValue, help string, optional bool) {
	s.addFlags(short, long, &value, help, optional, false)
}
