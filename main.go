package main

import (
	"bytes"
	"log"
	"os"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

func main() {
	Run(os.Getenv("SLACK_TOKEN"))
}
