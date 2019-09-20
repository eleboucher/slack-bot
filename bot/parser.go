package bot

import "strings"

//Parse the message
func Parse(r *Request) *CMD {
	var preffix string
	var option string

	s := strings.TrimSpace(r.Message)

	if s == "" {
		return nil
	}

	for _, cmdPrefix := range cmdPrefixes {
		if strings.HasPrefix(s, cmdPrefix) {
			preffix = cmdPrefix
		}
	}

	if preffix == "" {
		return nil
	}

	rawCommand := strings.TrimPrefix(s, preffix)
	splitted := strings.SplitN(rawCommand, " ", 2)

	if splitted[0] == "" {
		return nil
	}

	if len(splitted) >= 2 {
		option = strings.Join(strings.Fields(splitted[1]), " ")
	}

	return &CMD{
		Command:         strings.ToLower(splitted[0]),
		Option:          option,
		Channel:         r.Channel,
		User:            r.User,
		Timestamp:       r.Timestamp,
		ThreadTimestamp: r.ThreadTimestamp,
	}
}
