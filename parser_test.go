package main

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	cases := []struct {
		title            string
		r                *Request
		expectedResponse *CMD
	}{
		{
			title: "test command with argument",
			r: &Request{
				Message:         "!Test 123",
				Channel:         "1",
				User:            "1",
				Timestamp:       "1",
				ThreadTimestamp: "1",
			},
			expectedResponse: &CMD{
				Command:         "test",
				Option:          "123",
				Channel:         "1",
				User:            "1",
				Timestamp:       "1",
				ThreadTimestamp: "1",
			},
		},
		{
			title: "no command only the preffix",
			r: &Request{
				Message:         "!",
				Channel:         "1",
				User:            "1",
				Timestamp:       "1",
				ThreadTimestamp: "1",
			},
			expectedResponse: nil,
		},
		{
			title: "test command no argument",
			r: &Request{
				Message:         "!test",
				Channel:         "1",
				User:            "1",
				Timestamp:       "1",
				ThreadTimestamp: "1",
			},
			expectedResponse: &CMD{
				Command:         "test",
				Option:          "",
				Channel:         "1",
				User:            "1",
				Timestamp:       "1",
				ThreadTimestamp: "1",
			},
		},
		{
			title: "empty message",
			r: &Request{
				Message:         "",
				Channel:         "1",
				User:            "1",
				Timestamp:       "1",
				ThreadTimestamp: "1",
			},
			expectedResponse: nil,
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			cmd := Parse(c.r)
			if !reflect.DeepEqual(cmd, c.expectedResponse) {
				t.Errorf("%#v doesn't not correspond excepted responce %#v", cmd, c.expectedResponse)
			}
		})
	}
}
