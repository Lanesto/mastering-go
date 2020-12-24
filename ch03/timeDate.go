package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

type parser struct {
	expr   *regexp.Regexp
	format string
}

var logs = []string{
	"127.0.0.1 - - [16/Nov/2017:10:49:46 +0200] 325504",
	"127.0.0.1 - - [16/Nov/2017:10:16:41 +0200] \"GET /CVEN HTTP/1.1\" 200 12531 \"-\" \"Mozilla/5.0 AppleWebKit/537.36",
	"127.0.0.1 200 9412 - - [12/Nov/2017:06:26:05 +0200] \"GET \"http://www.mtsoukalos.eu/taxonomy/term/47\" 1507",
	"[12/Nov/2017:16:27:21 +0300]",
	"[12/Nov/2017:20:88:21 +0200]",
	"[12/Nov/2017:20:21 +0200]",
}

var parsers = [...]parser{
	parser{
		expr:   regexp.MustCompile(`.*\[(\d\d/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\].*`),
		format: "02/Jan/2006:15:04:05 -0700",
	},
	parser{
		expr:   regexp.MustCompile(`.*\[(\d\d/\w+/\d\d\d\d:\d\d:\d\d [+-]\d\d\d\d)\].*`),
		format: "02/Jan/2006:15:04 -0700",
	},
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %s dateformat\n", filepath.Base(args[0]))
		return
	}

	// format := "2006 Jan 02 15:04:05 -0700"
	format := args[1]

	for _, logEntry := range logs {
		ok := false
		for _, parser := range parsers {
			if parser.expr.MatchString(logEntry) {
				match := parser.expr.FindStringSubmatch(logEntry)
				dt, err := time.Parse(parser.format, match[1])
				if err == nil {
					newDate := dt.Format(format)
					fmt.Println(newDate)
					ok = true
				}
			}
			if ok {
				break
			}
		}
		if !ok {
			fmt.Println("no match for given formats")
		}
	}
}
